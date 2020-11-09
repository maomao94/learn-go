package persist

import (
	"context"
	"encoding/json"
	"learn-go/crawler/model"
	"testing"

	"github.com/olivere/elastic/v7"
)

func TestSave(t *testing.T) {
	expected := model.Profile{
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
	}

	id, err := save(expected)

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
		Type("zhenai").
		Id(id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)
	var actual model.Profile
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	if actual != expected {
		t.Errorf("expected %v;but was %v", expected, actual)
	}
}
