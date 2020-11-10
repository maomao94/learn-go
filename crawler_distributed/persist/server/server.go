package main

import (
	"learn-go/crawler_distributed/persist"
	"learn-go/crawler_distributed/rpcsupport"

	"github.com/olivere/elastic/v7"
)

func main() {
	serverRpc(":1234", "dating_profile")
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
