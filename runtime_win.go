// +build windows

package xdg

import (
	"os"
)

// RuntimeDir returns XDG runtime directory.
func RuntimeDir() string {
	// XDG_RUNTIME_DIR
	return alternate(os.Getenv(RuntimeDirEnv), os.Getenv("LOCALAPPDATA"))
}
