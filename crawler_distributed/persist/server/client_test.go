package main

import (
	"learn-go/crawler/engine"
	"learn-go/crawler/model"
	"learn-go/crawler_distributed/config"
	"learn-go/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	go serverRpc(host, "dating_profile")
	time.Sleep(time.Second)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	item := engine.Item{
		Url:  "http://localhost:8080/mock/album.zhenai.com/u/2219775402816308284",
		Type: "zhenai",
		Id:   "2219775402816308284",
		Payload: model.Profile{
			Name:       "断念肉嘟嘟",
			Gender:     "男",
			Age:        27,
			Height:     7,
			Weight:     57,
			Income:     "8001-10000元",
			Marriage:   "离异",
			Education:  "硕士",
			Occupation: "财务",
			Hokou:      "大连市",
			Xinzuo:     "天秤座",
			House:      "有房",
			Car:        "有车",
		},
	}
	result := ""
	client.Call(config.ItemSaverRpc, item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
