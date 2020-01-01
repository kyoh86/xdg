// +build darwin

package xdg

import (
	"os"
)

// RuntimeDir returns XDG runtime directory.
func RuntimeDir() string {
	// XDG_RUNTIME_DIR
	return altHome(os.Getenv(RuntimeDirEnv), "Library", "Application Support")
}
