package main

import (
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

var consumerKey string = os.Getenv("FX_DEVTOOLS_BOT_CONSUMER_KEY")
var consumerSecret string = os.Getenv("FX_DEVTOOLS_BOT_CONSUMER_SECRET")
var accessToken string = os.Getenv("FX_DEVTOOLS_BOT_ACCESS_TOKEN")
var accessTokenSecret string = os.Getenv("FX_DEVTOOLS_BOT_ACCESS_TOKEN_SECRET")
var twitterClient *anaconda.TwitterApi

func TwitterClient() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	twitterClient = anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	return twitterClient
}

func Tweet(status string) (tweet anaconda.Tweet, err error) {
	var client *anaconda.TwitterApi
	if twitterClient != nil {
		client = twitterClient
	} else {
		client = TwitterClient()
	}
	return client.PostTweet(status, url.Values{})
}
