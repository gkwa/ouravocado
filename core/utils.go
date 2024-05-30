package core

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"strconv"
	"strings"
)

func calculateChecksum(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	checksum := hex.EncodeToString(hash.Sum(nil))

	return checksum, nil
}

func countWords(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var wordCount int
	data := make([]byte, 1024)
	for {
		count, err := file.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}

		words := strings.Fields(string(data[:count]))
		wordCount += len(words)
	}

	return wordCount, nil
}

func formatSize(size int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB"}

	var i int
	fsize := float64(size)
	for i = 0; fsize >= 1024 && i < len(units)-1; i++ {
		fsize /= 1024
	}

	return strconv.FormatFloat(fsize, 'f', 2, 64) + " " + units[i]
}
