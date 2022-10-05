package main

import (
	"encoding/json"
	"github.com/dghubble/go-twitter/twitter"
	"log"
	"os"
)

func Save(trend *twitter.Trend) {

	j, err := json.Marshal(trend)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile("dump.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := file.Write([]byte(string(j))); err != nil {
		log.Fatal(err)
	}

	if err = file.Close(); err != nil {
		log.Fatal(err)
	}
}
