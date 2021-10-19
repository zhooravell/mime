package mime

import (
	"testing"

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
