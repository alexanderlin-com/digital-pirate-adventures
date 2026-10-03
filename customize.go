package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type customizeStep int

const (
	stepBroadside customizeStep = iota
	stepSwivel
	stepCrew
	stepSails
	stepRigging
)

var broadsideOptions = []string{
	"DDOS Cannon - Devastatin' damage, but accuracy be lackin'.",
	"DNS Leak Launchers - Medium damage n' accuracy; a safe choice.",
	"Zero Day Exploit Guns - High accuracy, but lower damage.",
}

var swivelOptions = []string{
	"Zip Bomb Swivels - Lower damage, confuses enemy accuracy.",
	"IP-Trace Swivels - Lower damage, boosts yer own accuracy.",
	"Logic Bomb Swivels - Pure RNG chaos.",
}

var crewOptions = []string{
	"Latency Musket Volley - Damage and disrupts enemy systems.",
	"Overclock Powder Kegs - Big damage, sacrifices armor.",
}

var sailsOptions = []string{
	"VPN Sails - More maneuverability, less speed.",
	"Fiber Optic Sails - Pure speed.",
}

var riggingOptions = []string{
	"Anti-virus Rigging - More defense, less speed.",
	"Tune-up Utility Rigging - Pure speed.",
}

func optionsFor(step customizeStep) []string {
	switch step {
	case stepBroadside:
		return broadsideOptions
	case stepSwivel:
		return swivelOptions
	case stepCrew:
		return crewOptions
	case stepSails:
		return sailsOptions
	case stepRigging:
		return riggingOptions
	}
	return nil
}

func promptFor(step customizeStep) string {
	switch step {
	case stepBroadside:
		return "Yer broadside cannons be yer main attack. What type o' cannon be yer choice?"
	case stepSwivel:
		return "Yer swivel guns be a fine secondary gun. What be yer choice o' swivel?"
	case stepCrew:
		return "Our ship be havin' room for extra crew abilities. What be yer choice?"
	case stepSails:
		return "Our ship's sails be the lifeblood o' our vessel. What kind o' sails be on yer ship?"
	case stepRigging:
		return "Our ship not be complete without proper riggin'. What kind o' riggin' be on yer ship?"
	}
	return ""
}

// customizeModel's step sequence is driven by the ship's HasSwivel/HasCrew flags
// rather than hardcoded per ship type, unlike the Java version where each
// Sloop/Frigate/Galleon subclass separately hardcoded its own customize() sequence.
type customizeModel struct {
	state *GameState
	steps []customizeStep
	index int
	menu  menu
}

func newCustomizeModel(state *GameState) customizeModel {
	steps := []customizeStep{stepBroadside}
	if state.Player.HasSwivel {
		steps = append(steps, stepSwivel)
	}
	if state.Player.HasCrew {
		steps = append(steps, stepCrew)
	}
	steps = append(steps, stepSails, stepRigging)

	return customizeModel{
		state: state,
		steps: steps,
		menu:  newMenu(optionsFor(steps[0])),
	}
}

func (m customizeModel) Init() tea.Cmd {
	return nil
}

func (m customizeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	// A screenDoneMsg command is async - a stray keypress can still reach this
	// same model before the swap to the next screen actually happens.
	if m.index >= len(m.steps) {
		return m, nil
	}

	if !m.menu.HandleKey(keyMsg) {
		return m, nil
	}

	m.apply(m.steps[m.index], m.menu.SelectedIndex())

	m.index++
	if m.index >= len(m.steps) {
		return m, func() tea.Msg { return screenDoneMsg{next: ScreenEmail} }
	}

	m.menu = newMenu(optionsFor(m.steps[m.index]))
	return m, nil
}

// apply records the chosen upgrade type and, for sails/rigging, applies the
// immediate stat change the Java Ship.setSails/setRigging applied right away
// (broadside/swivel/crew effects are deferred to when they're actually used
// in a fight, same as the Java version).
func (m customizeModel) apply(step customizeStep, choice int) {
	ship := &m.state.Player
	switch step {
	case stepBroadside:
		ship.BroadsideType = choice + 1
	case stepSwivel:
		ship.SwivelType = choice + 1
	case stepCrew:
		ship.CrewType = choice + 1
	case stepSails:
		ship.SailType = choice + 1
		if choice == 0 {
			ship.Dodge += 2
			ship.Speed -= 2
		} else {
			ship.Speed += 2
		}
	case stepRigging:
		ship.RiggingType = choice + 1
		if choice == 0 {
			ship.Armor += 0.1
			ship.Speed -= 2
		} else {
			ship.Speed += 2
		}
	}
}

func (m customizeModel) View() string {
	// Update can advance m.index to len(m.steps) and return a screenDoneMsg
	// command without having swapped the screen yet - that command is async,
	// so View() still gets called once more on this same model in the meantime.
	if m.index >= len(m.steps) {
		return ""
	}
	return panel(promptFor(m.steps[m.index]) + "\n\n" + m.menu.View())
}
