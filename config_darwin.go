// +build darwin

package xdg

import (
	"os"
	"path/filepath"
)

// ConfigHome returns a user XDG configuration directory (XDG_CONFIG_HOME).
func ConfigHome() string {
	return altHome(os.Getenv(ConfigHomeEnv), "Library", "Preferences")
}

// ConfigDirs returns system XDG data directories (XDG_CONFIG_DIRS).
func ConfigDirs() []string {
	// XDG_CONFIG_DIRS
	xdgDirs := filepath.SplitList(os.Getenv(ConfigDirsEnv))
	if len(xdgDirs) != 0 {
		return xdgDirs
	}
	return []string{
		filepath.Join("/", "Library", "Preferences"),
	}
}
