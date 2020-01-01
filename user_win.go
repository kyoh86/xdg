// +build windows

package xdg

import "os"

func DesktopDir() string     { return altUserProfile(os.Getenv(DesktopDirEnv), "Desktop") }
func DownloadDir() string    { return altUserProfile(os.Getenv(DownloadDirEnv), "Downloads") }
func DocumentsDir() string   { return altUserProfile(os.Getenv(DocumentsDirEnv), "Documents") }
func MusicDir() string       { return altUserProfile(os.Getenv(MusicDirEnv), "Music") }
func PicturesDir() string    { return altUserProfile(os.Getenv(PicturesDirEnv), "Pictures") }
func VideosDir() string      { return altUserProfile(os.Getenv(VideosDirEnv), "Videos") }
func TemplatesDir() string   { return altUserProfile(os.Getenv(TemplatesDirEnv), "Templates") }
func PublicShareDir() string { return alternate(os.Getenv(PublicShareDirEnv), os.Getenv("PUBLIC")) }
