package main

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/jessevdk/go-flags"
	"log"
	"net/http"
	"strconv"
)

type TrendsLoader interface {
	Load(string, string) ([]twitter.TrendsList, error)
}

type Version struct {
	Version string `json:"version"`
}

type Server struct {
	Loader TrendsLoader
}

func (s *Server) versionHandler(w http.ResponseWriter, _ *http.Request) {
	v := Version{
		Version: "0.0.1",
	}

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) getTrends(w http.ResponseWriter, r *http.Request) {
	loc := r.URL.Query().Get("location")
	date := r.URL.Query().Get("date")

	if len(loc) == 0 || len(date) == 0 {
		log.Fatal(fmt.Errorf("invalid argument"))
	}

	trends, err := s.Loader.Load(loc, date)
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

	var loader TrendsLoader
	switch config.Loader {
	case "dummy":
		loader = &DummyLoader{}
	case "json":
		loader, err = NewJsonLoader(config.DataPath)
		if err != nil {
			log.Fatal("invalid argument")
		}
	default:
		log.Fatal("invalid loader")
	}

	srv := Server{
		Loader: loader,
	}

	http.HandleFunc("/v0/version", srv.versionHandler)
	http.HandleFunc("/v0/trends", srv.getTrends)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port), nil))
}
