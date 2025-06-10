package osutils

import (
	"io/fs"
	"os"
	"runtime"
)

type MachineOSInfo struct {
	Name     string
	RootPath string
}

var (
	Windows = &MachineOSInfo{Name: "windows", RootPath: `C:\`}
	Linux   = &MachineOSInfo{Name: "linux", RootPath: `/`}
	Darwin  = &MachineOSInfo{Name: "darwin", RootPath: `/`}
)

var (
	reg = map[string]*MachineOSInfo{
		Windows.Name: Windows,
		Linux.Name:   Linux,
		Darwin.Name:  Darwin,
	}
	defaultOS = Linux
)

func CurrentOS() *MachineOSInfo {
	if info := reg[runtime.GOOS]; info != nil {
		return info
	}
	return defaultOS
}

func RootFS() fs.FS {
	osInfo := CurrentOS()
	return os.DirFS(osInfo.RootPath)
}
