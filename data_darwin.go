// +build darwin

package xdg

import (
	"os"
	"path/filepath"
)

// DataHome returns a user XDG data directory (XDG_DATA_HOME).
func DataHome() string {
	return altHome(os.Getenv(DataHomeEnv), "Library", "Application Support")
}

// DataDirs returns system XDG data directories (XDG_DATA_DIRS).
func DataDirs() []string {
	// XDG_DATA_DIRS
	xdgDirs := filepath.SplitList(os.Getenv(DataDirsEnv))
	if len(xdgDirs) != 0 {
		return xdgDirs
	}
	return []string{
		filepath.Join("/", "Library", "Application Support"),
	}
}
