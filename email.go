package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type emailStep int

const (
	stepPassphrase emailStep = iota
	stepLibOrSea
	stepLibraryIntro
	stepLibraryResponse
)

var libOrSeaOptions = []string{
	"Go to the Library",
	"Sail straight to The Pirate Bay",
}

type emailModel struct {
	state *GameState
	step  emailStep
	input textinput.Model
	menu  menu
}

func newEmailModel(state *GameState) emailModel {
	ti := textinput.New()
	ti.Placeholder = "yer secret passphrase"
	ti.Focus()
	return emailModel{state: state, step: stepPassphrase, input: ti, menu: newMenu(libOrSeaOptions)}
}

func (m emailModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m emailModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	switch m.step {
	case stepPassphrase:
		if keyMsg.String() == "enter" {
			m.step = stepLibOrSea
			return m, nil
		}
		var cmd tea.Cmd
		m.input, cmd = m.input.Update(msg)
		return m, cmd

	case stepLibOrSea:
		if m.menu.HandleKey(keyMsg) {
			if m.menu.SelectedIndex() == 0 {
				m.state.VisitedLibrary = true
				m.step = stepLibraryIntro
			} else {
				m.state.VisitedLibrary = false
				return m, func() tea.Msg { return screenDoneMsg{next: ScreenStorm} }
			}
		}
		return m, nil

	case stepLibraryIntro:
		if keyMsg.String() == "enter" {
			m.step = stepLibraryResponse
		}
		return m, nil

	case stepLibraryResponse:
		switch keyMsg.String() {
		case "y":
			m.state.HasTextbook = true
			return m, func() tea.Msg { return screenDoneMsg{next: ScreenStorm} }
		case "n":
			m.state.HasTextbook = false
			return m, func() tea.Msg { return screenDoneMsg{next: ScreenStorm} }
		}
		return m, nil
	}

	return m, nil
}

func (m emailModel) View() string {
	var content string
	switch m.step {
	case stepPassphrase:
		content = "Ahoy, it seems there be a fresh scroll in the inbox!\n\n" +
			"Enter yer secret passphrase to unfurl the message, matey:\n\n" + m.input.View() + "\n"
	case stepLibOrSea:
		content = "[Hey what's up, the torrent file for that new game just got released this morning.\n" +
			"ps if you get the chance could you send me the textbook for the class? ty.]\n\n" +
			"Hmm, it be lookin' like we ought to make a course for The Pirate Bay.\n" +
			"But mayhaps we should pay a visit to the library if fortune smiles upon us.\n\n" + m.menu.View()
	case stepLibraryIntro:
		content = "Arrr! Ye find yerself at the legendary Library Genesis, a digital treasure trove.\n\n" +
			"A salty dog named Keshav ambles over and extends a digital textbook to ye.\n\n" +
			"(press enter to continue)\n"
	case stepLibraryResponse:
		content = "Be ye claimin' this tome, matey? (y/n)\n"
	}
	return panel(content)
}
