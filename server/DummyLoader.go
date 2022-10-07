package main

import (
	//     "log"
	"github.com/dghubble/go-twitter/twitter"
)

type DummyLoader struct {
}

func (d *DummyLoader) Load(location string, datetime string) ([]twitter.TrendsList, error) {
	return []twitter.TrendsList{
		{
			Trends: []twitter.Trend{
				{
					Name:            "#SixTONES_GoodLuck",
					URL:             "http://twitter.com/search?q=%23SixTONES_GoodLuck",
					PromotedContent: "",
					Query:           "%23SixTONES_GoodLuck",
					TweetVolume:     81337,
				},
				{
					Name:            "#酒のツマミになる話",
					URL:             "http://twitter.com/search?q=%23%E9%85%92%E3%81%AE%E3%83%84%E3%83%9E%E3%83%9F%E3%81%AB%E3%81%AA%E3%82%8B%E8%A9%B1",
					PromotedContent: "",
					Query:           "%23%E9%85%92%E3%81%AE%E3%83%84%E3%83%9E%E3%83%9F%E3%81%AB%E3%81%AA%E3%82%8B%E8%A9%B1",
					TweetVolume:     41791,
				},
				{
					Name:            "#Mステ",
					URL:             "http://twitter.com/search?q=%23M%E3%82%B9%E3%83%86",
					PromotedContent: "",
					Query:           "%23M%E3%82%B9%E3%83%86",
					TweetVolume:     24058,
				},
				{
					Name:            "#與那城奨",
					URL:             "http://twitter.com/search?q=%23%E8%88%87%E9%82%A3%E5%9F%8E%E5%A5%A8",
					PromotedContent: "",
					Query:           "%23%E8%88%87%E9%82%A3%E5%9F%8E%E5%A5%A8",
					TweetVolume:     60793,
				},
				{
					Name:            "#アダムスファミリー2",
					URL:             "http://twitter.com/search?q=%23%E3%82%A2%E3%83%80%E3%83%A0%E3%82%B9%E3%83%95%E3%82%A1%E3%83%9F%E3%83%AA%E3%83%BC2",
					PromotedContent: "",
					Query:           "%23%E3%82%A2%E3%83%80%E3%83%A0%E3%82%B9%E3%83%95%E3%82%A1%E3%83%9F%E3%83%AA%E3%83%BC2",
					TweetVolume:     17161,
				},
				{
					Name:            "ニャオハ",
					URL:             "http://twitter.com/search?q=%E3%83%8B%E3%83%A3%E3%82%AA%E3%83%8F",
					PromotedContent: "",
					Query:           "%E3%83%8B%E3%83%A3%E3%82%AA%E3%83%8F",
					TweetVolume:     0,
				},
				{
					Name:            "梁山泊のアサシン",
					URL:             "http://twitter.com/search?q=%E6%A2%81%E5%B1%B1%E6%B3%8A%E3%81%AE%E3%82%A2%E3%82%B5%E3%82%B7%E3%83%B3",
					PromotedContent: "",
					Query:           "%E6%A2%81%E5%B1%B1%E6%B3%8A%E3%81%AE%E3%82%A2%E3%82%B5%E3%82%B7%E3%83%B3",
					TweetVolume:     0,
				},
				{
					Name:            "ウェンズデー",
					URL:             "http://twitter.com/search?q=%E3%82%A6%E3%82%A7%E3%83%B3%E3%82%BA%E3%83%87%E3%83%BC",
					PromotedContent: "",
					Query:           "%E3%82%A6%E3%82%A7%E3%83%B3%E3%82%BA%E3%83%87%E3%83%BC",
					TweetVolume:     0,
				},
				{
					Name:            "チーム8",
					URL:             "http://twitter.com/search?q=%E3%83%81%E3%83%BC%E3%83%A08",
					PromotedContent: "",
					Query:           "%E3%83%81%E3%83%BC%E3%83%A08",
					TweetVolume:     19024,
				},
				{
					Name:            "次のフェス",
					URL:             "http://twitter.com/search?q=%E6%AC%A1%E3%81%AE%E3%83%95%E3%82%A7%E3%82%B9",
					PromotedContent: "",
					Query:           "%E6%AC%A1%E3%81%AE%E3%83%95%E3%82%A7%E3%82%B9",
					TweetVolume:     0,
				},
				{
					Name:            "水タイプ",
					URL:             "http://twitter.com/search?q=%E6%B0%B4%E3%82%BF%E3%82%A4%E3%83%97",
					PromotedContent: "",
					Query:           "%E6%B0%B4%E3%82%BF%E3%82%A4%E3%83%97",
					TweetVolume:     0,
				},
				{
					Name:            "ホゲータ",
					URL:             "http://twitter.com/search?q=%E3%83%9B%E3%82%B2%E3%83%BC%E3%82%BF",
					PromotedContent: "",
					Query:           "%E3%83%9B%E3%82%B2%E3%83%BC%E3%82%BF",
					TweetVolume:     0,
				},
				{
					Name:            "オリオン",
					URL:             "http://twitter.com/search?q=%E3%82%AA%E3%83%AA%E3%82%AA%E3%83%B3",
					PromotedContent: "",
					Query:           "%E3%82%AA%E3%83%AA%E3%82%AA%E3%83%B3",
					TweetVolume:     22295,
				},
				{
					Name:            "コラボフェス",
					URL:             "http://twitter.com/search?q=%E3%82%B3%E3%83%A9%E3%83%9C%E3%83%95%E3%82%A7%E3%82%B9",
					PromotedContent: "",
					Query:           "%E3%82%B3%E3%83%A9%E3%83%9C%E3%83%95%E3%82%A7%E3%82%B9",
					TweetVolume:     93655,
				},
				{
					Name:            "aiko",
					URL:             "http://twitter.com/search?q=aiko",
					PromotedContent: "",
					Query:           "aiko",
					TweetVolume:     22461,
				},
				{
					Name:            "So What",
					URL:             "http://twitter.com/search?q=%22So+What%22",
					PromotedContent: "",
					Query:           "%22So+What%22",
					TweetVolume:     719070,
				},
				{
					Name:            "テゴマス",
					URL:             "http://twitter.com/search?q=%E3%83%86%E3%82%B4%E3%83%9E%E3%82%B9",
					PromotedContent: "",
					Query:           "%E3%83%86%E3%82%B4%E3%83%9E%E3%82%B9",
					TweetVolume:     0,
				},
				{
					Name:            "エリちゃん",
					URL:             "http://twitter.com/search?q=%E3%82%A8%E3%83%AA%E3%81%A1%E3%82%83%E3%82%93",
					PromotedContent: "",
					Query:           "%E3%82%A8%E3%83%AA%E3%81%A1%E3%82%83%E3%82%93",
					TweetVolume:     0,
				},
				{
					Name:            "BE ORIGINAL",
					URL:             "http://twitter.com/search?q=%22BE+ORIGINAL%22",
					PromotedContent: "",
					Query:           "%22BE+ORIGINAL%22",
					TweetVolume:     43887,
				},
				{
					Name:            "マンタロー",
					URL:             "http://twitter.com/search?q=%E3%83%9E%E3%83%B3%E3%82%BF%E3%83%AD%E3%83%BC",
					PromotedContent: "",
					Query:           "%E3%83%9E%E3%83%B3%E3%82%BF%E3%83%AD%E3%83%BC",
					TweetVolume:     0,
				},
				{
					Name:            "みずタイプ",
					URL:             "http://twitter.com/search?q=%E3%81%BF%E3%81%9A%E3%82%BF%E3%82%A4%E3%83%97",
					PromotedContent: "",
					Query:           "%E3%81%BF%E3%81%9A%E3%82%BF%E3%82%A4%E3%83%97",
					TweetVolume:     81072,
				},
				{
					Name:            "ほのおタイプ",
					URL:             "http://twitter.com/search?q=%E3%81%BB%E3%81%AE%E3%81%8A%E3%82%BF%E3%82%A4%E3%83%97",
					PromotedContent: "",
					Query:           "%E3%81%BB%E3%81%AE%E3%81%8A%E3%82%BF%E3%82%A4%E3%83%97",
					TweetVolume:     0,
				},
				{
					Name:            "元2世信者",
					URL:             "http://twitter.com/search?q=%E5%85%832%E4%B8%96%E4%BF%A1%E8%80%85",
					PromotedContent: "",
					Query:           "%E5%85%832%E4%B8%96%E4%BF%A1%E8%80%85",
					TweetVolume:     25946,
				},
				{
					Name:            "田淵さん",
					URL:             "http://twitter.com/search?q=%E7%94%B0%E6%B7%B5%E3%81%95%E3%82%93",
					PromotedContent: "",
					Query:           "%E7%94%B0%E6%B7%B5%E3%81%95%E3%82%93",
					TweetVolume:     0,
				},
				{
					Name:            "北斗の拳",
					URL:             "http://twitter.com/search?q=%E5%8C%97%E6%96%97%E3%81%AE%E6%8B%B3",
					PromotedContent: "",
					Query:           "%E5%8C%97%E6%96%97%E3%81%AE%E6%8B%B3",
					TweetVolume:     0,
				},
				{
					Name:            "B.I.Shadow",
					URL:             "http://twitter.com/search?q=B.I.Shadow",
					PromotedContent: "",
					Query:           "B.I.Shadow",
					TweetVolume:     0,
				},
				{
					Name:            "丸の内サディスティック",
					URL:             "http://twitter.com/search?q=%E4%B8%B8%E3%81%AE%E5%86%85%E3%82%B5%E3%83%87%E3%82%A3%E3%82%B9%E3%83%86%E3%82%A3%E3%83%83%E3%82%AF",
					PromotedContent: "",
					Query:           "%E4%B8%B8%E3%81%AE%E5%86%85%E3%82%B5%E3%83%87%E3%82%A3%E3%82%B9%E3%83%86%E3%82%A3%E3%83%83%E3%82%AF",
					TweetVolume:     0,
				},
				{
					Name:            "ユニゾン",
					URL:             "http://twitter.com/search?q=%E3%83%A6%E3%83%8B%E3%82%BE%E3%83%B3",
					PromotedContent: "",
					Query:           "%E3%83%A6%E3%83%8B%E3%82%BE%E3%83%B3",
					TweetVolume:     0,
				},
				{
					Name:            "雇用統計",
					URL:             "http://twitter.com/search?q=%E9%9B%87%E7%94%A8%E7%B5%B1%E8%A8%88",
					PromotedContent: "",
					Query:           "%E9%9B%87%E7%94%A8%E7%B5%B1%E8%A8%88",
					TweetVolume:     21793,
				},
				{
					Name:            "LINE流出",
					URL:             "http://twitter.com/search?q=LINE%E6%B5%81%E5%87%BA",
					PromotedContent: "",
					Query:           "LINE%E6%B5%81%E5%87%BA",
					TweetVolume:     0,
				},
				{
					Name:            "UNISON SQUARE GARDEN",
					URL:             "http://twitter.com/search?q=%22UNISON+SQUARE+GARDEN%22",
					PromotedContent: "",
					Query:           "%22UNISON+SQUARE+GARDEN%22",
					TweetVolume:     15067,
				},
				{
					Name:            "ほのお一択",
					URL:             "http://twitter.com/search?q=%E3%81%BB%E3%81%AE%E3%81%8A%E4%B8%80%E6%8A%9E",
					PromotedContent: "",
					Query:           "%E3%81%BB%E3%81%AE%E3%81%8A%E4%B8%80%E6%8A%9E",
					TweetVolume:     0,
				},
				{
					Name:            "黒スーツ",
					URL:             "http://twitter.com/search?q=%E9%BB%92%E3%82%B9%E3%83%BC%E3%83%84",
					PromotedContent: "",
					Query:           "%E9%BB%92%E3%82%B9%E3%83%BC%E3%83%84",
					TweetVolume:     10397,
				},
				{
					Name:            "最大4体ゲット",
					URL:             "http://twitter.com/search?q=%E6%9C%80%E5%A4%A74%E4%BD%93%E3%82%B2%E3%83%83%E3%83%88",
					PromotedContent: "",
					Query:           "%E6%9C%80%E5%A4%A74%E4%BD%93%E3%82%B2%E3%83%83%E3%83%88",
					TweetVolume:     24185,
				},
				{
					Name:            "iniの投稿動画",
					URL:             "http://twitter.com/search?q=ini%E3%81%AE%E6%8A%95%E7%A8%BF%E5%8B%95%E7%94%BB",
					PromotedContent: "",
					Query:           "ini%E3%81%AE%E6%8A%95%E7%A8%BF%E5%8B%95%E7%94%BB",
					TweetVolume:     0,
				},
				{
					Name:            "人達同士",
					URL:             "http://twitter.com/search?q=%E4%BA%BA%E9%81%94%E5%90%8C%E5%A3%AB",
					PromotedContent: "",
					Query:           "%E4%BA%BA%E9%81%94%E5%90%8C%E5%A3%AB",
					TweetVolume:     22919,
				},
				{
					Name:            "ウツホちゃん",
					URL:             "http://twitter.com/search?q=%E3%82%A6%E3%83%84%E3%83%9B%E3%81%A1%E3%82%83%E3%82%93",
					PromotedContent: "",
					Query:           "%E3%82%A6%E3%83%84%E3%83%9B%E3%81%A1%E3%82%83%E3%82%93",
					TweetVolume:     0,
				},
				{
					Name:            "ストレート勝ち",
					URL:             "http://twitter.com/search?q=%E3%82%B9%E3%83%88%E3%83%AC%E3%83%BC%E3%83%88%E5%8B%9D%E3%81%A1",
					PromotedContent: "",
					Query:           "%E3%82%B9%E3%83%88%E3%83%AC%E3%83%BC%E3%83%88%E5%8B%9D%E3%81%A1",
					TweetVolume:     0,
				},
				{
					Name:            "ファーストテイク",
					URL:             "http://twitter.com/search?q=%E3%83%95%E3%82%A1%E3%83%BC%E3%82%B9%E3%83%88%E3%83%86%E3%82%A4%E3%82%AF",
					PromotedContent: "",
					Query:           "%E3%83%95%E3%82%A1%E3%83%BC%E3%82%B9%E3%83%88%E3%83%86%E3%82%A4%E3%82%AF",
					TweetVolume:     0,
				},
				{
					Name:            "sakoさん",
					URL:             "http://twitter.com/search?q=sako%E3%81%95%E3%82%93",
					PromotedContent: "",
					Query:           "sako%E3%81%95%E3%82%93",
					TweetVolume:     0,
				},
				{
					Name:            "アイナちゃん",
					URL:             "http://twitter.com/search?q=%E3%82%A2%E3%82%A4%E3%83%8A%E3%81%A1%E3%82%83%E3%82%93",
					PromotedContent: "",
					Query:           "%E3%82%A2%E3%82%A4%E3%83%8A%E3%81%A1%E3%82%83%E3%82%93",
					TweetVolume:     0,
				},
				{
					Name:            "カワノルシア",
					URL:             "http://twitter.com/search?q=%E3%82%AB%E3%83%AF%E3%83%8E%E3%83%AB%E3%82%B7%E3%82%A2",
					PromotedContent: "",
					Query:           "%E3%82%AB%E3%83%AF%E3%83%8E%E3%83%AB%E3%82%B7%E3%82%A2",
					TweetVolume:     0,
				},
				{
					Name:            "会見中止要求",
					URL:             "http://twitter.com/search?q=%E4%BC%9A%E8%A6%8B%E4%B8%AD%E6%AD%A2%E8%A6%81%E6%B1%82",
					PromotedContent: "",
					Query:           "%E4%BC%9A%E8%A6%8B%E4%B8%AD%E6%AD%A2%E8%A6%81%E6%B1%82",
					TweetVolume:     19707,
				},
				{
					Name:            "旧統一教会側",
					URL:             "http://twitter.com/search?q=%E6%97%A7%E7%B5%B1%E4%B8%80%E6%95%99%E4%BC%9A%E5%81%B4",
					PromotedContent: "",
					Query:           "%E6%97%A7%E7%B5%B1%E4%B8%80%E6%95%99%E4%BC%9A%E5%81%B4",
					TweetVolume:     0,
				},
				{
					Name:            "木原選手",
					URL:             "http://twitter.com/search?q=%E6%9C%A8%E5%8E%9F%E9%81%B8%E6%89%8B",
					PromotedContent: "",
					Query:           "%E6%9C%A8%E5%8E%9F%E9%81%B8%E6%89%8B",
					TweetVolume:     0,
				},
				{
					Name:            "コバゴー",
					URL:             "http://twitter.com/search?q=%E3%82%B3%E3%83%90%E3%82%B4%E3%83%BC",
					PromotedContent: "",
					Query:           "%E3%82%B3%E3%83%90%E3%82%B4%E3%83%BC",
					TweetVolume:     0,
				},
				{
					Name:            "本人涙の続行",
					URL:             "http://twitter.com/search?q=%E6%9C%AC%E4%BA%BA%E6%B6%99%E3%81%AE%E7%B6%9A%E8%A1%8C",
					PromotedContent: "",
					Query:           "%E6%9C%AC%E4%BA%BA%E6%B6%99%E3%81%AE%E7%B6%9A%E8%A1%8C",
					TweetVolume:     0,
				},
				{
					Name:            "プエルトリコ",
					URL:             "http://twitter.com/search?q=%E3%83%97%E3%82%A8%E3%83%AB%E3%83%88%E3%83%AA%E3%82%B3",
					PromotedContent: "",
					Query:           "%E3%83%97%E3%82%A8%E3%83%AB%E3%83%88%E3%83%AA%E3%82%B3",
					TweetVolume:     0,
				},
				{
					Name:            "スプラフェス",
					URL:             "http://twitter.com/search?q=%E3%82%B9%E3%83%97%E3%83%A9%E3%83%95%E3%82%A7%E3%82%B9",
					PromotedContent: "",
					Query:           "%E3%82%B9%E3%83%97%E3%83%A9%E3%83%95%E3%82%A7%E3%82%B9",
					TweetVolume:     0,
				},
				{
					Name:            "清塚さん",
					URL:             "http://twitter.com/search?q=%E6%B8%85%E5%A1%9A%E3%81%95%E3%82%93",
					PromotedContent: "",
					Query:           "%E6%B8%85%E5%A1%9A%E3%81%95%E3%82%93",
					TweetVolume:     0,
				},
			},
			AsOf:      "2022-10-07T14:15:28Z",
			CreatedAt: "2022-10-05T21:42:57Z",
			Locations: []twitter.TrendsLocation{
				{
					Name:  "Fukuoka",
					WOEID: 1117099,
				},
			},
		},
	}, nil
}