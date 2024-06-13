package version

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/common-nighthawk/go-figure"
)

type model struct {
	message string
}

func FromMessage(message string) model {
	return model{
		message,
	}
}

func Display(m model) error {
	p := tea.NewProgram(m)
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
	asciiArt := figure.NewFigure(m.message, "larry3d", true).String()

	style := lipgloss.NewStyle().
		Bold(true).
		AlignHorizontal(lipgloss.Center).
		Foreground(lipgloss.Color("#0EA5E9")).
		Padding(2)

	return asciiArt + "\n" + style.Render("Build Simple Web Apps using Go,Tailwind,Templ and HTMX") + "\n"
}
