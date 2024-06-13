package generate

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	message        string
	errorInterface error
}

var (
	ERROR_STYLE   = lipgloss.NewStyle().AlignHorizontal(lipgloss.Left).Foreground(lipgloss.Color("#d32f2f")).Padding(1)
	SUCCESS_STYLE = lipgloss.NewStyle().AlignHorizontal(lipgloss.Left).Foreground(lipgloss.Color("#4caf50")).Padding(1)
)

func NewModel() model {
	return model{}
}

func DisplayError(errorInterface error) error {
	p := tea.NewProgram(model{errorInterface: errorInterface})
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}
	return nil
}

func DisplayMessage(message string) error {
	p := tea.NewProgram(model{message: message})
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}
	return nil
}

func (m model) Init() tea.Cmd {
	return tea.Quit
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.errorInterface != nil {
		return ERROR_STYLE.Render(m.errorInterface.Error()) + "\n"
	} else {
		return SUCCESS_STYLE.Render(m.message) + "\n"
	}
}
