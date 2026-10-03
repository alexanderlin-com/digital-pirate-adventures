package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// menu is a small reusable arrow-key-selectable list. Every later numbered-choice
// screen (storm events, fight actions, Pirate Bay) reuses this instead of each
// hand-rolling its own input handling the way every Ship.setX/Main.askX method
// did in the Java version.
type menu struct {
	options []string
	cursor  int
}

func newMenu(options []string) menu {
	return menu{options: options}
}

func (m *menu) up() {
	if m.cursor > 0 {
		m.cursor--
	}
}

func (m *menu) down() {
	if m.cursor < len(m.options)-1 {
		m.cursor++
	}
}

func (m *menu) SelectedIndex() int {
	return m.cursor
}

// HandleKey applies arrow-key navigation for a tea.KeyMsg and reports whether
// Enter was pressed (i.e. a selection was confirmed) this call.
func (m *menu) HandleKey(msg tea.KeyMsg) (confirmed bool) {
	switch msg.String() {
	case "up", "k":
		m.up()
	case "down", "j":
		m.down()
	case "enter":
		return true
	}
	return false
}

var (
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true)
	optionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("250"))
)

func (m *menu) View() string {
	view := ""
	for i, option := range m.options {
		if i == m.cursor {
			view += cursorStyle.Render("> "+option) + "\n"
		} else {
			view += optionStyle.Render("  "+option) + "\n"
		}
	}
	return view
}
