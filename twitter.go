package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// TwitterService exposes API to publish tweets to twitter account
type TwitterService interface {
	tweet(tweet) (*twitter.Tweet, *http.Response, error)
}

type tweeter struct {
	client *twitter.Client
}

func (t *tweeter) tweet(msg tweet) (*twitter.Tweet, *http.Response, error) {

	tweet, resp, err :=t.client.Statuses.Update(string(msg), nil)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n\n\n", resp)
	log.Printf("%+v\n", tweet)
	return tweet, resp, err
}

// NewTweeter creats a new TwitterService instance
func NewTweeter() (TwitterService, error) {

	// Pass in your consumer key (API Key) and your Consumer Secret (API Secret)
	config := oauth1.NewConfig(getEnvVar("TWITTER_CONSUMER_KEY"), getEnvVar("TWITTER_CONSUMER_SECRET"))
	// Pass in your Access Token and your Access Token Secret
	token := oauth1.NewToken(getEnvVar("TWITTER_ACCESS_TOKEN"),getEnvVar("TWITTER_ACCESS_SECRET"))

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// we can retrieve the user and verify if the credentials
	// we have used successfully allow us to log in!
	_, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	// log.Printf("User's ACCOUNT:\n%+v\n", user)
	return &tweeter{client: client}, nil
}

func getEnvVar(key string) string {
	val, exsit := os.LookupEnv(key)
	if !exsit {
		log.Fatalf("Environment variable with key: %s is mandatory", key)
	}
	return val
}
	
	



