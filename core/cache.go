package core

import (
	"encoding/json"
	"log/slog"
	"os"
)

func loadCacheFromFile(checksumCache map[string]FileInfo, index string) {
	data, err := os.ReadFile(index)
	if err != nil {
		if !os.IsNotExist(err) {
			slog.Error("error reading cache", "path", index, "error", err)
		}
		return
	}

	var fileInfos []FileInfo
	err = json.Unmarshal(data, &fileInfos)
	if err != nil {
		slog.Error("error unmarshaling cache", "path", index, "error", err)
		return
	}

	for _, fileInfo := range fileInfos {
		checksumCache[fileInfo.Path] = fileInfo
	}
}
