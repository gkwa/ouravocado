package core

import (
	"encoding/json"
	"log/slog"
	"os"
)

func loadCacheFromFile(checksumCache map[string]FileInfo) {
	data, err := os.ReadFile("index.json")
	if err != nil {
		if !os.IsNotExist(err) {
			slog.Error("error reading index.json", "error", err)
		}
		return
	}

	var fileInfos []FileInfo
	err = json.Unmarshal(data, &fileInfos)
	if err != nil {
		slog.Error("error unmarshaling index.json", "error", err)
		return
	}

	for _, fileInfo := range fileInfos {
		checksumCache[fileInfo.Path] = fileInfo
	}
}
