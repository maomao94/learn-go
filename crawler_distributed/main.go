package main

import (
	"bufio"
	"fmt"
	"io"
	"learn-go/crawler/engine"
	"learn-go/crawler/scheduler"
	"learn-go/crawler/zhenai/parser"
	"learn-go/crawler_distributed/config"
	itemsaver "learn-go/crawler_distributed/persist/client"
	worker "learn-go/crawler_distributed/worker/client"
	"regexp"

	"golang.org/x/text/encoding"

	"golang.org/x/net/html/charset"
)

func main() {
	itemChan, err := itemsaver.ItemSaver(
		fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	processor, err := worker.CreateProcessor()
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
