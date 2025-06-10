//go:build !linux && !darwin

package check

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

type shellHistoryCheck struct{}

func NewShellHistoryCheck() Check {
	return &shellHistoryCheck{}
}

func (s *shellHistoryCheck) Name() string {
	return "shell_history_windows"
}

func (s *shellHistoryCheck) Paths() []string {
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
	AddCheck(NewShellHistoryCheck())
}
