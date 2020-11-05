package parser

import (
	"fmt"
	"learn-go/crawler/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetch("http://localhost:8080/mock/www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", contents)
	ParseCityList(contents)
}
