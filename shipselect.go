package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type shipSelectModel struct {
	state     *GameState
	menu      menu
	confirmed bool
}

var shipOptions = []string{
	"Sloop   - fastest speed, lowest firepower",
	"Frigate - balanced speed n' firepower",
	"Galleon - superior firepower, slowest speed",
}

func newShipSelectModel(state *GameState) shipSelectModel {
	return shipSelectModel{state: state, menu: newMenu(shipOptions)}
}

func (m shipSelectModel) Init() tea.Cmd {
	return nil
}

func (m shipSelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	if m.confirmed {
		return m, func() tea.Msg { return screenDoneMsg{next: ScreenCustomize} }
	}

	if m.menu.HandleKey(keyMsg) {
		kind := ShipKind(m.menu.SelectedIndex())
		m.state.Player = NewShip(kind)
		m.confirmed = true
	}

	return m, nil
}

var summaryStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("214"))

func (m shipSelectModel) View() string {
	if m.confirmed {
		p := m.state.Player
		summary := summaryStyle.Render(fmt.Sprintf("Take %s ahead full, Captain %s!", m.state.ShipName, m.state.PlayerName)) + "\n\n"
		summary += fmt.Sprintf("Health %d  Speed %d  Dodge %d\n\n", p.Health, p.Speed, p.Dodge)
		summary += "(press any key to continue)\n"
		return panel(summary)
	}

	view := "What be the type o' yer ship, matey?\n\n"
	view += m.menu.View()
	return panel(view)
}
