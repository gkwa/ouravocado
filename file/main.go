package file

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

func ScanDirectories(dirs []string) ([]string, error) {
	var allResults []string
	for _, dir := range dirs {
		results, err := ScanDirectory(dir)
		if err != nil {
			return nil, err
		}
		allResults = append(allResults, results...)
	}
	return allResults, nil
}

func ScanDirectory(dir string) ([]string, error) {
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

func FilterOutIgnoredPaths(paths []string, ignorePaths []string) []string {
	var filteredPaths []string
	for _, path := range paths {
		if !ShouldIgnorePath(path, ignorePaths) {
			filteredPaths = append(filteredPaths, path)
		}
	}
	return filteredPaths
}

func FilterByExtensions(paths []string, includeExtensions ...string) []string {
	extMap := make(map[string]bool)
	for _, ext := range includeExtensions {
		extMap["."+strings.TrimLeft(strings.ToLower(ext), ".")] = true
	}

	var filteredPaths []string
	for _, path := range paths {
		ext := strings.ToLower(filepath.Ext(path))
		if extMap[ext] {
			filteredPaths = append(filteredPaths, path)
		}
	}
	return filteredPaths
}

func ShouldIgnorePath(p string, ignores []string) bool {
	p = strings.ToLower(p)
	for _, ignore := range ignores {
		if strings.Contains(p, strings.ToLower(ignore)) {
			return true
		}
	}
	return false
}
