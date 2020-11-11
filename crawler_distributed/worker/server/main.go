package main

import (
	"fmt"
	"learn-go/crawler_distributed/config"
	"learn-go/crawler_distributed/rpcsupport"
	"learn-go/crawler_distributed/worker"
)

func main() {
	rpcsupport.ServerRpc(
		fmt.Sprintf(":%d", config.WorkPort0),
		worker.CrawlService{})
}
