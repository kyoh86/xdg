// +build windows

package xdg

func altUserProfile(path string, suffix ...string) string {
	if path != "" {
		return path
	}
	home := os.Getenv("USERPROFILE")
	if home != "" {
		return filepath.Join(append([]string{home}, suffix...)...)
	}
	return ""
}
