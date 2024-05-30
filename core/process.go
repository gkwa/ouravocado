package core

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/taylormonacelli/somespider"
)

var cacheRelativePath string

func init() {
	cacheRelativePath = filepath.Join("ouravocado", "index.json")
}

func ProcessDirectories(dirs []string, verbose bool, ignorePaths []string) error {
	index, err := somespider.GenPath(cacheRelativePath)
	if err != nil {
		slog.Error("generating cache path failed", "error", err)
		return err
	}

	checksumCache := make(map[string]FileInfo)
	loadCacheFromFile(checksumCache)

	var allResults []string
	for _, dir := range dirs {
		results, err := scanDirectory(dir)
		if err != nil {
			return err
		}
		allResults = append(allResults, results...)
	}

	filteredResults := filterResults(allResults, ignorePaths)

	includeExtensions := []string{
		".md",
		".txt",
		".org",
	}
	filteredResults = filterByExtensions(filteredResults, includeExtensions)

	var fileInfos []FileInfo
	for _, path := range filteredResults {
		fileInfo, err := generateFileInfo(path, checksumCache)
		if err != nil {
			slog.Error("error generating file info", "path", path, "error", err)
			continue
		}
		fileInfos = append(fileInfos, fileInfo)
	}

	jsonData, err := json.MarshalIndent(fileInfos, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(index, jsonData, 0o644)
	if err != nil {
		return err
	}

	fmt.Printf("Updated %s\n", index)

	return nil
}

func filterByExtensions(paths []string, includeExtensions []string) []string {
	var filteredPaths []string
	for _, path := range paths {
		ext := strings.ToLower(filepath.Ext(path))
		for _, includeExt := range includeExtensions {
			if ext == strings.ToLower(includeExt) {
				filteredPaths = append(filteredPaths, path)
				break
			}
		}
	}
	return filteredPaths
}

func scanDirectory(dir string) ([]string, error) {
	var results []string

	slog.Debug("processing directory", "dir", dir)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			slog.Error("error walking directory", "path", path, "error", err)
			return err
		}

		if !info.IsDir() {
			results = append(results, path)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return results, nil
}

func filterResults(results []string, ignorePaths []string) []string {
	var filteredResults []string
	for _, result := range results {
		if !shouldIgnorePath(result, ignorePaths) {
			filteredResults = append(filteredResults, result)
		}
	}
	return filteredResults
}

func shouldIgnorePath(path string, ignorePaths []string) bool {
	for _, ignorePath := range ignorePaths {
		if strings.Contains(path, ignorePath) {
			return true
		}
	}
	return false
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