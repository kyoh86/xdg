package xdg

import "testing"

func TestPath(t *testing.T) {
	t.Log(CacheHome())
	t.Log(DataHome())
	t.Log(DataDirs())
	t.Log(ConfigHome())
	t.Log(ConfigDirs())
	t.Log(RuntimeDir())
	t.Log(DesktopDir())
	t.Log(DownloadDir())
	t.Log(DocumentsDir())
	t.Log(MusicDir())
	t.Log(PicturesDir())
	t.Log(VideosDir())
	t.Log(TemplatesDir())
	t.Log(PublicShareDir())
}
