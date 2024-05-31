package core

import (
	"encoding/json"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/gkwa/ouravocado/file"
	"github.com/taylormonacelli/somespider"
)

var (
	cacheRelativePath        = filepath.Join("ouravocado", "index.json")
	DefaultIncludeExtensions = []string{".md", ".txt", ".org"}
	index, _                 = somespider.GenPath(cacheRelativePath)
)

func ProcessDirectories(dirs []string, verbose bool, ignorePaths []string, includeExtensions []string) error {
	var filePaths []string
	var err error

	checksumCache := make(map[string]FileInfo)
	loadCacheFromFile(checksumCache, index)

	filePaths, err = file.ScanDirectories(dirs)
	if err != nil {
		return err
	}

	filePaths = file.FilterOutIgnoredPaths(filePaths, ignorePaths)
	filePaths = file.FilterByExtensions(filePaths, includeExtensions)

	fileInfos, err := generateFileInfos(filePaths, checksumCache)
	if err != nil {
		return err
	}

	err = saveFileInfosToCache(index, fileInfos)
	if err != nil {
		return err
	}

	return nil
}

func generateFileInfos(paths []string, checksumCache map[string]FileInfo) ([]FileInfo, error) {
	var fileInfos []FileInfo
	for _, path := range paths {
		fileInfo, err := generateFileInfo(path, checksumCache)
		if err != nil {
			slog.Error("error generating file info", "path", path, "error", err)
			continue
		}
		fileInfos = append(fileInfos, fileInfo)
	}
	return fileInfos, nil
}

func generateFileInfo(path string, checksumCache map[string]FileInfo) (FileInfo, error) {
	cachedFileInfo, ok := checksumCache[path]
	if ok {
		slog.Debug("checksum match, skipping word count", "path", path)
		return cachedFileInfo, nil
	}

	info, err := os.Stat(path)
	if err != nil {
		return FileInfo{}, err
	}

	checksum, err := calculateChecksum(path)
	if err != nil {
		return FileInfo{}, err
	}

	slog.Debug("calculating word count", "path", path)
	wordCount, err := countWords(path)
	if err != nil {
		return FileInfo{}, err
	}

	slog.Debug("word count", "path", path, "count", wordCount)

	size := info.Size()
	sizeFriendly := formatSize(size)

	fileInfo := FileInfo{
		Path:         path,
		WordCount:    wordCount,
		Size:         size,
		SizeFriendly: sizeFriendly,
		FastChecksum: checksum,
	}

	checksumCache[path] = fileInfo

	return fileInfo, nil
}

func saveFileInfosToCache(index string, fileInfos []FileInfo) error {
	jsonData, err := json.MarshalIndent(fileInfos, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(index, jsonData, 0o644)
	if err != nil {
		return err
	}

	slog.Info("Updated cache", "path", index)
	return nil
}
