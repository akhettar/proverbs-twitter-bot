package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// S3Client ...
type S3Client struct {
	*s3manager.Downloader
}

func (sc3 *S3Client) download(bucket,filename string) (*os.File, error) {

	// 1. Create temporary file where to downlaod the tweets to
	file, err := os.Create("tempFile")

	if err != nil {
		log.Fatal(err)
	}

	// create S3 Get Object request
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	}

	// Write the contents of S3 Object to the file
	n, err := sc3.Download(file, input)

	if err != nil {
		return 0, fmt.Errorf("failed to download file, %v", err)
	}
	fmt.Printf("file downloaded, %d bytes\n", n)
	return file, nil
}

// NewS3Client ...
func NewS3Client() *S3Client {
	return &S3Client{s3manager.NewDownloader(session.Must(session.NewSession()))}
}
