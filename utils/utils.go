package utils

import (
	"fmt"
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
)

func ByteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func CurrentOS() *MachineOSInfo {
	if runtime.GOOS == Windows.Name {
		return Windows
	}
	return Linux
}

func RootFS() fs.FS {
	osInfo := CurrentOS()
	return os.DirFS(osInfo.RootPath)
}
