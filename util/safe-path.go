package util

import (
	"path/filepath"
	"strings"
)

func IsSafePath(baseDir, userPath string) bool {
	if !filepath.IsAbs(userPath) {
		return false
	}

	// yes, this is ai shit and i am not sure if this is safe!!

	// 1. Ensure paths are Clean (removes /./ and redundant //)
	base := filepath.Clean(baseDir)
	target := filepath.Clean(userPath)

	// 2. Calculate the relative path from base to target
	// This is the "Gold Standard" for subpath checking.
	rel, err := filepath.Rel(base, target)
	if err != nil {
		return false
	}

	// 3. If the relative path starts with "..", it escaped the base.
	// If the relative path IS "..", it's the parent directory.
	if strings.HasPrefix(rel, "..") {
		return false
	}

	// 4. On Linux, if the user provides an absolute path that
	// points somewhere else entirely, Rel will return a path
	// starting with "../" to get there.
	return true
}
