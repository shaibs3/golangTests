package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/shaibs3/golangTests/s3"
)

func main() {
	file, err := os.Create("test_file")
	if err != nil {
		fmt.Println("Failed to create file", err)
		return
	}

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	},
	)

	client := s3manager.NewDownloader(sess)
	downloader := s3.NewDownloader(client, &sync.Mutex{}, 3)
	numBytes, _ := downloader.Download(file, "my_test_key", "my_test_bucket")

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
	file.Close()
}
