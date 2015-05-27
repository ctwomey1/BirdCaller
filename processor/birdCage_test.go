package processor

import (
	"os"
	"testing"

	"github.com/ChimeraCoder/anaconda"
)

var birdsToCage []anaconda.Tweet = []anaconda.Tweet{anaconda.Tweet{Text: "today i'm @ bills sandwhiches corner of main and 1st"}, anaconda.Tweet{Text: "today i ate some toast and cheese, it wasn't very good"}, anaconda.Tweet{Text: "I smashed my elbow on the corner of a counter, ouchie"}, anaconda.Tweet{Text: "hanging out at the corner of 1st and main"}}

func TestPuttingInCage(t *testing.T) {
	birdcage := JSONFileCage{Filename: "tmp/test.json"}
	birdcage.putInCage(&birdsToCage)
	if _, err := os.Stat("/tmp/test.json"); os.IsNotExist(err) {
		t.Error("file doesn't exist")

	}
	os.Remove("tmp/test.json")

}
