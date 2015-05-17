package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	ConsumerKey    string `json:"consumer-key"`
	ConsumerSecret string `json:"consumer-secret"`
	AccessToken    string `json:"access-token"`
	AccessSecrect  string `json:"access-secret"`
}

func init() {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(configuration)
}
func main() {

}
