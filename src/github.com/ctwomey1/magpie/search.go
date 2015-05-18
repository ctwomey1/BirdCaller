package main

import (
	"log"
	"regexp"

	"github.com/ChimeraCoder/anaconda"
)

func searchTimeline(timeline []anaconda.Tweet, searchPattern string) []anaconda.Tweet {
	var matchingTweets []anaconda.Tweet
	regex, err := regexp.Compile(searchPattern)
	if err != nil {
		log.Printf("Regular expression %v could not compile because %v", searchPattern, err)
	} else {
		for i := 0; i < len(timeline); i++ {
			match := regex.MatchString(timeline[i].Text)
			if match {
				matchingTweets = append(matchingTweets, timeline[i])
			}
		}
	}
	return matchingTweets
}
