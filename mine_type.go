package mime

import (
	"errors"
	"strings"
)

var (
	ErrBlankMimeType = errors.New("mine type should not be blank")
	ErrBlankExt      = errors.New("extension should not be blank")
)

// GetExtensions Gets the extensions for the given MIME type.
// Return a slice of extensions (first one is the preferred one)
func GetExtensions(mimeType string) ([]string, error) {
	mimeType = strings.TrimSpace(mimeType)

	if mimeType == "" {
		return nil, ErrBlankMimeType
	}

	return nil, nil
}

// GetMimeTypes Gets the MIME types for the given extension.
// Return a slice of MIME types (first one is the preferred one)
func GetMimeTypes(ext string) ([]string, error) {
	ext = strings.TrimSpace(ext)

	if ext == "" {
		return nil, ErrBlankExt
	}

	return nil, nil
}
