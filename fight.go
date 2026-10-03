package main

import (
	"fmt"
	"math/rand"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type fightActionKind int

const (
	actionBroadside fightActionKind = iota
	actionSwivel
	actionCrew
	actionFlee
)

// fightActions mirrors Fight.playerTurn's dynamic menu from the Java version:
// Swivel/Crew only appear when the ship actually has them, Flee is always last.
func fightActions(ship Ship) ([]fightActionKind, []string) {
	kinds := []fightActionKind{actionBroadside}
	labels := []string{"Broadside Cannon - let loose with the big guns."}
	if ship.HasSwivel {
		kinds = append(kinds, actionSwivel)
		labels = append(labels, "Swivel Gun - a quicker, more precise shot.")
	}
	if ship.HasCrew {
		kinds = append(kinds, actionCrew)
		labels = append(labels, "Crew Ability - call on the crew for somethin' special.")
	}
	kinds = append(kinds, actionFlee)
	labels = append(labels, "Flee - try to lose 'em and slip away.")
	return kinds, labels
}

type fightStep int

const (
	stepFightIntro fightStep = iota
	stepChoosing
	stepRoundResult
	stepFled
	stepDead
)

type fightModel struct {
	state   *GameState
	actions []fightActionKind
	menu    menu
	step    fightStep
	outcome string
	fled    bool
}

func newFightModel(state *GameState) fightModel {
	kinds, labels := fightActions(state.Player)
	return fightModel{state: state, actions: kinds, menu: newMenu(labels), step: stepFightIntro}
}

func (m fightModel) Init() tea.Cmd {
	return nil
}

func (m fightModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	switch m.step {
	case stepFightIntro:
		m.step = stepChoosing
		return m, nil

	case stepChoosing:
		if !m.menu.HandleKey(keyMsg) {
			return m, nil
		}
		action := m.actions[m.menu.SelectedIndex()]
		m.outcome, m.fled = m.resolveRound(action)
		m.step = stepRoundResult
		return m, nil

	case stepRoundResult:
		if m.fled {
			m.step = stepFled
			return m, nil
		}
		if m.state.Player.Health <= 0 {
			m.step = stepDead
			return m, nil
		}
		kinds, labels := fightActions(m.state.Player)
		m.actions = kinds
		m.menu = newMenu(labels)
		m.step = stepChoosing
		return m, nil

	case stepFled:
		return m, func() tea.Msg { return screenDoneMsg{next: ScreenPirateBay} }

	case stepDead:
		return m, tea.Quit
	}

	return m, nil
}

// resolveRound mirrors Fight.run's per-round structure from the Java version:
// initiative decides who acts first, both act once per round (unless the
// player flees, or the player is already dead from the enemy's hit when the
// enemy went first - the same "zombie turn" guard the Java version needed).
func (m fightModel) resolveRound(action fightActionKind) (string, bool) {
	var lines []string
	var fled bool

	resolvePlayer := func() {
		switch action {
		case actionBroadside:
			lines = append(lines, m.playerBroadside())
		case actionSwivel:
			lines = append(lines, m.playerSwivel())
		case actionCrew:
			lines = append(lines, m.playerCrew())
		case actionFlee:
			var text string
			text, fled = m.attemptFlee()
			lines = append(lines, text)
		}
	}

	if initiative(m.state.Player.Speed, m.state.Enemy.Speed) {
		resolvePlayer()
		if !fled {
			lines = append(lines, m.enemyTurn())
		}
	} else {
		lines = append(lines, m.enemyTurn())
		if m.state.Player.Health > 0 {
			resolvePlayer()
		}
	}

	return strings.Join(lines, "\n"), fled
}

func (m fightModel) applyBandwidthCheck() string {
	if m.state.Player.Bandwidth <= 0 {
		m.state.Player.Health -= d20()
		return "We've stretched our bandwidth thin! Our ship be takin' damage!"
	}
	return ""
}

func withSelfDamage(line, selfDamageText string) string {
	if selfDamageText == "" {
		return line
	}
	return line + "\n" + selfDamageText
}

func (m fightModel) playerBroadside() string {
	ship := &m.state.Player
	var baseDamage, accuracyMod int
	switch ship.BroadsideType {
	case 1:
		baseDamage, accuracyMod = rand.Intn(16)+15, -3
	case 3:
		baseDamage, accuracyMod = rand.Intn(11)+5, 5
	default:
		baseDamage, accuracyMod = rand.Intn(11)+10, 0
	}

	ship.Bandwidth -= 10
	selfDamageText := m.applyBandwidthCheck()

	if resolveHit(ship.Accuracy, accuracyMod, m.state.Enemy.Dodge) {
		damage := int(float64(baseDamage) * ship.Damage)
		m.damageEnemy(damage)
		return withSelfDamage(fmt.Sprintf("Yer broadside cannons roar! Ye deal %d damage to the privateer!", damage), selfDamageText)
	}
	return withSelfDamage("Yer broadside cannons miss their mark!", selfDamageText)
}

func (m fightModel) playerSwivel() string {
	ship := &m.state.Player
	ship.Bandwidth -= 5
	selfDamageText := m.applyBandwidthCheck()

	var line string
	switch ship.SwivelType {
	case 1:
		if resolveHit(ship.Accuracy, 0, m.state.Enemy.Dodge) {
			damage := int(float64(rand.Intn(10)+1) * ship.Damage)
			m.damageEnemy(damage)
			m.state.Enemy.Accuracy -= 2
			line = fmt.Sprintf("Yer zip bomb swivel hits for %d damage and scrambles their targeting systems!", damage)
		} else {
			line = "Yer zip bomb swivel shot goes wide!"
		}
	case 2:
		if resolveHit(ship.Accuracy, 0, m.state.Enemy.Dodge) {
			damage := int(float64(rand.Intn(10)+1) * ship.Damage)
			m.damageEnemy(damage)
			ship.Accuracy += 2
			line = fmt.Sprintf("Yer IP-trace swivel hits for %d damage and sharpens yer own aim!", damage)
		} else {
			line = "Yer IP-trace swivel shot goes wide!"
		}
	case 3:
		damage := int(float64(rand.Intn(26)) * ship.Damage)
		m.damageEnemy(damage)
		line = fmt.Sprintf("Yer logic bomb swivel unleashes pure chaos, dealin' %d damage!", damage)
	default:
		line = "Ye haven't settled on a swivel gun load, matey!"
	}
	return withSelfDamage(line, selfDamageText)
}

func (m fightModel) playerCrew() string {
	ship := &m.state.Player
	ship.Bandwidth -= 15
	selfDamageText := m.applyBandwidthCheck()

	var line string
	switch ship.CrewType {
	case 1:
		if resolveHit(ship.Accuracy, 0, m.state.Enemy.Dodge) {
			damage := int(float64(rand.Intn(11)+8) * ship.Damage)
			m.damageEnemy(damage)
			m.state.Enemy.Speed -= 2
			line = fmt.Sprintf("Yer crew's musket volley hits for %d damage and slows their pursuit!", damage)
		} else {
			line = "Yer crew's musket volley misses!"
		}
	case 2:
		hit := resolveHit(ship.Accuracy, 0, m.state.Enemy.Dodge)
		ship.Armor *= 0.5
		if hit {
			damage := int(float64(rand.Intn(11)+10) * ship.Damage * 2)
			m.damageEnemy(damage)
			line = fmt.Sprintf("Yer crew overclocks the powder kegs! The blast hits for %d damage, but yer defenses take a hit from the strain!", damage)
		} else {
			line = "Yer crew overclocks the powder kegs, but the shot goes wide! Yer defenses still took a hit from the strain!"
		}
	default:
		line = "Ye haven't settled on a crew ability, matey!"
	}
	return withSelfDamage(line, selfDamageText)
}

func (m fightModel) damageEnemy(damage int) {
	m.state.Enemy.Health -= damage
	if m.state.Enemy.Health < 0 {
		m.state.Enemy.Health = 0
	}
}

// attemptFlee mirrors Fight.attemptFlee: suppressing the enemy enough (health
// at 0) guarantees escape, otherwise it's the same speed + d20 contest used
// for initiative. This is a flee-or-die fight, not a kill-the-enemy one - see
// the fight-system plan from earlier in the rewrite for why.
func (m fightModel) attemptFlee() (string, bool) {
	if m.state.Enemy.Health <= 0 {
		return "Their systems be fried from yer assault! Ye slip away unopposed, matey!", true
	}
	if initiative(m.state.Player.Speed, m.state.Enemy.Speed) {
		return "With a burst of speed, ye leave the privateer in yer wake! Ye've made yer escape!", true
	}
	return "Ye try to break away, but the privateer keeps pace! No luck this time, matey!", false
}

func (m fightModel) enemyTurn() string {
	attack := rand.Intn(6) + 1
	var name string
	var baseDamage int
	switch attack {
	case 1:
		name, baseDamage = "IP Tracker Swivel Gun", rand.Intn(11)+5
	case 2:
		name, baseDamage = "Firewall Broadside", rand.Intn(16)+10
	case 3:
		name, baseDamage = "Data Snare Chainshot", rand.Intn(11)+5
	case 4:
		name, baseDamage = "Proxy Buster Musket Volley", rand.Intn(11)+5
	case 5:
		name, baseDamage = "Encryption Jammer Grapeshot", rand.Intn(11)+8
	default:
		name, baseDamage = "Codebreaker Broadside", rand.Intn(16)+10
	}

	line := fmt.Sprintf("The privateer unleashes a %s!", name)

	if !resolveHit(m.state.Enemy.Accuracy, 0, m.state.Player.Dodge) {
		return line + "\nYe manage to avoid the attack!"
	}

	ship := &m.state.Player
	damage := mitigate(baseDamage, ship.Armor)
	ship.Health -= damage
	line += fmt.Sprintf("\nIt strikes true, dealin' %d damage!", damage)

	switch attack {
	case 1, 2:
		ship.Speed -= 2
		line += "\nYer ship's speed be hamperin' from the hit!"
	case 3:
		ship.Dodge -= 2
		line += "\nYer ship's maneuverability be hamperin' from the hit!"
	case 4, 5:
		ship.Armor -= 0.05
		line += "\nYer ship's defenses be weakenin' from the hit!"
	case 6:
		ship.Accuracy -= 2
		line += "\nYer ship's targeting systems be jammed from the hit!"
	}

	return line
}

// statsHeader shows both combatants' health side by side so the player can
// see at a glance how close they are to dying or to suppressing the enemy
// enough to guarantee an escape.
func (m fightModel) statsHeader() string {
	p, e := m.state.Player, m.state.Enemy
	return healthBar("You", p.Health, p.MaxHealth, 20) + "\n" +
		healthBar("Privateer", e.Health, e.MaxHealth, 20) + "\n\n"
}

func (m fightModel) View() string {
	var content string
	switch m.step {
	case stepFightIntro:
		content = "A shadow looms - an ominous privateer, relentless in their pursuit of buccaneers like ye!\n\n" +
			"Defeat could mean a life behind bars for yer digital transgressions.\n\n(press any key to continue)\n"
	case stepChoosing:
		content = m.statsHeader() + "What be yer move, captain?\n\n" + m.menu.View()
	case stepRoundResult:
		content = m.statsHeader() + m.outcome + "\n\n(press any key to continue)\n"
	case stepFled:
		content = "The privateer fades into the digital fog behind ye. Ye've made yer escape!\n\n(press any key to continue)\n"
	case stepDead:
		content = "\n\nYou are dead.\n\n*RE4 Leon death sound*\n"
	}
	return panel(content)
}
