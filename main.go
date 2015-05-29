package main

import (
	"encoding/json"
	"flag"
	"log"
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
var jsonFile *string
var filename *string
var jsonOut *bool

func init() {
	searchCriteria = flag.String("search", "", "Search criteria to use when searching tweets")
	screenNames = flag.String("screenNames", "", "Comma delimited list of twitter screen names to use when searching")
	jsonOut = flag.Bool("json", false, "Wether or not ")
	filename = flag.String("file", "timeline.json", "json file to save timeline to, defaults to timeline.json")
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
	var cage BirdCage
	if *jsonOut {
		cage = &JSONFileCage{Filename: *filename}
	}
	bi := BirdInterpreter{ScreenNames: strings.Split(*screenNames, ","), SearchCriteria: *searchCriteria, API: *api, Cage: cage}
	bi.Search()
	bi.Parse()

}
