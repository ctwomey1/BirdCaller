package processor

import (
	"testing"

	"github.com/ChimeraCoder/anaconda"
)

var unfilteredTimeline []anaconda.Tweet = []anaconda.Tweet{anaconda.Tweet{Text: "today i'm @ bills sandwhiches corner of main and 1st"}, anaconda.Tweet{Text: "today i ate some toast and cheese, it wasn't very good"}, anaconda.Tweet{Text: "I smashed my elbow on the corner of a counter, ouchie"}, anaconda.Tweet{Text: "hanging out at the corner of 1st and main"}}

func TestBirdInterpreterParsing(t *testing.T) {
	bi := BirdInterpreter{Timeline: unfilteredTimeline, SearchCriteria: "corner of .* and .*"}
	bi.Parse()

	if len(bi.Timeline) != 2 {

		t.Errorf("Timeline search should return %v values but returned %v", 2, len(bi.Timeline))
	}

}

func TestBirdInterpreterSaving(t *testing.T) {
	bi := BirdInterpreter{Timeline: unfilteredTimeline}

	bi.Save()
}
