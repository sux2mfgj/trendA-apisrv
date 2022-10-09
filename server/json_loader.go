package main

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"log"
	"os"
	"path"
	"strings"
)

type JsonLoader struct {
	Datapath string
}

var _ TrendsLoader = (*JsonLoader)(nil)

func NewJsonLoader(datapath string) (*JsonLoader, error) {
	info, err := os.Stat(datapath)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", datapath)
	}

	return &JsonLoader{
		Datapath: datapath,
	}, nil
}

func (j *JsonLoader) validateFilePath(path string) bool {
	return !strings.HasPrefix(path, j.Datapath)
}

func (j *JsonLoader) Load(location string, datetime string) ([]twitter.TrendsList, error) {
	filepath := path.Join(j.Datapath, location, datetime+".json")

	if !j.validateFilePath(filepath) {
		log.Println("query tried to access invalid path:", filepath)
		return nil, fmt.Errorf("Invalid access")
	}

	byteValue, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var trendlists []twitter.TrendsList

	err = json.Unmarshal(byteValue, &trendlists)
	if err != nil {
		return nil, err
	}

	return trendlists, nil
}
