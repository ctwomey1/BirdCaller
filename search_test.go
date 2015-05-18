package main

import (
	"testing"

	"github.com/ChimeraCoder/anaconda"
)

func TestBasicRegexCompileAndMatch(t *testing.T) {
	timeline := []anaconda.Tweet{anaconda.Tweet{Text: "test"}}
	m := SearchText(timeline, "test")
	if m == nil {
		t.Error("no matches found")
	}

	if len(m) != 1 {
		t.Error("found too many matches!")
	}

}

func TestAdvancedRegex(t *testing.T) {
	timeline := []anaconda.Tweet{anaconda.Tweet{Text: "corner of Broadway and hargrave"}, anaconda.Tweet{Text: "today i ate some toast and cheese, it wasn't very good"}, anaconda.Tweet{Text: "I smashed my elbow on the corner of a counter, ouchie"}, anaconda.Tweet{Text: "hanging out at the corner of 1st and main"}}
	m := SearchText(timeline, "corner.*and")
	if m == nil {
		t.Error("no matches found")
	}

	if len(m) != 2 {
		t.Errorf("Found %v matches instead of 2", len(m))
	}
}
