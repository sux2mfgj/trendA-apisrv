package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"testing"
)

func TestSaveToFile(t *testing.T) {
	trend := twitter.Trend{
		Name:            "#天皇杯",
		URL:             "http://twitter.com/search?q=%23%E5%A4%A9%E7%9A%87%E6%9D%AF",
		PromotedContent: "",
		Query:           "%23%E5%A4%A9%E7%9A%87%E6%9D%AF",
		TweetVolume:     14438,
	}

	Save(&trend)
}
