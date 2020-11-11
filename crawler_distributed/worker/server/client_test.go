package main

import (
	"fmt"
	"learn-go/crawler_distributed/config"
	"learn-go/crawler_distributed/rpcsupport"
	"learn-go/crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServerRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	request := worker.Request{
		Url: "http://localhost:8080/mock/album.zhenai.com/u/2219775402816308284",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "断念肉嘟嘟",
		},
	}

	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, request, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
