package utils

import (
	"os"
	"time"
)

// ReadLastNBytes reads the last n bytes from the given file path.
func ReadLastNBytes(path string, n int64) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}

	size := fi.Size()
	start := size - n
	if start < 0 {
		start = 0
	}

	buf := make([]byte, size-start)
	_, err = file.ReadAt(buf, start)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// GetFileMetadata returns the modification time and size of the file at the given path.
func GetFileMetadata(path string) (modTime time.Time, size int64, err error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return time.Time{}, 0, err
	}
	return fileInfo.ModTime(), fileInfo.Size(), nil
}
