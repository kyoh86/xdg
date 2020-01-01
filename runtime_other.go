// +build !windows,!darwin

package xdg

import (
	"os"
	"path/filepath"
	"strconv"
)

// RuntimeDir returns XDG runtime directory.
func RuntimeDir() string {
	// XDG_RUNTIME_DIR
	return alternate(
		os.Getenv(RuntimeDirEnv),
		filepath.Join("/", "run", "user", strconv.Itoa(os.Getuid())),
	)
}
