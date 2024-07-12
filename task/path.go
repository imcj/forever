package task

import (
	"os"
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

func (p Path) IsAbs() bool {
	return filepath.IsAbs(p.String())
}

func (p Path) IsExists() bool {
	if _, err := os.Stat(p.String()); os.IsNotExist(err) {
		return false
	}
	return true
}

func (p Path) Join(elements ...Path) Path {
	elements = append([]Path{p}, elements...)
	var converted []string
	for _, path := range elements {
		converted = append(converted, path.String())
	}
	return Path(filepath.Join(converted...))
}

func EmptyPath() Path {
	return Path("")
}

func (p Path) CommandPath(dir Path) Path {
	if !p.IsAbs() {
		full := dir.Join(p)

		if full.IsExists() {
			return full
		}

		return p
	}

	return p
}
