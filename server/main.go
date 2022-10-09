package main

import (
	"encoding/json"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/jessevdk/go-flags"
	"log"
	"net/http"
	"strconv"
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

var opts struct {
	Config string `short:"c" description:"config file path"`
}

func main() {
	parser := flags.NewParser(&opts, flags.Default)

	_, err := parser.Parse()
	if err != nil {
		switch err.(type) {
		case flags.ErrorType:
			log.Fatal("a")
		default:
			log.Fatal("b")
		}
	}

	config, err := LoadConfigFile(opts.Config)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/v0/version", VersionHandler)
	http.HandleFunc("/v0/trends", GetTrends)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port), nil))
}
