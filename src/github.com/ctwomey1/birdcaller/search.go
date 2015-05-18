package main

import (
	"log"
	"regexp"

	"github.com/ChimeraCoder/anaconda"
)

func SearchText(timeline []anaconda.Tweet, searchPattern string) []anaconda.Tweet {
	c := make(chan anaconda.Tweet)
	go func() {
		regex, err := regexp.Compile(searchPattern)
		if err != nil {
			log.Printf("Regular expression %v could not compile because %v", searchPattern, err)
		} else {
			for i := 0; i < len(timeline); i++ {
				match := regex.MatchString(timeline[i].Text)
				if match {
					c <- timeline[i]
				}
			}
		}
		close(c)
	}()
	var tweets []anaconda.Tweet
	for tweet := range c {
		tweets = append(tweets, tweet)
	}
	return tweets
}
