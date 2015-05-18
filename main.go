package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

type SearchCriteriaStruct struct {
	ScreenNames   []string `json:"screen_names"`
	SearchPattern string   `json:"searchPattern"`
}
type Configuration struct {
	ConsumerKey    string               `json:"consumer-key"`
	ConsumerSecret string               `json:"consumer-secret"`
	AccessToken    string               `json:"access-token"`
	AccessSecrect  string               `json:"access-secret"`
	SearchCriteria SearchCriteriaStruct `json:"searchCriteria"`
}

var config Configuration

func init() {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("Config loaded with the following values %+v\n", &config)
}

func main() {
	anaconda.SetConsumerKey(config.ConsumerKey)
	anaconda.SetConsumerSecret(config.ConsumerSecret)
	api := anaconda.NewTwitterApi(config.AccessToken, config.AccessSecrect)
	v := url.Values{}
	for _, screenName := range config.SearchCriteria.ScreenNames {
		v.Set("screen_name", screenName)
		timeline, err := api.GetUserTimeline(v)
		if err != nil {
			panic(err)
		}
		fmt.Println(SearchText(timeline, config.SearchCriteria.SearchPattern))
	}
}
