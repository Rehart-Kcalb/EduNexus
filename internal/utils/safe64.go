package utils

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"time"

	"github.com/gabriel-vasile/mimetype"
)

var base64Pattern = regexp.MustCompile(`^data:(.*?);base64,(.*)$`)

func SaveBase64ToFile(base64String, storageDir string) (string, error) {
	// Strip out metadata prefix if present
	matches := base64Pattern.FindStringSubmatch(base64String)
	if len(matches) == 3 {
		base64String = matches[2]
	}

	// Decode the base64 string
	data, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return "", fmt.Errorf("invalid base64 data: %w", err)
	}

	// Detect MIME type
	mimeType := mimetype.Detect(data)
	ext := mimeType.Extension()
	if ext == "" {
		return "", fmt.Errorf("unsupported file type")
	}

	// Create a file in the storage directory with the correct extension
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(storageDir, filename)
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	return filePath, nil
}
