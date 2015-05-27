package processor

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

type processor interface {
	Search()
	Parse()
	Save()
}

type BirdInterpreter struct {
	API            anaconda.TwitterApi
	ScreenNames    []string
	SearchCriteria string
	Timeline       []anaconda.Tweet
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
	bi.Timeline = SearchText(bi.Timeline, bi.SearchCriteria)
}

func (bi *BirdInterpreter) Save() {
	//log.Println(SearchText(bi.Timeline, bi.SearchCriteria))
}
