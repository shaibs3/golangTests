package phase_3

import (
	"log"
	"os"

	"github.com/avast/retry-go"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/shaibs3/golangTests/phase_2"
)

type Locker interface {
	Lock()
	Unlock()
}

type FinalDownloader struct {
	client     phase_2.S3Downloader
	lock       Locker
	numRetries int
}

func NewFinalDownloader(client phase_2.S3Downloader, lock Locker, retries int) *FinalDownloader {
	return &FinalDownloader{
		client:     client,
		lock:       lock,
		numRetries: retries,
	}
}

func (fd *FinalDownloader) Download(file *os.File, key, bucket string) (int64, error) {
	var numBytes int64
	s3Obj := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	fd.lock.Lock()

	err := retry.Do(
		func() error {
			var err error
			numBytes, err = fd.client.Download(file, s3Obj)
			if err != nil {
				return err
			}

			return err
		},
		retry.Attempts(uint(fd.numRetries)),
		retry.OnRetry(func(n uint, err error) {
			log.Printf("Retrying request after error: %v", err)
		}),
	)
	if err != nil {
		return -1, err
	}

	fd.lock.Unlock()
	return numBytes, err
}
