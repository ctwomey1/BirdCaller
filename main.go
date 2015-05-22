package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
)

type Configuration struct {
	ConsumerKey    string `json:"consumer-key"`
	ConsumerSecret string `json:"consumer-secret"`
	AccessToken    string `json:"access-token"`
	AccessSecrect  string `json:"access-secret"`
}

var config Configuration
var searchCriteria *string
var screenNames *string

func init() {
	searchCriteria = flag.String("search", "", "Search criteria to use when searching tweets")
	screenNames = flag.String("screenNames", "", "Comma delimited list of twitter screen names to use when searching")
	file, err := os.Open("conf.json")
	if err != nil {
		log.Fatalln("Error opening conf.json!", err)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("error:", err)
	}

	log.Printf("Config loaded with the following values %+v\n", &config)
}

func main() {

	flag.Parse()

	anaconda.SetConsumerKey(config.ConsumerKey)
	anaconda.SetConsumerSecret(config.ConsumerSecret)
	api := anaconda.NewTwitterApi(config.AccessToken, config.AccessSecrect)
	v := url.Values{}
	for _, screenName := range strings.Split(*screenNames, ",") {
		v.Set("screen_name", screenName)
		timeline, err := api.GetUserTimeline(v)
		if err != nil {
			panic(err)
		}
		fmt.Println(SearchText(timeline, *searchCriteria))
	}
}
