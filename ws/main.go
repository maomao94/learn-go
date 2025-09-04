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
	// 升级连接
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

	// ✅ 正确的 Pong 处理：收到 Pong 时刷新超时
	conn.SetPongHandler(func(appData string) error {
		conn.SetReadDeadline(time.Now().Add(s.config.ReadTimeout))
		return nil
	})

	// ✅ 可选：收到 Ping 时立刻回复 Pong（有些实现需要）
	conn.SetPingHandler(func(appData string) error {
		log.Printf("收到 Ping，自动回复 Pong")
		return conn.WriteMessage(websocket.PongMessage, []byte(appData))
	})

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

	// 消息处理循环
	for {
		msgType, data, err := conn.ReadMessage()
		if err != nil {
			log.Printf("读取消息失败: %v", err)
			break
		}

		// ⚡ 协议 Ping/Pong 已经由 handler 自动处理，这里只关心 TextMessage
		if msgType != websocket.TextMessage {
			log.Printf("忽略非文本消息类型: %d", msgType)
			continue
		}

		log.Printf("收到消息: %s", string(data))

		// 应用级心跳：如果客户端发 "ping"
		if string(data) == "ping" {
			if err := conn.WriteMessage(websocket.TextMessage, []byte("pong")); err != nil {
				log.Printf("发送 pong 失败: %v", err)
				break
			}
			continue
		}

		// 👉 这里继续你的 JSON 消息处理逻辑
		var msg Message
		if err := json.Unmarshal(data, &msg); err != nil {
			log.Printf("不是 JSON，按普通文本回显")
			_ = conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("echo: %s", string(data))))
			continue
		}

		// … 认证逻辑、业务逻辑省略 …
	}

	log.Printf("客户端断开连接: %s", clientID)
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
