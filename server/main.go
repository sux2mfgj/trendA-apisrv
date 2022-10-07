package main

import (
	"encoding/json"
	"github.com/dghubble/go-twitter/twitter"
	"log"
	"net/http"
)

type Loader interface {
	Load(string, string) ([]twitter.TrendsList, error)
}

func GetTrends(w http.ResponseWriter, r *http.Request) {
	d := DummyLoader{}
	trends, err := d.Load("", "")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(trends)
	if err != nil {
		log.Fatal(err)
	}
}

type Version struct {
	Version string `json:"version"`
}

func VersionHandler(w http.ResponseWriter, _ *http.Request) {
	v := Version{
		Version: "0.0.1",
	}

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/v0/version", VersionHandler)
	http.HandleFunc("/v0/trends", GetTrends)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
