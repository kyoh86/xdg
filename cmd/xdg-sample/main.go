// +build sample

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kyoh86/xdg"
)

func main() {
	fmt.Println("ConfigHome() will get a user level configuration directory path: ", xdg.ConfigHome())

	os.Setenv(xdg.ConfigHomeEnv, "/.config")
	fmt.Println("...And `XDG_CONFIG_HOME` envar will be supported: ", xdg.ConfigHome())

	fmt.Println(strings.Repeat("-", 48))

	fmt.Println("ConfigDirs() will get system level configuration directories: ", xdg.ConfigDirs())

	os.Setenv(xdg.ConfigDirsEnv, "/.config:/.config2")
	fmt.Println("...And `XDG_CONFIG_DIRS` envar will be able to change it:", xdg.ConfigDirs())

	fmt.Println(strings.Repeat("-", 48))

	fmt.Println("DataHome() will get a user level data directory path: ", xdg.DataHome())

	os.Setenv(xdg.DataHomeEnv, "/.data")
	fmt.Println("...And `XDG_DATA_HOME` envar will be supported: ", xdg.DataHome())

	fmt.Println(strings.Repeat("-", 48))

	fmt.Println("DataDirs() will get system level data directories: ", xdg.DataDirs())

	os.Setenv(xdg.DataDirsEnv, "/.data:/.data2")
	fmt.Println("...And `XDG_DATA_DIRS` envar will be able to change it:", xdg.DataDirs())

	fmt.Println(strings.Repeat("-", 48))

	fmt.Println("CacheHome() will get a cache directory path: ", xdg.CacheHome())

	os.Setenv(xdg.CacheHomeEnv, "/.cache")
	fmt.Println("...And `XDG_CACHE_HOME` envar will be supported: ", xdg.CacheHome())

	fmt.Println(strings.Repeat("-", 48))

	fmt.Println("RuntimeDir() will get a runtime directory path: ", xdg.RuntimeDir())

	os.Setenv(xdg.RuntimeDirEnv, "/.runtime")
	fmt.Println("...And `XDG_RUNTIME_DIR` envar will be supported: ", xdg.RuntimeDir())

	fmt.Println(strings.Repeat("-", 48))

	fmt.Println("DesktopDir() will get a desktop directory path: ", xdg.DesktopDir())
	os.Setenv(xdg.DesktopDirEnv, "/.Desktop")
	fmt.Println("...And `XDG_DESKTOP_DIR` envar will be supported: ", xdg.DesktopDir())

	fmt.Println("DownloadDir() will get a download directory path: ", xdg.DownloadDir())
	os.Setenv(xdg.DownloadDirEnv, "/.Download ")
	fmt.Println("...And `XDG_Download _DIR` envar will be supported: ", xdg.DownloadDir())

	fmt.Println("DocumentsDir() will get a documents directory path: ", xdg.DocumentsDir())
	os.Setenv(xdg.DocumentsDirEnv, "/.Documents ")
	fmt.Println("...And `XDG_Documents _DIR` envar will be supported: ", xdg.DocumentsDir())

	fmt.Println("MusicDir() will get a music directory path: ", xdg.MusicDir())
	os.Setenv(xdg.MusicDirEnv, "/.Music ")
	fmt.Println("...And `XDG_Music _DIR` envar will be supported: ", xdg.MusicDir())

	fmt.Println("PicturesDir() will get a pictures directory path: ", xdg.PicturesDir())
	os.Setenv(xdg.PicturesDirEnv, "/.Pictures ")
	fmt.Println("...And `XDG_Pictures _DIR` envar will be supported: ", xdg.PicturesDir())

	fmt.Println("VideosDir() will get a videos directory path: ", xdg.VideosDir())
	os.Setenv(xdg.VideosDirEnv, "/.Videos ")
	fmt.Println("...And `XDG_Videos _DIR` envar will be supported: ", xdg.VideosDir())

	fmt.Println("TemplatesDir() will get a templates directory path: ", xdg.TemplatesDir())
	os.Setenv(xdg.TemplatesDirEnv, "/.Templates ")
	fmt.Println("...And `XDG_Templates _DIR` envar will be supported: ", xdg.TemplatesDir())

	fmt.Println("PublicShareDir() will get a public share directory path: ", xdg.PublicShareDir())
	os.Setenv(xdg.PublicShareDirEnv, "/.PublicShare ")
	fmt.Println("...And `XDG_PublicShare_DIR` envar will be supported: ", xdg.PublicShareDir())
}
