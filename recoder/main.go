package main

import (
	"log"
	"os"
)

var japan_woeids = map[string]int64{
	"Kitakyushu": 1110809,
	"Saitama":    1116753,
	"Chiba":      1117034,
	"Fukuoka":    1117099,
	"Hamamatsu":  1117155,
	"Hiroshima":  1117227,
	"Kawasaki":   1117502,
	"Kobe":       1117545,
	"Kumamoto":   1117605,
	"Nagoya":     1117817,
	"Niigata":    1117881,
	"Sagamihara": 1118072,
	"Sapporo":    1118108,
	"Sendai":     1118129,
	"Takamatsu":  1118285,
	"Tokyo":      1118370,
	"Yokohama":   1118550,
	"Okinawa":    2345896,
	"Osaka":      15015370,
	"Kyoto":      15015372,
	"Okayama":    90036018,
	"Japan":      23424856,
}

func main() {

	clientID := os.Getenv("TRENDYA_TWITTER_CLIENT_ID")
	clientSecret := os.Getenv("TRENDYA_TWITTER_CLIENT_SECRET")

	if len(clientID) == 0 || len(clientSecret) == 0 {
		log.Fatal("Should be set environment variables named `TRENDYA_TWITTER_CLIENT_ID` and `TRENDYA_TWITTER_CLIENT_SECRET`.")
		return
	}

	datapath := os.Getenv("TRENDYA_DATAPATH")
	if len(datapath) == 0 {
		datapath = "./data"
	}

	loader := NewLoader(clientID, clientSecret)

	for place, id := range japan_woeids {
		trendList, err := loader.LoadTrends(id)
		if err != nil {
			log.Fatal(err)
		}

		err = StoreTrends(trendList, datapath, place)
		if err != nil {
			log.Fatal(err)
		}
	}
}
