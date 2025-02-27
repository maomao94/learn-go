package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/panjf2000/gnet/v2"
	"io"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

const (
	startFlag       = 0xEB90
	endFlag         = 0xEB90
	xmlRegisterData = `<?xml version="1.0" encoding="UTF-8"?>
<PatrolDevice>
    <SendCode>testDog</SendCode>
    <ReceiveCode>Server01</ReceiveCode>
    <Type>251</Type>
    <Code>1222</Code>
    <Command>1</Command>
    <Time>2022-01-01 12:02:34</Time>
    <Items/>
</PatrolDevice>`
	xmlBizData = `<PatrolDevice>
    <SendCode>testDog</SendCode>
    <ReceiveCode>Server01</ReceiveCode>
    <Code>变电站编码</Code>
    <Type>1</Type>
    <Time>2025-01-09 12:00:00</Time>
    <Items>
    <Item patroldevice_name="设备A" patroldevice_code="12345" time="2025-01-09T12:00:00" type="11" value="30" value_unit="低" unit="%"/>
    <Item patroldevice_name="设备B" patroldevice_code="67890" time="2025-01-09T12:05:00" type="20" value="0" value_unit="正常" unit="1"/>
    <Item patroldevice_name="设备C" patroldevice_code="54321" time="2025-01-09T12:10:00" type="31" value="0" value_unit="正常" unit="2"/>
    <Item patroldevice_name="设备D" patroldevice_code="11223" time="2025-01-09T12:15:00" type="41" value="1" value_unit="异常" unit="3"/>
    <Item patroldevice_name="设备E" patroldevice_code="44556" time="2025-01-09T12:20:00" type="210" value="1" value_unit="报警" unit="4"/>
    <Item patroldevice_name="设备F" patroldevice_code="78901" time="2025-01-09T12:25:00" type="414" value="1" value_unit="空闲状态" unit="5"/>
    <Item patroldevice_name="设备G" patroldevice_code="23456" time="2025-01-09T12:30:00" type="613" value="1" value_unit="任务模式" unit="6"/>
    <Item patroldevice_name="设备H" patroldevice_code="33445" time="2025-01-09T12:35:00" type="2012" value="3" value_unit="飞行中" unit="7"/>
    </Items>
</PatrolDevice>`

	xmlBizData1 = `<PatrolDevice>
    <SendCode>testDog</SendCode>
    <ReceiveCode>Server01</ReceiveCode>
    <Code>变电站编码</Code>
    <Type>64</Type>
    <Items>
        <Item task_patrolled_id="任务执行ID001" device_id="设备点位ID001,设备点位ID002" is_alarm="1" confirm_people="确认人A" confirm_date="2025-01-09 12:10:00"/>
        <Item task_patrolled_id="任务执行ID002" device_id="设备点位ID003" is_alarm="2" confirm_people="确认人B" confirm_date="2025-01-09 12:15:00"/>
        <Item task_patrolled_id="任务执行ID003" device_id="设备点位ID004" is_alarm="1" confirm_people="确认人C" confirm_date="2025-01-09 12:20:00"/>
        <Item task_patrolled_id="任务执行ID004" device_id="设备点位ID005,设备点位ID006" is_alarm="1" confirm_people="确认人D" confirm_date="2025-01-09 12:25:00"/>
        <Item task_patrolled_id="任务执行ID005" device_id="设备点位ID007" is_alarm="0" confirm_people="确认人E" confirm_date="2025-01-09 12:30:00"/>
        <Item task_patrolled_id="任务执行ID006" device_id="设备点位ID008,设备点位ID009,设备点位ID010" is_alarm="2" confirm_people="确认人F" confirm_date="2025-01-09 12:35:00"/>
        <Item task_patrolled_id="任务执行ID007" device_id="设备点位ID011" is_alarm="1" confirm_people="确认人G" confirm_date="2025-01-09 12:40:00"/>
    </Items>
</PatrolDevice>`

	xmlBizData3 = `<PatrolDevice>
    <SendCode>testDog</SendCode>
    <ReceiveCode>Server01</ReceiveCode>
    <Code>变电站编码</Code>
    <Type>11</Type>
    <Time>2025-01-09 12:00:00</Time>
    <Items>
        <Item time="2025-01-09 12:00:00" type="1" file_path="/path/to/deviceA_model.xml"/>
        <Item time="2025-01-09 12:05:00" type="2" file_path="/path/to/region_host_model.xml"/>
        <Item time="2025-01-09 12:10:00" type="3" file_path="/path/to/robot_model.xml"/>
        <Item time="2025-01-09 12:15:00" type="4" file_path="/path/to/camera_model.xml"/>
        <Item time="2025-01-09 12:20:00" type="5" file_path="/path/to/drone_model.xml"/>
        <Item time="2025-01-09 12:25:00" type="6" file_path="/path/to/voice_model.xml"/>
        <Item time="2025-01-09 12:30:00" type="7" file_path="/path/to/task_model.xml"/>
        <Item time="2025-01-09 12:35:00" type="8" file_path="/path/to/maintenance_config.xml"/>
        <Item time="2025-01-09 12:40:00" type="9" file_path="/path/to/map_file.xml"/>
        <Item time="2025-01-09 12:45:00" type="10" file_path="/path/to/maintenance_record.xml"/>
        <Item time="2025-01-09 12:50:00" type="11" file_path="/path/to/linkage_config.xml"/>
        <Item time="2025-01-09 12:55:00" type="12" file_path="/path/to/alarm_threshold_model.xml"/>
    </Items>
</PatrolDevice>`

	xmlHeartData = `<?xml version="1.0" encoding="UTF-8"?>
<PatrolDevice>
    <SendCode>testDog</SendCode>
    <ReceiveCode>Server01</ReceiveCode>
    <Type>251</Type>
    <Code/>
    <Command>2</Command>
    <Time>2022-01-01 12:02:34</Time>
    <Items/>
</PatrolDevice>`

	xmlCallback2513 = `<?xml version="1.0" encoding="UTF-8"?>
<PatrolDevice>
    <SendCode>Server01</SendCode>
    <ReceiveCode>testDog</ReceiveCode>
    <Type>251</Type>
    <Code>xmlCallback2513</Code>
    <Command>3</Command>
    <Time>2022-01-01 12:02:34</Time>
    <Items><Item test="1" fly = "1" test1= "第一个"/><Item test="2" fly = "2" test1= "第二个"/><Item test="3" fly = "3" test1= "第三个"/></Items>
</PatrolDevice>`
	xmlCallback2514 = `<?xml version="1.0" encoding="UTF-8"?>
<PatrolDevice>
    <SendCode>Server01</SendCode>
    <ReceiveCode>testDog</ReceiveCode>
    <Type>251</Type>
    <Code>xmlCallback2514</Code>
    <Command>4</Command>
    <Time>2022-01-01 12:02:34</Time>
    <Items><Item test="1" fly = "1" test1= "第一个"/><Item test="2" fly = "2" test1= "第二个"/><Item test="3" fly = "3" test1= "第三个"/></Items>
</PatrolDevice>`
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
			"  XMLContent:    \n %s,\n"+
			"  EndFlag:       0x%04X\n"+
			"}",
		m.StartFlag, m.TransmitSeq, m.ReceiveSeq, m.SessionSource, m.XMLLength, m.XMLContent, m.EndFlag)
}

// 转换为小端字节序
func toBytes(v interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// 转换为大端字节序
func toBytesBig(v interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// 构造并写入消息内容
func writeBuffer(msg Message, buf *bytes.Buffer) {
	if startFlagBytes, err := toBytesBig(msg.StartFlag); err == nil {
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
	if endFlagBytes, err := toBytesBig(msg.EndFlag); err == nil {
		buf.Write(endFlagBytes)
	}
}

// 客户端事件处理器
type clientEventHandler struct {
	*gnet.BuiltinEventEngine
	con        net.Conn
	ReceiveSeq int64
}

var seq int64 // 定义一个 int64 类型的计数器

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
		TransmitSeq:   seq,
		ReceiveSeq:    0,
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

func callback(msg Message, con gnet.Conn, xml string) {
	// 构造消息
	call := Message{
		StartFlag:     startFlag,
		TransmitSeq:   msg.TransmitSeq,
		ReceiveSeq:    msg.TransmitSeq,
		SessionSource: 0x00,
		XMLLength:     int32(len(xml)),
		XMLContent:    xml,
		EndFlag:       endFlag,
	}
	// 构造字节流
	buf := new(bytes.Buffer)
	writeBuffer(call, buf)
	hexStr := hex.EncodeToString(buf.Bytes())
	fmt.Printf("callback: %s\n", hexStr)
	con.Write(buf.Bytes())
}

func (c *clientEventHandler) OnClose(_ gnet.Conn, _ error) (action gnet.Action) {
	fmt.Println("OnClose")
	return gnet.Shutdown
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
			// 判断是否是 EOF 错误
			if err == io.EOF {
				// 读取完成，没有更多数据
				fmt.Println("End of file (EOF) reached.")
				break // 或者执行其它逻辑处理
			}
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

				if strutil.ContainsAny(msg.XMLContent, []string{"<Type>1</Type>", "<Type>11</Type>", "<Type>64</Type>"}) {
					callback(msg, conn, xmlCallback2513)
				} else if strutil.ContainsAny(msg.XMLContent, []string{"<Type>41</Type>"}) {
					callback(msg, conn, xmlCallback2514)
				}

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
	if len(data) < 25 {
		return fmt.Errorf("invalid message length")
	}

	// 解析 StartFlag (大端字节序)
	msg.StartFlag = binary.BigEndian.Uint16(data[:2])

	// 解析 TransmitSeq (小端字节序)
	msg.TransmitSeq = int64(binary.LittleEndian.Uint64(data[2:10]))

	// 解析 ReceiveSeq (小端字节序)
	msg.ReceiveSeq = int64(binary.LittleEndian.Uint64(data[10:18]))

	// 解析 SessionSource
	msg.SessionSource = data[18]

	// 解析 XMLLength（4 字节，小端字节序）
	msg.XMLLength = int32(binary.LittleEndian.Uint32(data[19:23]))

	// 检查 XMLLength 是否大于 0，如果是，解析 XMLContent
	if msg.XMLLength > 0 {
		// 解析 XMLContent
		xmlContentStart := 23
		xmlContentEnd := xmlContentStart + int(msg.XMLLength)
		if len(data) < xmlContentEnd {
			return fmt.Errorf("invalid XML length")
		}
		msg.XMLContent = string(data[xmlContentStart:xmlContentEnd])

		// 解析 EndFlag (小端字节序)
		msg.EndFlag = binary.BigEndian.Uint16(data[xmlContentEnd : xmlContentEnd+2])
	} else {
		// 如果 XMLLength 为 0，说明没有 XML 内容，直接解析 EndFlag
		msg.EndFlag = binary.BigEndian.Uint16(data[23:25])
	}
	return nil
}

func (c *clientEventHandler) OnTick() (delay time.Duration, action gnet.Action) {
	fmt.Println("OnTick")
	delay = 10 * time.Second
	// 构造消息
	msg := Message{
		StartFlag:     startFlag,
		TransmitSeq:   atomic.AddInt64(&seq, 1),
		ReceiveSeq:    0,
		SessionSource: 0x00,
		XMLLength:     int32(len(xmlHeartData)),
		XMLContent:    xmlHeartData,
		EndFlag:       endFlag,
	}
	fmt.Printf("send 心跳: %d\n", seq)
	if c.con != nil {
		// 构造字节流
		buf := new(bytes.Buffer)
		writeBuffer(msg, buf)
		hexStr := hex.EncodeToString(buf.Bytes())
		fmt.Printf("send 心跳: %s\n", hexStr)
		c.con.Write(buf.Bytes())

		// 构建 biz
		msgBiz := Message{
			StartFlag:     startFlag,
			TransmitSeq:   atomic.AddInt64(&seq, 1),
			ReceiveSeq:    0,
			SessionSource: 0x00,
			XMLLength:     int32(len(xmlBizData3)),
			XMLContent:    xmlBizData3,
			EndFlag:       endFlag,
		}
		// 构造字节流
		buf1 := new(bytes.Buffer)
		writeBuffer(msgBiz, buf1)
		c.con.Write(buf1.Bytes())
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
	wg.Add(1)
	conn, err := cli.Dial("tcp", "127.0.0.1:7100")
	if err != nil {
		log.Fatalf("Failed to dial gnet client: %v", err)
	}
	clientEV.con = conn
	wg.Wait()
}
