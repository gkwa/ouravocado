// Package file provides utility functions for scanning and filtering file paths.
//
// To scan directories and retrieve file paths:
//
//	filePaths, err := file.ScanDirectories(dirs)
//	if err != nil {
//		// Handle error
//	}
//
// To filter file paths based on ignored paths:
//
//	filteredPaths := file.FilterOutIgnoredPaths(filePaths, ignorePaths)
//
// To filter file paths based on file extensions:
//
//	filteredPaths := file.FilterByExtensions(filePaths, includeExtensions)
//
// Example usage:
//
//	dirs := []string{"/path/to/dir1", "/path/to/dir2"}
//	ignorePaths := []string{".git", "node_modules"}
//	includeExtensions := []string{".go", ".md"}
//
//	filePaths, err := file.ScanDirectories(dirs)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	filePaths = file.FilterOutIgnoredPaths(filePaths, ignorePaths)
//	filePaths = file.FilterByExtensions(filePaths, includeExtensions)
//
//	for _, path := range filePaths {
//		fmt.Println(path)
//	}
package file
