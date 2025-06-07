package check

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/sundowndev/covermyass/v2/utils"
)

type shellHistoryCheckLinux struct{}
type shellHistoryCheckWindows struct{}

func NewShellHistoryCheckLinux() Check {
	return &shellHistoryCheckLinux{}
}

func NewShellHistoryCheckWindows() Check {
	return &shellHistoryCheckWindows{}
}

func (s *shellHistoryCheckLinux) Name() string {
	return "shell_history_linux"
}

func (s *shellHistoryCheckWindows) Name() string {
	return "shell_history_windows"
}

func (s *shellHistoryCheckLinux) Paths() []string {
	var paths []string

	users := map[string]string{
		"root": "/root",
		"home": "/home/*",
	}

	suffixes := []string{
		"/.bash_history",
		"/.zsh_history",
		"/.node_repl_history",
		"/.python_history",
		"/.mysql_history",
	}

	for _, prefix := range users {
		for _, suf := range suffixes {
			paths = append(paths, fmt.Sprintf("%s%s", prefix, suf))
		}
	}
	return paths
}

func (s *shellHistoryCheckWindows) Paths() []string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logrus.Error(err)
		return []string{}
	}

	return []string{
		filepath.Join(homeDir, "AppData", "Roaming", "Microsoft", "Windows", "PowerShell", "PSReadLine", "ConsoleHost_history.txt"), // PowerShell
		filepath.Join(homeDir, ".node_repl_history"),
		filepath.Join(homeDir, ".python_history"),
		filepath.Join(homeDir, ".bash_history"), // git bash saves it
	}
}

func init() {
	// check windows
	if utils.CurrentOS() == utils.Windows {
		AddCheck(NewShellHistoryCheckWindows())
		return
	}

	// otherwise linux
	AddCheck(NewShellHistoryCheckLinux())
}
