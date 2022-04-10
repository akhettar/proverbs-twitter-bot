package main

import (
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func getClient(creds *Credentials) (*twitter.Client, error) {
	// Pass in your consumer key (API Key) and your Consumer Secret (API Secret)
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	// Pass in your Access Token and your Access Token Secret
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

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
	return client, nil
}

func getAppClient(creds *Credentials) (*twitter.Client, error) {

	// oauth2 configures a client that uses app credentials to keep a fresh token
	config := &clientcredentials.Config{
		ClientID:     creds.ConsumerKey,
		ClientSecret: creds.ConsumerSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext)

	// Twitter client
	client := twitter.NewClient(httpClient)

	return client, nil

}

func main() {
	consumerKey := "o37I7QrJe6Jf7DqRuT7hHeuCo"
	consumerSecret := "1HglkuCH30gceeSPGiFeSS4j5m3WGkUEeE92P09eN2tnUAJxKU"
	accessToken := "1508039357898317828-CI5Mj4vMp9i21qy5FsXQCxrU5fmBjP"
	accessSecret := "mgoyZI5MLoijqGNyBEgOpNkj9zekD9EU2SO1jJZs0rh6R"

	fmt.Println("Go-Twitter Bot v0.01")

	creds := Credentials{
		AccessToken:       accessToken,
		AccessTokenSecret: accessSecret,
		ConsumerKey:       consumerKey,
		ConsumerSecret:    consumerSecret,
	}

	fmt.Printf("%+v\n", creds)

	client, err := getClient(&creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}

	// Print out the pointer to our client
	// for now so it doesn't throw errors
	//fmt.Printf("%+v\n", client)

	tweet, resp, err := client.Statuses.Update("Habit is a second nature", nil)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n\n\n", resp)
	log.Printf("%+v\n", tweet)

}
