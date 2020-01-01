// +build darwin

package xdg

import "os"

func DesktopDir() string     { return altHome(os.Getenv(DesktopDirEnv), "Desktop") }
func DownloadDir() string    { return altHome(os.Getenv(DownloadDirEnv), "Downloads") }
func DocumentsDir() string   { return altHome(os.Getenv(DocumentsDirEnv), "Documents") }
func MusicDir() string       { return altHome(os.Getenv(MusicDirEnv), "Music") }
func PicturesDir() string    { return altHome(os.Getenv(PicturesDirEnv), "Pictures") }
func VideosDir() string      { return altHome(os.Getenv(VideosDirEnv), "Movies") }
func TemplatesDir() string   { return altHome(os.Getenv(TemplatesDirEnv), "Templates") }
func PublicShareDir() string { return altHome(os.Getenv(PublicShareDirEnv), "Public") }
