package main

import (
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// TweetParser instance type
type TweetParser struct {}

type tweet string

func newTweetParser() TweetParser {
	return TweetParser{}
}

func (tp TweetParser) parse(file io.Reader) tweet {
	var tweets[]string
	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatalf("failed to read the content of the file: %v", err)
	}

	reg := regexp.MustCompile("(\\d+)( )(.*)")
	results := reg.FindAllStringSubmatch(string(bytes), -1)
	for _, tweet := range results {
		tweets = append(tweets, strings.TrimSpace(tweet[3]))
	}
	log.Println(tweets)

	// return a random tweet from the list
	rand.Seed(time.Now().UnixNano())
	return tweet(tweets[rand.Intn(len(tweets) -1)])
}
