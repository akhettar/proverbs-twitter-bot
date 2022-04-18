package main

import (
	"os"
	"reflect"
	"testing"
)

func TestTweetParser_parse(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		tp   TweetParser
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"Testing loading valid file",
			TweetParser{},
			args{"tweets.txt"},
			[]string{"Absence makes the heart grow fonder.", "Men are best loved furthest off.", "Absence diminishes little passions and increases great ones.", "Absence sharpens love, presence strengthens it."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tp := TweetParser{}
			f, _ := os.Open(tt.args.filename)
			if got := tp.parse(f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TweetParser.parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
