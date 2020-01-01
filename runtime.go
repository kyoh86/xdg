package xdg

const (
	// RuntimeDirEnv is the name of the environment variable holding a runtime directory path.
	RuntimeDirEnv = "XDG_RUNTIME_DIR"
)

// FindRuntimeFile finds a file from the XDG runtime directory.
// If one cannot be found, an error `ErrNotFound` be returned.
func FindRuntimeFile(rel ...string) (string, error) {
	return findFile([]string{RuntimeDir()}, rel)
}
