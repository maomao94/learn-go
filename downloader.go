package main

import (
	"fmt"
	"io/ioutil"
	"learn-go/retriever/testing"
	"net/http"
)

func getRetriever() retriever {
	return testing.Retriever{}
}

func retrieve(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}

//something that can "Get"
type retriever interface {
	Get(string) string
}

func main() {
	//retriever := infra.Retriever{}
	retriever := getRetriever()
	//fmt.Println(retrieve("http://www.baidu.com"))
	fmt.Println(retriever.Get("http://www.baidu.com"))
}
