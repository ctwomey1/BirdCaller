package processor

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/ChimeraCoder/anaconda"
)

var unfilteredTimeline []anaconda.Tweet = []anaconda.Tweet{anaconda.Tweet{Text: "today i'm @ bills sandwhiches corner of main and 1st"}, anaconda.Tweet{Text: "today i ate some toast and cheese, it wasn't very good"}, anaconda.Tweet{Text: "I smashed my elbow on the corner of a counter, ouchie"}, anaconda.Tweet{Text: "hanging out at the corner of 1st and main"}}

func TestBirdInterpreterParsing(t *testing.T) {
	file, _ := ioutil.TempFile(os.TempDir(), "parsingTest.json")
	bi := BirdInterpreter{Timeline: unfilteredTimeline, SearchCriteria: "corner of .* and .*", Cage: &JSONFileCage{Filename: file.Name()}}
	bi.Parse()
	var jsonfile []anaconda.Tweet
	marshalledfile, _ := ioutil.ReadFile(file.Name())
	_ = json.Unmarshal(marshalledfile, &jsonfile)
	if len(jsonfile) != 2 {

		t.Errorf("Timeline search should return %v values but returned %v", 2, len(jsonfile))
	}

}
