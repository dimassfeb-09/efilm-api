package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifyFile(t *testing.T) {

	fileTypes := []string{"image/jpeg", "image/png", "image/jpg"}
	for _, fileType := range fileTypes {
		result := VerfiyFileType(fileType)
		t.Run("expect true", func(t *testing.T) {
			assert.True(t, result, "is should be true")
		})
	}

	fileTypes = []string{"application/pdf", "application/json"}
	for _, fileType := range fileTypes {
		result := VerfiyFileType(fileType)
		t.Run("expect true", func(t *testing.T) {
			assert.False(t, result, "is should be false")
		})
	}
}
