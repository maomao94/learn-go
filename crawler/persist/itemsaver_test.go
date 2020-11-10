package persist

import (
	"context"
	"encoding/json"
	"learn-go/crawler/engine"
	"learn-go/crawler/model"
	"testing"

	"github.com/olivere/elastic/v7"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "http://localhost:8080/mock/album.zhenai.com/u/2219775402816308284",
		Type: "zhenai",
		Id:   "2219775402816308284",
		Payload: model.Profile{
			Name:       "断念肉嘟嘟",
			Gender:     "男",
			Age:        27,
			Height:     7,
			Weight:     57,
			Income:     "8001-10000元",
			Marriage:   "离异",
			Education:  "硕士",
			Occupation: "财务",
			Hokou:      "大连市",
			Xinzuo:     "天秤座",
			House:      "有房",
			Car:        "有车",
		},
	}

	err := save(expected)

	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().
		Index("dating_profile").
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)
	var actual engine.Item
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	if actual != expected {
		t.Errorf("expected %v;but was %v", expected, actual)
	}

	_, err = client.
		DeleteIndex("dating_profile").
		Do(context.Background())
	if err != nil {
		panic(nil)
	}
}
