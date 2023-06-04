package phase_2_test

import (
	"errors"
	"os"
	"sync"
	"testing"

	"github.com/shaibs3/golangTests/mocks"
	"github.com/shaibs3/golangTests/phase_2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var file *os.File

func TestBasic(t *testing.T) {
	// create and setup the mock
	s := mocks.NewS3Downloader(t)
	s.EXPECT().Download(mock.Anything, mock.Anything).Return(64, nil)

	// call the api
	down := phase_2.NewDownloader(s, &sync.Mutex{}, 3)
	n, err := down.Download(file, "test", "test")

	// verify assertions
	assert.NoError(t, err)
	assert.Equal(t, int64(64), n)
}

func TestError(t *testing.T) {
	// create and setup the mock
	s := mocks.NewS3Downloader(t)
	down := phase_2.NewDownloader(s, &sync.Mutex{}, 3)
	s.EXPECT().Download(mock.Anything, mock.Anything).Return(0, errors.New("error"))

	// call the api
	n, err := down.Download(file, "test", "test")

	// verify assertions
	assert.Error(t, err)
	assert.Equal(t, int64(-1), n)
}

func TestRetrySuccess(t *testing.T) {
	// create and setup the mock
	s := mocks.NewS3Downloader(t)
	down := phase_2.NewDownloader(s, &sync.Mutex{}, 5)
	s.EXPECT().Download(mock.Anything, mock.Anything).Times(4).Return(int64(0), errors.New("error"))
	s.EXPECT().Download(mock.Anything, mock.Anything).Return(int64(0), nil)

	// call the api
	n, err := down.Download(file, "test", "test")

	// verify assertions
	assert.NoError(t, err)
	assert.NotEqual(t, int64(-1), n)
	s.AssertNumberOfCalls(t, "Download", 5)
}

func TestMain(m *testing.M) {
	file, _ := os.Create("key")
	defer file.Close()
	exitVal := m.Run()
	os.Exit(exitVal)
}
