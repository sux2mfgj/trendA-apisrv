package main

// OAuth2
import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"log"
	"os"
)

func main() {

	clientID := os.Getenv("TRENDYA_TWITTER_CLIENT_ID")
	clientSecret := os.Getenv("TRENDYA_TWITTER_CLIENT_SECRET")

	if len(clientID) == 0 || len(clientSecret) == 0 {
		log.Fatal("Should be set environment variables named `TRENDYA_TWITTER_CLIENT_ID` and `TRENDYA_TWITTER_CLIENT_SECRET`.")
	}

	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	httpClient := config.Client(oauth2.NoContext)

	client := twitter.NewClient(httpClient)

	trend, _, err := client.Trends.Place(23424856, &twitter.TrendsPlaceParams{})
	if err != nil {
		fmt.Println("err {}", err)
	}

	for _, tweet := range trend {
		for _, t := range tweet.Trends {
			Save(&t)
			fmt.Println(t.Name)
		}
	}
}
