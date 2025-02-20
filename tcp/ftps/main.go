package main

import (
	"crypto/tls"
	"fmt"
	"github.com/jlaffaye/ftp"
	"log"
	"os"
)

func main() {
	var opts []ftp.DialOption
	// FTPS服务器配置
	server := "10.10.1.213:10012" // 显式FTPS通常使用21端口
	username := "test"
	password := "123qwe,."

	// TLS配置（根据服务器证书情况调整）
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // 跳过证书验证（仅用于测试，生产环境应设为false）
		//ServerName:         "ftp.example.com",
		ClientSessionCache: tls.NewLRUClientSessionCache(1), // 启用会话缓存
		VerifyConnection: func(state tls.ConnectionState) error {
			fmt.Printf("TLS Session Resumed: %v\n", state.DidResume)
			return nil
		},
		SessionTicketsDisabled: false,
	}

	//opts = append(opts, ftp.DialWithExplicitTLS(tlsConfig))
	opts = append(opts, ftp.DialWithTLS(tlsConfig)) // 使用隐式TLS连接
	opts = append(opts, ftp.DialWithDebugOutput(os.Stdout))
	opts = append(opts, ftp.DialWithDisabledEPSV(true))

	// 创建带显式TLS的FTP客户端
	client, err := ftp.Dial(server, opts...)
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	defer client.Quit()

	// 登录
	if err := client.Login(username, password); err != nil {
		log.Fatal("登录失败:", err)
	}
	fmt.Println("成功连接到FTPS服务器")

	// 示例：列出根目录文件
	fmt.Println("\n根目录文件列表:")
	entries, err := client.List("/home/test")
	if err != nil {
		log.Fatal("列出目录失败:", err)
	}

	for _, entry := range entries {
		fmt.Println(entry.Name)
	}

	// 示例：上传文件
	file, err := os.Open("localfile.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("本地文件不存在:", err)
		}
		log.Fatal("打开文件失败:", err)
	}
	defer file.Close()

	err = client.Stor("remotefile.txt", file)
	if err != nil {
		log.Fatal("上传失败:", err)
	}
	fmt.Println("\n文件上传成功")
}
