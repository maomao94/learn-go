package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// 定义消息类型
const (
	MessageTypeText   = "text"
	MessageTypeJSON   = "json"
	MessageTypeAuth   = "auth"
	MessageTypeError  = "error"
	MessageTypeSystem = "system"
)

// 消息结构
type Message struct {
	Type    string          `json:"type"`
	Content interface{}     `json:"content"`
	Time    string          `json:"time"`
	ID      int             `json:"id,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
}

// 认证请求
type AuthRequest struct {
	Token string `json:"token"`
}

// 服务器配置
type ServerConfig struct {
	Port           int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxMessageSize int64
	AuthRequired   bool
}

// WebSocket服务器
type WebSocketServer struct {
	config      ServerConfig
	upgrader    websocket.Upgrader
	clients     map[string]*websocket.Conn
	clientsMu   sync.Mutex
	messageID   int
	messageIDMu sync.Mutex
}

// 创建新服务器
func NewWebSocketServer(config ServerConfig) *WebSocketServer {
	return &WebSocketServer{
		config: config,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // 允许所有来源的连接
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		clients: make(map[string]*websocket.Conn),
	}
}

// 生成唯一消息ID
func (s *WebSocketServer) nextMessageID() int {
	s.messageIDMu.Lock()
	defer s.messageIDMu.Unlock()
	s.messageID++
	return s.messageID
}

// 发送系统消息给客户端
func (s *WebSocketServer) sendSystemMessage(conn *websocket.Conn, content string) error {
	msg := Message{
		Type:    MessageTypeSystem,
		Content: content,
		Time:    time.Now().Format(time.RFC3339),
		ID:      s.nextMessageID(),
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return conn.WriteMessage(websocket.TextMessage, data)
}

// 发送错误消息给客户端
func (s *WebSocketServer) sendErrorMessage(conn *websocket.Conn, content string) error {
	msg := Message{
		Type:    MessageTypeError,
		Content: content,
		Time:    time.Now().Format(time.RFC3339),
		ID:      s.nextMessageID(),
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return conn.WriteMessage(websocket.TextMessage, data)
}

// 处理客户端连接
func (s *WebSocketServer) handleConnection(w http.ResponseWriter, r *http.Request) {
	// 升级HTTP连接为WebSocket
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("升级连接失败: %v", err)
		return
	}
	defer conn.Close()

	// 设置连接超时
	conn.SetReadDeadline(time.Now().Add(s.config.ReadTimeout))
	conn.SetWriteDeadline(time.Now().Add(s.config.WriteTimeout))

	// 生成客户端ID
	clientID := fmt.Sprintf("client-%d", time.Now().UnixNano())
	log.Printf("新客户端连接: %s", clientID)

	// 添加到客户端列表
	s.clientsMu.Lock()
	s.clients[clientID] = conn
	s.clientsMu.Unlock()

	defer func() {
		// 从客户端列表移除
		s.clientsMu.Lock()
		delete(s.clients, clientID)
		s.clientsMu.Unlock()
		log.Printf("客户端断开连接: %s", clientID)
	}()

	// 发送欢迎消息
	if err := s.sendSystemMessage(conn, "欢迎连接到测试WebSocket服务器"); err != nil {
		log.Printf("发送欢迎消息失败: %v", err)
		return
	}

	// 如果需要认证，等待客户端发送认证信息
	authenticated := !s.config.AuthRequired
	if !authenticated {
		if err := s.sendSystemMessage(conn, "请发送认证信息 (type: auth, content: {\"token\": \"your-token\"})"); err != nil {
			log.Printf("发送认证提示失败: %v", err)
			return
		}
	}

	// 消息处理循环
	for {
		// 读取消息（获取消息类型）
		msgType, data, err := conn.ReadMessage()
		if err != nil {
			log.Printf("读取消息失败: %v", err)
			break
		}
		// 重置超时
		conn.SetReadDeadline(time.Now().Add(s.config.ReadTimeout))
		conn.SetWriteDeadline(time.Now().Add(s.config.WriteTimeout))

		// 关键修复：处理WebSocket控制帧（Ping/Pong）
		if msgType == websocket.PingMessage {
			// 收到Ping帧，立即返回Pong帧（保持连接活性）
			if err := conn.WriteMessage(websocket.PongMessage, data); err != nil {
				log.Printf("发送Pong响应失败: %v", err)
				break
			}
			continue // 处理完Ping帧，继续循环
		}

		// 只处理文本消息（忽略其他类型如二进制消息）
		if msgType != websocket.TextMessage {
			log.Printf("忽略非文本消息类型: %d", msgType)
			continue
		}

		// 处理JSON文本消息
		var msg Message
		if err := json.Unmarshal(data, &msg); err != nil {
			// 如果不是JSON格式，当作普通文本处理
			response := Message{
				Type:    MessageTypeText,
				Content: fmt.Sprintf("收到文本消息: %s", string(data)),
				Time:    time.Now().Format(time.RFC3339),
				ID:      s.nextMessageID(),
			}
			responseData, _ := json.Marshal(response)
			if err := conn.WriteMessage(websocket.TextMessage, responseData); err != nil {
				log.Printf("发送消息失败: %v", err)
				break
			}
			continue
		}

		// 处理认证消息
		if msg.Type == MessageTypeAuth && !authenticated {
			var authReq AuthRequest
			if err := json.Unmarshal([]byte(msg.Content.(string)), &authReq); err != nil {
				if err := s.sendErrorMessage(conn, "无效的认证格式"); err != nil {
					log.Printf("发送错误消息失败: %v", err)
					break
				}
				continue
			}

			// 简单的认证逻辑：token不为空即通过
			if authReq.Token != "" {
				authenticated = true
				if err := s.sendSystemMessage(conn, "认证成功"); err != nil {
					log.Printf("发送认证成功消息失败: %v", err)
					break
				}
				log.Printf("客户端 %s 认证成功", clientID)
			} else {
				if err := s.sendErrorMessage(conn, "认证失败：token不能为空"); err != nil {
					log.Printf("发送错误消息失败: %v", err)
					break
				}
			}
			continue
		}

		// 未认证的客户端不能发送其他消息
		if !authenticated {
			if err := s.sendErrorMessage(conn, "请先认证"); err != nil {
				log.Printf("发送错误消息失败: %v", err)
				break
			}
			continue
		}

		// 回声消息：将收到的消息原样返回
		response := Message{
			Type:    msg.Type,
			Content: msg.Content,
			Data:    msg.Data,
			Time:    time.Now().Format(time.RFC3339),
			ID:      s.nextMessageID(),
		}

		responseData, err := json.Marshal(response)
		if err != nil {
			log.Printf("序列化响应消息失败: %v", err)
			if err := s.sendErrorMessage(conn, "处理消息失败"); err != nil {
				log.Printf("发送错误消息失败: %v", err)
				break
			}
			continue
		}

		// 发送响应
		if err := conn.WriteMessage(websocket.TextMessage, responseData); err != nil {
			log.Printf("发送消息失败: %v", err)
			break
		}
	}
}

// 启动服务器
func (s *WebSocketServer) Start() error {
	http.HandleFunc("/ws", s.handleConnection)

	addr := fmt.Sprintf(":%d", s.config.Port)
	log.Printf("WebSocket服务器启动，地址: ws://localhost%s/ws", addr)
	return http.ListenAndServe(addr, nil)
}

func main() {
	// 命令行参数
	port := flag.Int("port", 18080, "服务器端口")
	authRequired := flag.Bool("auth", false, "是否需要认证")
	flag.Parse()

	// 创建服务器配置
	config := ServerConfig{
		Port:           *port,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxMessageSize: 1024 * 1024, // 1MB
		AuthRequired:   *authRequired,
	}

	// 创建并启动服务器
	server := NewWebSocketServer(config)
	if err := server.Start(); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
