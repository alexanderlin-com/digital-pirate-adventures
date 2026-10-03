package main

import tea "github.com/charmbracelet/bubbletea"

type Screen int

const (
	ScreenIntro Screen = iota
	ScreenSetup
	ScreenShipSelect
	ScreenCustomize
	ScreenEmail
	ScreenStorm
	ScreenFight
	ScreenPirateBay
)

// screenDoneMsg is how a screen signals "I'm finished, advance to the next one."
// It's returned as a tea.Cmd from a screen's Update, which Bubble Tea then feeds
// back through RootModel.Update below, where it's intercepted and used to swap
// in the next screen's model.
type screenDoneMsg struct {
	next Screen
}

type RootModel struct {
	state   *GameState
	screen  Screen
	current tea.Model
}

func NewRootModel() RootModel {
	state := NewGameState()
	return RootModel{
		state:   state,
		screen:  ScreenIntro,
		current: newIntroModel(),
	}
}

func (m RootModel) Init() tea.Cmd {
	return m.current.Init()
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == "ctrl+c" {
		return m, tea.Quit
	}

	if done, ok := msg.(screenDoneMsg); ok {
		m.screen = done.next
		switch done.next {
		case ScreenSetup:
			m.current = newSetupModel(m.state)
		case ScreenShipSelect:
			m.current = newShipSelectModel(m.state)
		case ScreenCustomize:
			m.current = newCustomizeModel(m.state)
		case ScreenEmail:
			m.current = newEmailModel(m.state)
		case ScreenStorm:
			m.current = newStormModel(m.state)
		case ScreenFight:
			m.current = newFightModel(m.state)
		case ScreenPirateBay:
			m.current = newPirateBayModel(m.state)
		}
		return m, m.current.Init()
	}

	updated, cmd := m.current.Update(msg)
	m.current = updated
	return m, cmd
}

func (m RootModel) View() string {
	return m.current.View()
}
