package task_test

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestFilePath_Slash(t *testing.T) {
	v := filepath.FromSlash("C:/Users/")
	assert.Equal(t, "C:\\Users\\", v)

	v = filepath.FromSlash("C:\\Users\\")
	assert.Equal(t, "C:\\Users\\", v)
}
