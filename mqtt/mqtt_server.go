package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"time"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	s := make(chan os.Signal, 1)
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("emqx_go_client")

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅主题
	if token := c.Subscribe("test/#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	go func() {
		// 发布消息
		c.Publish("test/1", 0, false, "Hello World")
		c.Publish("test/2", 0, false, "Hello World")
		c.Publish("test/3", 0, false, "Hello World")
		c.Publish("test/4", 0, false, "Hello World")
		c.Publish("test/5", 0, false, "Hello World")
	}()
	<-s
	fmt.Println("信号：", s)

	//// 发布消息
	//token := c.Publish("test/1", 0, false, "Hello World")
	//token.Wait()
	//
	//time.Sleep(6 * time.Second)
	//
	//// 取消订阅
	//if token := c.Unsubscribe("testtopic/#"); token.Wait() && token.Error() != nil {
	//	fmt.Println(token.Error())
	//	os.Exit(1)
	//}
	//
	//// 断开连接
	//c.Disconnect(250)
	//time.Sleep(1 * time.Second)
}
