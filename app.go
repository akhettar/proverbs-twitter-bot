package main

import (
	"log"
	
)

func main() {

	// 1. Downlod the file
	file, err := NewS3Client().download(getEnvVar("BUCKET_NAME"), getEnvVar("FILE_NAME"))
	if err != nil {
		log.Fatalf("failed to downalod the file: %s from the bucket: %s", getEnvVar("FILE_NAME") ,getEnvVar("BUCKET_NAME"))
	}
	
	// 2. Create twitter client
	twitter, err := NewTweeter()
	if err != nil {
		log.Fatal(err)
	}

	// 4. Push the tweet
	tweet := newTweetParser().parse(file)
	if _, _, err := twitter.tweet(tweet); err != nil {
		log.Printf("failed to push the tweets")
	}

	log.Printf("tweet: %s successfully pushed", tweet)

}
