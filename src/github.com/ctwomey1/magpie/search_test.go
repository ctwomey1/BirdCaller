package main

import (
	"testing"

	"github.com/ChimeraCoder/anaconda"
)

func TestBasicRegexCompileAndMatch(t *testing.T) {
	timeline := []anaconda.Tweet{anaconda.Tweet{Text: "test"}}
	m := searchTimeline(timeline, "test")
	if m == nil {
		t.Error("no matches found")
	}
	if len(m) != 1 {
		t.Error("found too many matches!")
	}

}
