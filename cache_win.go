// +build windows

package xdg

import (
	"os"
	"path/filepath"
)

// CacheHome returns a XDG cache directory (XDG_CACHE_HOME).
func CacheHome() string {
	return alternate(os.Getenv(CacheHomeEnv), filepath.Join(os.Getenv("LOCALAPPDATA"), "cache"))
}
