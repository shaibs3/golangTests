package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/shaibs3/golangTests/phase_1"
	"github.com/spf13/pflag"
)

func main() {
	key, bucket, file, client := setup()

	downloader := phase_1.NewDownloader(client, &sync.Mutex{}, 3)
	numBytes, _ := downloader.Download(file, key, bucket)

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
	file.Close()
}

func setup() (string, string, *os.File, *s3manager.Downloader) {
	// define the application's flags
	outFile := pflag.String("out_file", "test_file", "output file")
	key := pflag.String("key", "test_key", "the s3 key")
	bucket := pflag.String("bucket", "test_bucket", "the s3 bucket")
	pflag.Parse()

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	session, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	},
	)
	client := s3manager.NewDownloader(session)
	// create an output file to write the S3 Object contents to.
	file, err := os.Create(*outFile)
	if err != nil {
		fmt.Println("Failed to create file", err)
		return "", "", nil, nil
	}
	return *key, *bucket, file, client
}
