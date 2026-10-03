package main

import (
	"fmt"
	"math/rand"

	tea "github.com/charmbracelet/bubbletea"
)

type stormStep int

const (
	stepStormIntro stormStep = iota
	stepLightningChoice
	stepLightningOutcome
	stepWhirlwindChoice
	stepWhirlwindOutcome
	stepEruptionChoice
	stepEruptionOutcome
	stepWaveChoice
	stepWaveOutcome
	stepStormEnd
)

var lightningOptions = []string{
	"Ready the VPN defenses and face the lightning head-on!",
	"Dodge the strike with a quick maneuver.",
	"Brave the storm and sail straight through.",
}

var whirlwindOptions = []string{
	"Ready the swivel gun and blast a path through.",
	"Navigate with caution, takin' it slow.",
	"Analyze the data patterns for a shortcut.",
}

// Eruption and wave are ported faithfully from the Java Storm.eruption/wave: all
// three choices there are flavor-only (every switch case was an empty `break;`,
// only the bad-input default did anything), so they stay flavor-only here too.
var eruptionOptions = []string{
	"Search for a weakness in the firewall.",
	"Blast through with the DDOS cannon.",
	"Wait for the firewall to dissipate naturally.",
}

var waveOptions = []string{
	"Raise encryption defenses against the wave.",
	"Ride the wave for a speed boost.",
	"Slacken the sails and weather it with caution.",
}

type stormModel struct {
	state       *GameState
	step        stormStep
	menu        menu
	lastOutcome string
}

func newStormModel(state *GameState) stormModel {
	return stormModel{state: state, step: stepStormIntro}
}

func (m stormModel) Init() tea.Cmd {
	return nil
}

func (m stormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	switch m.step {
	case stepStormIntro:
		m.step = stepLightningChoice
		m.menu = newMenu(lightningOptions)

	case stepLightningChoice:
		if m.menu.HandleKey(keyMsg) {
			m.lastOutcome = m.resolveLightning(m.menu.SelectedIndex())
			m.step = stepLightningOutcome
		}
	case stepLightningOutcome:
		m.step = stepWhirlwindChoice
		m.menu = newMenu(whirlwindOptions)

	case stepWhirlwindChoice:
		if m.menu.HandleKey(keyMsg) {
			m.lastOutcome = m.resolveWhirlwind(m.menu.SelectedIndex())
			m.step = stepWhirlwindOutcome
		}
	case stepWhirlwindOutcome:
		m.step = stepEruptionChoice
		m.menu = newMenu(eruptionOptions)

	case stepEruptionChoice:
		if m.menu.HandleKey(keyMsg) {
			m.lastOutcome = "Ye make yer move, but the firewall's defenses hold firm for now - naught to be done but press on."
			m.step = stepEruptionOutcome
		}
	case stepEruptionOutcome:
		m.step = stepWaveChoice
		m.menu = newMenu(waveOptions)

	case stepWaveChoice:
		if m.menu.HandleKey(keyMsg) {
			m.lastOutcome = "The wave crashes past, leavin' the ship a little worse for wear but still afloat."
			m.step = stepWaveOutcome
		}
	case stepWaveOutcome:
		m.step = stepStormEnd

	case stepStormEnd:
		return m, func() tea.Msg { return screenDoneMsg{next: ScreenFight} }
	}

	return m, nil
}

// resolveLightning ports Storm.lightning from the Java version faithfully,
// including its slightly odd case-3 damage formula.
func (m stormModel) resolveLightning(choice int) string {
	ship := &m.state.Player
	switch choice {
	case 0:
		score := d20()
		if ship.SailType == 1 {
			score += 5
		}
		if score > 11 {
			return "Our defenses held strong, and the lightning be naught but a spark!"
		}
		ship.Health -= rand.Intn(25)
		return "Alas! The lightning strikes true, and we've taken damage!"
	case 1:
		score := d20() + ship.Dodge/2
		if score > 11 {
			return "Our ship be nimble as the tide and avoids the lightning!"
		}
		ship.Health -= rand.Intn(25)
		return "Alas! The lightning strikes true, and we've taken damage!"
	default:
		damage := rand.Float64()
		ship.Health -= int(damage - damage*(ship.Armor*2))
		return "We've braced for the impact and most of it be negated!"
	}
}

func (m stormModel) resolveWhirlwind(choice int) string {
	ship := &m.state.Player
	switch choice {
	case 0:
		score := d20()
		if ship.HasSwivel {
			score += 5
		}
		if score > 12 {
			return "With a thunderous roar, we blasted the whirlwind clear!"
		}
		ship.Health -= rand.Intn(25)
		return "The whirlwind struck back, dealin' damage to our brave vessel!"
	case 1:
		score := d20()
		ship.Bandwidth -= rand.Intn(25)
		if score >= 5 {
			return "We skirted the whirlwind, but it cost us some bandwidth."
		}
		ship.Health -= rand.Intn(25)
		return "We took it slow, yet the whirlwind still found us!"
	default:
		score := d20() + (ship.Dodge+ship.Speed)/3
		if score >= 10 {
			return "Avast! We spied a shortcut through the whirlwind!"
		}
		ship.Health -= rand.Intn(25)
		return "We tried to cut through, but still it struck us! Cursed luck!"
	}
}

func (m stormModel) View() string {
	var content string
	switch m.step {
	case stepStormIntro:
		content = "Ye be sailin' onward, but a fearsome Bitstorm be upon us!\n\n(press any key to continue)\n"
	case stepLightningChoice:
		content = "A Data Surge Lightning Strike be on the horizon! What be yer course, captain?\n\n" + m.menu.View()
	case stepLightningOutcome, stepWhirlwindOutcome, stepEruptionOutcome, stepWaveOutcome:
		content = m.lastOutcome + "\n\n(press any key to continue)\n"
	case stepWhirlwindChoice:
		content = "A data debris whirlwind be upon us! What be our course, captain?\n\n" + m.menu.View()
	case stepEruptionChoice:
		content = "A firewall eruption be upon us! What be yer orders, captain?\n\n" + m.menu.View()
	case stepWaveChoice:
		content = "A monster data packet wave approaches! What be yer orders, captain?\n\n" + m.menu.View()
	case stepStormEnd:
		p := m.state.Player
		content = fmt.Sprintf("The Bitstorm be behind us, me hearties!\n\nFinal stats for %s aboard %s:\n\n%s\n%s\n\n(press any key to continue)\n",
			m.state.PlayerName, m.state.ShipName,
			healthBar("Health", p.Health, p.MaxHealth, 20),
			healthBar("Bandwidth", p.Bandwidth, p.MaxBandwidth, 20))
	}
	return panel(content)
}
