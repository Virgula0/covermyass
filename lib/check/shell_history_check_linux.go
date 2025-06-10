//go:build !windows

package check

import (
	"fmt"
)

type shellHistoryCheck struct{}

func NewShellHistoryCheck() Check {
	return &shellHistoryCheck{}
}

func (s *shellHistoryCheck) Name() string {
	return "shell_history_linux"
}

func (s *shellHistoryCheck) Paths() []string {
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

func init() {
	AddCheck(NewShellHistoryCheck())
}
