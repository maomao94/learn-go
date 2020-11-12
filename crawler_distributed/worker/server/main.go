package main

import (
	"flag"
	"fmt"
	"learn-go/crawler_distributed/rpcsupport"
	"learn-go/crawler_distributed/worker"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	rpcsupport.ServerRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{})
}
