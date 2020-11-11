package client

import (
	"fmt"
	"learn-go/crawler/engine"
	"learn-go/crawler_distributed/config"
	"learn-go/crawler_distributed/rpcsupport"
	"learn-go/crawler_distributed/worker"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(
		fmt.Sprintf(":%d", config.WorkPort0))
	if err != nil {
		return nil, err
	}

	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializedRequest(req)
		var sResult worker.ParseResult
		err = client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult)
	}, nil
}
