package main

import (
	"encoding/json"
	"github.com/dghubble/go-twitter/twitter"
	"log"
	"os"
	"path"
	"time"
)

func StoreTrend(trend *twitter.Trend, dir string, location string) (string, error) {
	now := time.Now()
	const layout = "2006_01_02_15.json"

	base := path.Join(dir, location)

	if err := os.MkdirAll(base, os.ModePerm); err != nil {
		log.Fatal(err)
		return "", err
	}

	filepath := path.Join(base, now.Format(layout))

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	j, err := json.Marshal(trend)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	if _, err := file.Write([]byte(string(j))); err != nil {
		log.Fatal(err)
		return "", err
	}

	if err = file.Close(); err != nil {
		log.Fatal(err)
		return "", err
	}

	return filepath, nil
}
