package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents)
	const resultSize = 470
	var expectedUrls = []string{
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/aba",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/akesu",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/alashanmeng",
	}
	var expectedCities = []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}
	if len(result.Request) != resultSize {
		t.Errorf("result should have %d "+
			"requests; but had %d", resultSize, len(result.Request))
	}
	for i, url := range expectedUrls {
		if result.Request[i].Url != url {
			t.Errorf("expected url #%d: %s; but "+
				"was %s", i, url, result.Request[i].Url)
		}
	}
	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected url #%d: %s; but "+
				"was %s", i, city, result.Items[i].(string))
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d "+
			"requests; but had %d", resultSize, len(result.Items))
	}
}
