package s3

import (
	"log"
	"os"
	"sync"

	"github.com/avast/retry-go"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Downloader struct {
	client     *s3manager.Downloader
	lock       *sync.Mutex
	numRetries uint
	key        string
	bucket     string
}

func NewDownloader(client *s3manager.Downloader, lock *sync.Mutex,
	key, bucket string, numRetries uint,
) *Downloader {
	return &Downloader{
		client:     client,
		lock:       lock,
		numRetries: numRetries,
		key:        key,
		bucket:     bucket,
	}
}

func (s3Client *Downloader) Download(file *os.File) (int64, error) {
	var numBytes int64
	s3Obj := &s3.GetObjectInput{
		Bucket: aws.String(s3Client.bucket),
		Key:    aws.String(s3Client.key),
	}
	s3Client.lock.Lock()

	err := retry.Do(
		func() error {
			var err error
			numBytes, err = s3Client.client.Download(file, s3Obj)
			if err != nil {
				return err
			}

			return err
		},
		retry.Attempts(s3Client.numRetries),
		retry.OnRetry(func(n uint, err error) {
			log.Printf("Retrying request after error: %v", err)
		}),
	)
	if err != nil {
		return -1, err
	}

	s3Client.lock.Unlock()
	return numBytes, err
}
