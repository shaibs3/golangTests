package phase_2

import (
	"io"
	"log"
	"os"
	"sync"

)

type S3Downloader interface {
	Download(w io.WriterAt, input *s3.GetObjectInput, options ...func(*s3manager.Downloader)) (n int64, err error)
}
type Downloader struct {
	client     S3Downloader
	lock       *sync.Mutex
	numRetries int
}

func NewDownloader(client S3Downloader, lock *sync.Mutex, retries int) *Downloader {
	return &Downloader{
		client:     client,
		lock:       lock,
		numRetries: retries,
	}
}

func (s3Client *Downloader) Download(file *os.File, key, bucket string) (int64, error) {
	var numBytes int64
	s3Obj := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
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
		retry.Attempts(uint(s3Client.numRetries)),
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
