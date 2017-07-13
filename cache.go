package xdg

import "os"

const (
	CacheHomeEnv = "XDG_CACHE_HOME"
	CacheDirsEnv = "XDG_CACHE_DIRS"
)

// CacheHome returns a user XDG data directory (XDG_CACHE_HOME).
func CacheHome() string {
	return altHome(os.Getenv(CacheHomeEnv), ".cache")
}

// FindCacheFile finds a file from the XDG data directory.
// If one cannot be found, an error `ErrNotFound` be returned.
func FindCacheFile(rel ...string) (string, error) {
	return findFile([]string{CacheHome()}, rel)
}
