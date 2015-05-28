package processor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ChimeraCoder/anaconda"
)

type BirdCage interface {
	putInCage(timeline []anaconda.Tweet) error
}

type JSONFileCage struct {
	Filename string
}

func (j *JSONFileCage) putInCage(timeline []anaconda.Tweet) error {
	data, err := json.Marshal(timeline)
	if err == nil {
		ioutil.WriteFile(j.Filename, data, 644)
	} else {
		log.Fatalf("failed to create file %v, error %v", j.Filename, err)
	}
	return err
}

type ConsoleCage struct {
}

func (j *ConsoleCage) putInCage(timeline []anaconda.Tweet) error {
	fmt.Printf("%+v", timeline)
}
