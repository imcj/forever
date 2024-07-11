package task

import (
	"path/filepath"
	"runtime"
)

type Path string

func (p Path) String() string {
	if runtime.GOOS == "windows" {
		return filepath.FromSlash(string(p))
	} else {
		return string(p)
	}
}
