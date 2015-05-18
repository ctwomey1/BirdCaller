package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

type Configuration struct {
	ConsumerKey    string `json:"consumer-key"`
	ConsumerSecret string `json:"consumer-secret"`
	AccessToken    string `json:"access-token"`
	AccessSecrect  string `json:"access-secret"`
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
	fmt.Println(config)
}
func main() {
	anaconda.SetConsumerKey(config.ConsumerKey)
	anaconda.SetConsumerSecret(config.ConsumerSecret)
	api := anaconda.NewTwitterApi(config.AccessToken, config.AccessSecrect)
	v := url.Values{}
	v.Set("screen_name", "codingvelocity")
	timeline, err := api.GetUserTimeline(v)
	if err != nil {
		panic(err)
	}

	fmt.Print(timeline)
}
