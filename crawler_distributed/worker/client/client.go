package client

import (
	"learn-go/crawler/engine"
	"learn-go/crawler_distributed/config"
	"learn-go/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientsChan chan *rpc.Client) engine.Processor {
	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializedRequest(req)
		var sResult worker.ParseResult
		c := <-clientsChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult)
	}
}
