package tui

import (
	"os/exec"
	"runtime"

	tea "github.com/charmbracelet/bubbletea"
)

type repositoryOpenMsg struct {
	err error
}

var openRepositoryURLCmd = defaultOpenRepositoryURLCmd

func defaultOpenRepositoryURLCmd() tea.Cmd {
	return func() tea.Msg {
		var cmd *exec.Cmd
		switch runtime.GOOS {
		case "darwin":
			cmd = exec.Command("open", repositoryFooterURL)
		case "windows":
			cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", repositoryFooterURL)
		default:
			cmd = exec.Command("xdg-open", repositoryFooterURL)
		}

		return repositoryOpenMsg{err: cmd.Run()}
	}
}
