package main

import (
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	// consumerKey := "o37I7QrJe6Jf7DqRuT7hHeuCo"
	// consumerSecret := "1HglkuCH30gceeSPGiFeSS4j5m3WGkUEeE92P09eN2tnUAJxKU"
	// accessToken := "1508039357898317828-5FsffejCmYX7VPdYqJNwRqcNONWqD9"
	// accessSecret := "gJo0LxhlPzyQfxW2oX4rwh9yONHIaTzr4WXVhvU9m7ri9"

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		IncludeEmail: twitter.Bool(true),
	}
	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	fmt.Printf("User's Name:%+v\n", user.Name)

	// send tweet
	// Send a Tweet
	tweet, resp, err := client.Statuses.Update("First tweets test", nil)

	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(tweet)
	fmt.Println(resp)

}
