package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/panjf2000/gnet/v2"
	"log"
	"net"
	"sync"
	"time"
)

const (
	startFlag       = 0xEB90
	endFlag         = 0xEB90
	xmlRegisterData = `<?xml version="1.0" encoding="UTF-8"?>
<PatrolHost>
    <SendCode>Client01</SendCode>
    <ReceiveCode>Server01</ReceiveCode>
    <Type>251</Type>
    <Code/>
    <Command>1</Command>
    <Time>2022-01-01 12:02:34</Time>
    <Items/>
</PatrolHost>`
	xmlHeartData = `<?xml version="1.0" encoding="UTF-8"?>
<PatrolHost>
    <SendCode>Client01</SendCode>
    <ReceiveCode>Server01</ReceiveCode>
    <Type>251</Type>
    <Code/>
    <Command>2</Command>
    <Time>2022-01-01 12:02:34</Time>
    <Items/>
</PatrolHost>`
)

type Message struct {
	StartFlag     uint16
	TransmitSeq   int64
	ReceiveSeq    int64
	SessionSource uint8
	XMLLength     int32
	XMLContent    string
	EndFlag       uint16
}

// String 方法用于美观地打印 Message 结构体
func (m Message) String() string {
	return fmt.Sprintf(
		"Message{\n"+
			"  StartFlag:     0x%04X,\n"+
			"  TransmitSeq:   %d,\n"+
			"  ReceiveSeq:    %d,\n"+
			"  SessionSource: 0x%02X,\n"+
			"  XMLLength:     %d bytes,\n"+
			"  XMLContent:    %s,\n"+
			"  EndFlag:       0x%04X\n"+
			"}",
		m.StartFlag, m.TransmitSeq, m.ReceiveSeq, m.SessionSource, m.XMLLength, m.XMLContent, m.EndFlag)
}

// 转换为小端字节序
func toBytes(v interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// 构造并写入消息内容
func writeBuffer(msg Message, buf *bytes.Buffer) {
	if startFlagBytes, err := toBytes(msg.StartFlag); err == nil {
		buf.Write(startFlagBytes)
	}
	if transmitSeqBytes, err := toBytes(msg.TransmitSeq); err == nil {
		buf.Write(transmitSeqBytes)
	}
	if receiveSeqBytes, err := toBytes(msg.ReceiveSeq); err == nil {
		buf.Write(receiveSeqBytes)
	}
	if sessionSourceBytes, err := toBytes(msg.SessionSource); err == nil {
		buf.Write(sessionSourceBytes)
	}
	if xmlLengthBytes, err := toBytes(msg.XMLLength); err == nil {
		buf.Write(xmlLengthBytes)
	}
	buf.WriteString(msg.XMLContent)
	if endFlagBytes, err := toBytes(msg.EndFlag); err == nil {
		buf.Write(endFlagBytes)
	}
}

// 客户端事件处理器
type clientEventHandler struct {
	*gnet.BuiltinEventEngine
	con         net.Conn
	TransmitSeq int64
	ReceiveSeq  int64
}

func (c *clientEventHandler) OnBoot(e gnet.Engine) (action gnet.Action) {
	fmt.Println("OnBoot")
	return
}

func (c *clientEventHandler) OnShutdown(_ gnet.Engine) {
	fmt.Println("OnShutdown")
	return
}

func (c *clientEventHandler) OnOpen(_ gnet.Conn) (out []byte, action gnet.Action) {
	fmt.Println("OnOpen")
	// 构造消息
	msg := Message{
		StartFlag:     startFlag,
		TransmitSeq:   c.TransmitSeq + 1,
		ReceiveSeq:    c.ReceiveSeq,
		SessionSource: 0x00,
		XMLLength:     int32(len(xmlRegisterData)),
		XMLContent:    xmlRegisterData,
		EndFlag:       endFlag,
	}

	// 构造字节流
	buf := new(bytes.Buffer)
	writeBuffer(msg, buf)
	hexStr := hex.EncodeToString(buf.Bytes())
	fmt.Printf("send: %s\n", hexStr)
	return buf.Bytes(), gnet.None
}

func (c *clientEventHandler) OnClose(_ gnet.Conn, _ error) (action gnet.Action) {
	fmt.Println("OnClose")
	return
}

func (c *clientEventHandler) OnTraffic(conn gnet.Conn) (action gnet.Action) {
	fmt.Println("OnTraffic")

	// 设置一个缓冲区来存储每次读取的数据，假设每次读取 1024 字节
	buf := make([]byte, 1024)

	// 用于拼接接收到的完整数据
	var fullData []byte

	// 不断读取数据，直到接收到完整消息
	for {
		n, err := conn.Read(buf)
		if err != nil {
			// 错误处理，读取数据失败时返回
			fmt.Printf("Read error: %v\n", err)
			return gnet.None
		}

		// 将读取到的数据追加到 fullData 中
		fullData = append(fullData, buf[:n]...)

		// 检查是否收到足够的字节以解析消息
		if len(fullData) >= 25 { // 25 字节是固定头部的长度
			// 尝试解析消息
			var msg Message
			err = parseMessage(fullData, &msg)
			if err == nil {
				// 如果解析成功，打印消息并清除已解析的数据
				fmt.Printf("Parsed message: %+v\n", msg)

				// 从 fullData 中去掉已解析的部分
				// 25 字节头部 + XML 内容的长度字节（msg.XMLLength）
				fullData = fullData[25+int(msg.XMLLength):]

				// 处理完当前消息后，继续读取新的数据
				break
			} else {
				// 如果消息不完整，输出提示并继续等待更多数据
				fmt.Println("Message is incomplete, waiting for more data.")
				continue
			}
		}
	}

	return gnet.None
}

// 解析字节流为 Message 结构体
func parseMessage(data []byte, msg *Message) error {
	// 确保最小的消息长度为 18 字节（包括头部信息和 XMLLength）
	if len(data) < 18 {
		return fmt.Errorf("invalid message length")
	}

	// 解析 StartFlag (大端字节序)
	msg.StartFlag = binary.BigEndian.Uint16(data[:2])

	// 解析 TransmitSeq (大端字节序)
	msg.TransmitSeq = int64(binary.BigEndian.Uint64(data[2:10]))

	// 解析 ReceiveSeq (大端字节序)
	msg.ReceiveSeq = int64(binary.BigEndian.Uint64(data[10:18]))

	// 解析 SessionSource
	msg.SessionSource = data[18]

	// 解析 XMLLength（4 字节，大端字节序）
	msg.XMLLength = int32(binary.BigEndian.Uint32(data[19:23]))

	// 检查 XMLLength 是否大于 0，如果是，解析 XMLContent
	if msg.XMLLength > 0 {
		// 解析 XMLContent
		xmlContentStart := 23
		xmlContentEnd := xmlContentStart + int(msg.XMLLength)
		if len(data) < xmlContentEnd {
			return fmt.Errorf("invalid XML length")
		}
		msg.XMLContent = string(data[xmlContentStart:xmlContentEnd])

		// 解析 EndFlag (大端字节序)
		msg.EndFlag = binary.BigEndian.Uint16(data[xmlContentEnd : xmlContentEnd+2])
	} else {
		// 如果 XMLLength 为 0，说明没有 XML 内容，直接解析 EndFlag
		msg.EndFlag = binary.BigEndian.Uint16(data[23:25])
	}
	return nil
}

func (c *clientEventHandler) OnTick() (delay time.Duration, action gnet.Action) {
	fmt.Println("OnTick")
	delay = 1 * time.Second
	// 构造消息
	msg := Message{
		StartFlag:     startFlag,
		TransmitSeq:   time.Now().Unix(),
		ReceiveSeq:    time.Now().Unix() + 1,
		SessionSource: 0x00,
		XMLLength:     int32(len(xmlHeartData)),
		XMLContent:    xmlHeartData,
		EndFlag:       endFlag,
	}

	// 构造字节流
	buf := new(bytes.Buffer)
	writeBuffer(msg, buf)
	if c.con != nil {
		hexStr := hex.EncodeToString(buf.Bytes())
		fmt.Printf("send: %s\n", hexStr)
		_, err := c.con.Write(buf.Bytes())
		if err != nil {
			return 0, gnet.Shutdown
		}
	}
	return
}

func startClient(address string) {
	client := &clientEventHandler{}
	// 使用 gnet 客户端连接到服务器
	err := gnet.Run(client, address, gnet.WithMulticore(false))
	if err != nil {
		log.Fatalf("Failed to start gnet client: %v", err)
	}
}

func main() {
	var wg sync.WaitGroup
	// 客户端目标服务器地址
	var err error
	clientEV := &clientEventHandler{}
	cli, err := gnet.NewClient(clientEV, gnet.WithTicker(true))
	if err != nil {
		log.Fatalf("Failed to create gnet client: %v", err)
	}
	err = cli.Start()
	if err != nil {
		log.Fatalf("Failed to start gnet client: %v", err)
	}
	defer cli.Stop()
	conn, err := cli.Dial("tcp", "127.0.0.1:7100")
	if err != nil {
		log.Fatalf("Failed to dial gnet client: %v", err)
	}
	clientEV.con = conn
	wg.Add(1)
	wg.Wait()
}
