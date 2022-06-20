package main

import (
	"context"
	"log"

	 "github.com/aws/aws-lambda-go/lambda"
)

func hanler(context.Context) error {
	// run the lambda
	// 1. Downlod the file
	// os.Setenv("BUCKET_NAME", "english-proverbs-cirta")
	// os.Setenv("FILE_NAME", "proverbs.txt")
	file, err := NewS3Client().download(getEnvVar("BUCKET_NAME"), getEnvVar("FILE_NAME"))
	if err != nil {
		log.Println(err)
		log.Fatalf("failed to downalod the file: %s from the bucket: %s", getEnvVar("FILE_NAME") ,getEnvVar("BUCKET_NAME"))
	}
	
	// 2. Create twitter client
	twitter, err := NewTweeter()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// 4. Push the tweet
	tweet := newTweetParser().parse(file)
	if _, _, err := twitter.tweet(tweet); err != nil {
		log.Printf("failed to push the tweets")
		return err
	}

	log.Printf("tweet: %s successfully pushed", tweet)
	return nil

}

func main() {
	lambda.Start(hanler)
}
