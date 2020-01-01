package xdg

import "testing"

func TestCache(t *testing.T) {
	t.Log(CacheHome())
	t.Log(DataHome())
	t.Log(DataDirs())
	t.Log(ConfigHome())
	t.Log(ConfigDirs())
	t.Log(RuntimeDir())
}
