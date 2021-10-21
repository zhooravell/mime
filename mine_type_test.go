package mime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetMimeTypes_EmptyExtensions(t *testing.T) {
	tests := map[string]string{
		"blank":    "",
		"new line": "\n",
		"tab":      "\t",
		"spaces":   "   ",
	}

	for name, ext := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := GetMimeTypes(ext)

			assert.Nil(t, res)
			assert.ErrorIs(t, err, ErrBlankExt)
		})
	}
}

func TestGetExtensions_EmptyMimeTypes(t *testing.T) {
	tests := map[string]string{
		"blank":    "",
		"new line": "\n",
		"tab":      "\t",
		"spaces":   "   ",
	}

	for name, mt := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := GetExtensions(mt)

			assert.Nil(t, res)
			assert.ErrorIs(t, err, ErrBlankMimeType)
		})
	}
}

func TestGetMimeTypes_NotFound(t *testing.T) {
	tests := map[string]string{
		"uuid":     "dae5b3d9-3c52-4716-9cc6-b6a05fa256a4",
		"Jon Snow": "Jon Snow",
		"time":     time.Now().Format("02-Mar-2006 15:04:05"),
	}

	for name, mt := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := GetMimeTypes(mt)

			assert.Nil(t, res)
			assert.ErrorIs(t, err, ErrMimeTypesNotFound)
		})
	}
}

func TestGetExtensions_NotFound(t *testing.T) {
	tests := map[string]string{
		"uuid":       "00000000-0000-0000-0000-000000000000",
		"Arya Stark": "Arya Stark",
		"time":       time.Now().Format("02-Mar-2006 15:04:05"),
	}

	for name, mt := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := GetExtensions(mt)

			assert.Nil(t, res)
			assert.ErrorIs(t, err, ErrExtensionsNotFound)
		})
	}
}

func TestGetExtensions_OK(t *testing.T) {
	tests := map[string][]string{
		"application/pdf": {"pdf"},
		"text/yaml":       {"yaml", "yml"},
		"text/css":        {"css"},
		"image/png":       {"png"},
	}

	for name, exts := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := GetExtensions(name)

			assert.Nil(t, err)
			assert.Equal(t, exts, res)
		})
	}
}

func TestGetMimeTypes_OK(t *testing.T) {
	tests := map[string][]string{
		".twig":  {"text/x-twig"},
		"tar.gz": {"application/x-compressed-tar"},
	}

	for name, exts := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := GetMimeTypes(name)

			assert.Nil(t, err)
			assert.Equal(t, exts, res)
		})
	}
}
