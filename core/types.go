package core

type FileInfo struct {
	Path         string `json:"path"`
	WordCount    int    `json:"wordCount"`
	Size         int64  `json:"size"`
	SizeFriendly string `json:"sizeFriendly"`
	FastChecksum string `json:"fastChecksum"`
}
