package mime

import (
	"errors"
	"strings"
)

var (
	ErrBlankMimeType      = errors.New("mine type should not be blank")
	ErrBlankExt           = errors.New("extension should not be blank")
	ErrMimeTypesNotFound  = errors.New("mine types not found")
	ErrExtensionsNotFound = errors.New("extensions not found")
)

// GetExtensions Gets the extensions for the given MIME type.
// Return a slice of extensions (first one is the preferred one)
func GetExtensions(mimeType string) ([]string, error) {
	mimeType = strings.TrimSpace(mimeType)

	if mimeType == "" {
		return nil, ErrBlankMimeType
	}

	if e, ok := mimeTypesExtensions[mimeType]; ok {
		return e, nil
	}

	return nil, ErrExtensionsNotFound
}

// GetMimeTypes Gets the MIME types for the given extension.
// Return a slice of MIME types (first one is the preferred one)
func GetMimeTypes(ext string) ([]string, error) {
	ext = strings.TrimSpace(ext)
	ext = strings.TrimPrefix(ext, ".")

	if ext == "" {
		return nil, ErrBlankExt
	}

	if mt, ok := extensionMimeTypes[ext]; ok {
		return mt, nil
	}

	return nil, ErrMimeTypesNotFound
}
