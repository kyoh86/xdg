package xdg

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	DataHomeEnv = "XDG_DATA_HOME"
	DataDirsEnv = "XDG_DATA_DIRS"
)

// DataHome returns a user XDG data directory (XDG_DATA_HOME).
func DataHome() string {
	return altHome(os.Getenv(DataHomeEnv), ".local", "share")
}

// DataDirs returns system XDG data directories (XDG_DATA_DIRS).
func DataDirs() []string {
	// XDG_DATA_DIRS
	xdgDirs := alternate(
		os.Getenv(DataDirsEnv),
		strings.Join([]string{
			filepath.Join("/", "usr", "local", "share"),
			filepath.Join("/", "usr", "share"),
		}, string(filepath.ListSeparator)),
	)
	return filepath.SplitList(xdgDirs)
}

// AllDataDirs returns all XDG data directories.
func AllDataDirs() []string {
	var dirs []string

	// XDG_DATA_HOME
	if home := DataHome(); home != "" {
		dirs = append(dirs, home)
	}

	// XDG_DATA_DIRS
	dirs = append(dirs, DataDirs()...)

	return dirs
}

// FindDataFile finds a file from the XDG data directory.
// If one cannot be found, an error `ErrNotFound` be returned.
func FindDataFile(rel ...string) (string, error) {
	return findFile(AllDataDirs(), rel)
}
