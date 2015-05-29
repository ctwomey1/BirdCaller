package main

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

type processor interface {
	Search()
	Parse()
}

type BirdInterpreter struct {
	API            anaconda.TwitterApi
	ScreenNames    []string
	SearchCriteria string
	Timeline       []anaconda.Tweet
	Cage           BirdCage
}

func (bi *BirdInterpreter) Search() {
	var err error
	v := url.Values{}
	for _, screenName := range bi.ScreenNames {
		v.Set("screen_name", screenName)
		bi.Timeline, err = bi.API.GetUserTimeline(v)
		if err != nil {
			panic(err)
		}

	}
}

func (bi *BirdInterpreter) Parse() {
	timeline := SearchText(bi.Timeline, bi.SearchCriteria)
	err := bi.Cage.PutInCage(timeline)
	if err != nil {
		panic(err)
	}
}
