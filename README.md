# xdg

Light weight helper functions in golang to get config,
data, cache and some user directories according to
the [XDG Base Directory](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html) and XDG User Directory Specification.

[![Go Report Card](https://goreportcard.com/badge/github.com/kyoh86/xdg)](https://goreportcard.com/report/github.com/kyoh86/xdg)
[![Coverage Status](https://img.shields.io/codecov/c/github/kyoh86/xdg.svg)](https://codecov.io/gh/kyoh86/xdg)

## Install

```console
$ go get github.com/kyoh86/xdg
```

## Usage

For example:

```go
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
```

And you can run sample easily.

```console
$ make sample
go run -tags=sample ./cmd/xdg-sample/main.go
ConfigHome() will get a user level configuration directory path:  /home/kyoh86/.config/
...And `XDG_CONFIG_HOME` envar will be supported:  /.config
------------------------------------------------
ConfigDirs() will get system level configuration directories:  [/etc/xdg]
...And `XDG_CONFIG_DIRS` envar will be able to change it: [/.config /.config2]
------------------------------------------------
DataHome() will get a user level data directory path:  /home/kyoh86/.local/share
...And `XDG_DATA_HOME` envar will be supported:  /.data
------------------------------------------------
DataDirs() will get system level data directories:  [/usr/local/share /usr/share]
...And `XDG_DATA_DIRS` envar will be able to change it: [/.data /.data2]
------------------------------------------------
CacheHome() will get a cache directory path:  /home/kyoh86/.cache
...And `XDG_CACHE_HOME` envar will be supported:  /.cache
------------------------------------------------
RuntimeDir() will get a runtime directory path:  /run/user/1000
...And `XDG_RUNTIME_DIR` envar will be supported:  /.runtime
------------------------------------------------
DesktopDir() will get a desktop directory path:  /home/kyoh86/Desktop
...And `XDG_DESKTOP_DIR` envar will be supported:  /.Desktop
DownloadDir() will get a download directory path:  /home/kyoh86/Downloads
...And `XDG_Download _DIR` envar will be supported:  /.Download 
DocumentsDir() will get a documents directory path:  /home/kyoh86/Documents
...And `XDG_Documents _DIR` envar will be supported:  /.Documents 
MusicDir() will get a music directory path:  /home/kyoh86/Music
...And `XDG_Music _DIR` envar will be supported:  /.Music 
PicturesDir() will get a pictures directory path:  /home/kyoh86/Pictures
...And `XDG_Pictures _DIR` envar will be supported:  /.Pictures 
VideosDir() will get a videos directory path:  /home/kyoh86/Videos
...And `XDG_Videos _DIR` envar will be supported:  /.Videos 
TemplatesDir() will get a templates directory path:  /home/kyoh86/Templates
...And `XDG_Templates _DIR` envar will be supported:  /.Templates 
PublicShareDir() will get a public share directory path:  /home/kyoh86/Public
...And `XDG_PublicShare_DIR` envar will be supported:  /.PublicShare 
```

## Default path

The package defines defaults for XDG directories.

#### XDG Base Directory

| XDG Variable    | Mac OS                          | Windows                              | Others                           |
| :---            | :-----                          | :---                                 | :---                             |
| XDG_CONFIG_HOME | `~/Library/Preferences`         | `%LOCALAPPDATA%`                     | `~/.config`                      |
| XDG_CONFIG_DIRS | `/Library/Preferences`          | `%PROGRAMDATA%`                      | `/etc/xdg`                       |
| XDG_DATA_HOME   | `~/Library/Application Support` | `%LOCALAPPDATA%`                     | `~/.local/share`                 |
| XDG_DATA_DIRS   | `/Library/Application Support`  | `%APPDATA%\Roaming`, `%PROGRAMDATA%` | `/usr/local/share`, `/usr/share` |
| XDG_CACHE_HOME  | `~/Library/Caches`              | `%LOCALAPPDATA%\cache`               | `~/.cache`                       |
| XDG_RUNTIME_DIR | `~/Library/Application Support` | `%LOCALAPPDATA%`                     | `/run/user/UID`                  |

#### XDG User Directory

| XDG Variable        | Mac OS        | Windows                   | Others        |
| :---                | :-----        | :---                      | :---          |
| XDG_DESKTOP_DIR     | `~/Desktop`   | `%USERPROFILE%/Desktop`   | `~/Desktop`   |
| XDG_DOWNLOAD_DIR    | `~/Downloads` | `%USERPROFILE%/Downloads` | `~/Downloads` |
| XDG_DOCUMENTS_DIR   | `~/Documents` | `%USERPROFILE%/Documents` | `~/Documents` |
| XDG_PICTURES_DIR    | `~/Pictures`  | `%USERPROFILE%/Pictures`  | `~/Pictures`  |
| XDG_MUSIC_DIR       | `~/Music`     | `%USERPROFILE%/Music`     | `~/Music`     |
| XDG_VIDEOS_DIR      | `~/Movies`    | `%USERPROFILE%/Videos`    | `~/Videos`    |
| XDG_TEMPLATES_DIR   | `~/Templates` | `%USERPROFILE%/Templates` | `~/Templates` |
| XDG_PUBLICSHARE_DIR | `~/Public`    | `%PUBLIC%`                | `~/Public`    |

# LICENSE

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](http://www.opensource.org/licenses/MIT)

This is distributed under the [MIT License](http://www.opensource.org/licenses/MIT).
