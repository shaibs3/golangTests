package finalDownloader_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/shaibs3/golangTests/finalDownloader"
	"github.com/shaibs3/golangTests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBasic(t *testing.T) {
	s := mocks.NewS3Downloader(t)
	s.EXPECT().Download(mock.Anything, mock.Anything).Return(0, nil)
	lock := mocks.NewLocker(t)
	lock.EXPECT().Lock().Return()
	lock.EXPECT().Unlock().Return()
	down := finalDownloader.NewFinalDownloader(s, lock, 3)

	file, err := os.Create("key")
	if err != nil {
		fmt.Println("Failed to create file", err)
		return
	}

	defer file.Close()
	n, err := down.Download(file, "test", "test")
	assert.NoError(t, err)
	assert.Equal(t, int64(0), n)
}

func TestError(t *testing.T) {
	s := mocks.NewS3Downloader(t)
	lock := mocks.NewLocker(t)
	lock.EXPECT().Lock().Return()
	lock.EXPECT().Unlock().Return()
	down := finalDownloader.NewFinalDownloader(s, lock, 3)
	s.EXPECT().Download(mock.Anything, mock.Anything).Return(0, errors.New("error"))
	file, err := os.Create("key")
	if err != nil {
		fmt.Println("Failed to create file", err)
		return
	}

	defer file.Close()
	n, err := down.Download(file, "test", "test")
	assert.Error(t, err)
	assert.Equal(t, int64(-1), n)
}

func TestRetrySuccess(t *testing.T) {
	s := mocks.NewS3Downloader(t)
	lock := mocks.NewLocker(t)
	lock.EXPECT().Lock().Return()
	lock.EXPECT().Unlock().Return()
	down := finalDownloader.NewFinalDownloader(s, lock, 5)
	s.EXPECT().Download(mock.Anything, mock.Anything).Times(4).Return(int64(0), errors.New("error"))
	s.EXPECT().Download(mock.Anything, mock.Anything).Return(int64(0), nil)
	file, _ := os.Create("key")
	defer file.Close()
	n, err := down.Download(file, "test", "test")
	assert.NoError(t, err)
	assert.NotEqual(t, int64(-1), n)
	s.AssertNumberOfCalls(t, "Download", 5)
}
