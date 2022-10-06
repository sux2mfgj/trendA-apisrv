package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type TrendsLoader struct {
	Client *twitter.Client
}

const TwitterTokenURL = "https://api.twitter.com/oauth2/token"

func NewLoader(clientID string, clientSecret string) *TrendsLoader {
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     TwitterTokenURL,
	}

	httpClient := config.Client(oauth2.NoContext)

	client := twitter.NewClient(httpClient)

	return &TrendsLoader{
		Client: client,
	}
}

func (loader *TrendsLoader) LoadTrends(place int64) ([]twitter.TrendsList, error) {

	trends, _, err := loader.Client.Trends.Place(place, &twitter.TrendsPlaceParams{})
	if err != nil {
		return nil, err
	}

	return trends, err
}
