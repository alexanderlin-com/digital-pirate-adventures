package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type setupStep int

const (
	stepPlayerName setupStep = iota
	stepShipName
)

type setupModel struct {
	state *GameState
	step  setupStep
	input textinput.Model
}

func newSetupModel(state *GameState) setupModel {
	ti := textinput.New()
	ti.Placeholder = "yer name, matey"
	ti.Focus()
	return setupModel{state: state, step: stepPlayerName, input: ti}
}

func (m setupModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m setupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == "enter" {
		value := m.input.Value()
		if value == "" {
			return m, nil
		}

		switch m.step {
		case stepPlayerName:
			m.state.PlayerName = value
			m.step = stepShipName
			m.input.Reset()
			m.input.Placeholder = "yer ship's name"
			return m, nil
		case stepShipName:
			m.state.ShipName = value
			return m, func() tea.Msg { return screenDoneMsg{next: ScreenShipSelect} }
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m setupModel) View() string {
	var prompt string
	switch m.step {
	case stepPlayerName:
		prompt = "Ahoy! Before we set sail, what be yer name, matey?"
	case stepShipName:
		prompt = "And what be the name of our ship?"
	}
	return panel(prompt + "\n\n" + m.input.View() + "\n")
}
