package persist

import (
	"learn-go/crawler/engine"
	"learn-go/crawler/persist"

	"github.com/olivere/elastic/v7"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(
	item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	if err == nil {
		*result = "ok"
	}
	return err
}
