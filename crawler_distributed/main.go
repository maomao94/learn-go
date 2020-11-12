package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"learn-go/crawler/engine"
	"learn-go/crawler/scheduler"
	"learn-go/crawler/zhenai/parser"
	"learn-go/crawler_distributed/config"
	itemsaver "learn-go/crawler_distributed/persist/client"
	"learn-go/crawler_distributed/rpcsupport"
	worker "learn-go/crawler_distributed/worker/client"
	"log"
	"net/rpc"
	"regexp"
	"strings"

	"golang.org/x/text/encoding"

	"golang.org/x/net/html/charset"
)

var (
	itemSaverHost = flag.String("itemsaver_hots", "", "itemsaver host")

	workerHost = flag.String("worker_host", "",
		"worker_host")
)

func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(
		fmt.Sprintf(":%d", *itemSaverHost))
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHost, ","))
	processor := worker.CreateProcessor(pool)
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url: "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList, config.ParseCityList),
	})
}

func determineEncoding(
	r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(content []byte) {
	re := regexp.MustCompile(`<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(content, -1)
	for _, m := range matches {
		fmt.Printf("City: %s,URL: %s\n", m[2], m[1])
	}
	fmt.Printf("Matches found: %d\n", len(matches))
}

func createClientPool(host []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range host {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("error connecting to %s: %v",
				h, err)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
