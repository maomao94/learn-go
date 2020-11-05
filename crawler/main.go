package main

import (
	"bufio"
	"fmt"
	"io"
	"learn-go/crawler/engine"
	"learn-go/crawler/zhenai/parser"
	"regexp"

	"golang.org/x/text/encoding"

	"golang.org/x/net/html/charset"
)

func main() {
	//resp, err := http.Get("http://localhost:8080/mock/www.zhenai.com/zhenghun")
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer resp.Body.Close()
	//
	//if resp.StatusCode != http.StatusOK {
	//	fmt.Println("Error: status code", resp.StatusCode)
	//	return
	//}
	//
	//e := determineEncoding(resp.Body)
	//
	//utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	//all, err := ioutil.ReadAll(utf8Reader)
	//if err != nil {
	//	panic(err)
	//}
	//
	//printCityList(all)
	//fmt.Printf("%s\n", all)

	engine.Run(engine.Request{
		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
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
