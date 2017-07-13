package xdg

import (
	"os"
	"path/filepath"
)

const (
	ConfigHomeEnv = "XDG_CONFIG_HOME"
	ConfigDirsEnv = "XDG_CONFIG_DIRS"
)

// ConfigHome returns a user XDG configuration directory (XDG_CONFIG_HOME).
func ConfigHome() string {
	return altHome(os.Getenv(ConfigHomeEnv), ".config")
}

// ConfigDirs returns system XDG configuration directories (XDG_CONFIG_DIRS).
func ConfigDirs() []string {
	// XDG_CONFIG_DIRS
	xdgDirs := alternate(
		os.Getenv(ConfigDirsEnv),
		filepath.Join("/", "etc", "xdg"),
	)
	return filepath.SplitList(xdgDirs)
}

// AllConfigDirs returns all XDG configuration directories.
func AllConfigDirs() []string {
	var dirs []string

	// XDG_CONFIG_HOME
	if home := ConfigHome(); home != "" {
		dirs = append(dirs, home)
	}

	// XDG_CONFIG_DIRS
	dirs = append(dirs, ConfigDirs()...)

	return dirs
}

// FindConfigFile finds a file from the XDG configuration directory.
// If one cannot be found, an error `ErrNotFound` be returned.
func FindConfigFile(rel ...string) (string, error) {
	return findFile(AllConfigDirs(), rel)
}
