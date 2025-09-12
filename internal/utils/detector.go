package utils

import (
	"strings"
)

func DetectMediaType(extension string) string {
	switch strings.ToLower(extension) {
	case "mp4", "mkv", "webm", "avi", "mov":
		return strings.Join([]string{"video", extension}, "/")
	case "mp3", "aac", "wav", "flac", "m4a":
		return strings.Join([]string{"audio", extension}, "/")
	default:
		return strings.Join([]string{"other", extension}, "/")
	}
}
