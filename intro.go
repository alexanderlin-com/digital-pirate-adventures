package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const introText = `Ahoy there fellow digital pirate!
Welcome aboard yer trust ship and prepare to sail the internet seas!
I be yer trusty virtual quartermaster.`

type introModel struct {
	badInput bool
}

func newIntroModel() introModel {
	return introModel{}
}

func (m introModel) Init() tea.Cmd {
	return nil
}

func (m introModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	switch keyMsg.String() {
	case "y":
		return m, func() tea.Msg { return screenDoneMsg{next: ScreenSetup} }
	case "n":
		return m, tea.Quit
	default:
		m.badInput = true
		return m, nil
	}
}

var titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("214"))

func (m introModel) View() string {
	view := banner() + "\n"
	view += titleStyle.Render(introText) + "\n\n"
	view += "Ready to embark on this digital adventure? (y/n)\n"
	if m.badInput {
		view += "\nArr, I didn't quite catch that. Try 'y' for yes or 'n' for no, matey!\n"
	}
	return panel(view)
}
