package main

import (
	"fmt"
	"learn-go/crawler_distributed/config"
	"learn-go/crawler_distributed/persist"
	"learn-go/crawler_distributed/rpcsupport"

	"github.com/olivere/elastic/v7"
)

func main() {
	serverRpc(
		fmt.Sprintf(":%d", config.ItemSaverPort),
		config.ElasticIndex)
}

func serverRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return rpcsupport.ServerRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
