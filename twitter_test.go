package main

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/dghubble/go-twitter/twitter"
)

func Test_tweeter_tweet(t *testing.T) {
	type fields struct {
		client *twitter.Client
	}
	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *twitter.Tweet
		want1   *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &tweeter{
				client: tt.fields.client,
			}
			got, got1, err := tr.tweet(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("tweeter.tweet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tweeter.tweet() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("tweeter.tweet() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
