package client

import (
	"learn-go/crawler/engine"
	"learn-go/crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan<- engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item "+
				"#%d: %v", itemCount, item)
			itemCount++

			//call RPC to save item
			result := ""
			client.Call("ItemSaverService.Save",
				item, &result)
			if err != nil {
				log.Printf("Item Saver: error "+
					"saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}
