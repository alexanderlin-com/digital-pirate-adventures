package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type pirateBayStep int

const (
	stepBayIntro pirateBayStep = iota
	stepBayOutcome
	stepBayChoice
	stepBayEndingText
	stepBayClose
)

var nextPlanOptions = []string{
	"Explore the Pirate Bay - be on the lookout for more treasures to plunder.",
	"Engage in Digital Trade - seed the new game torrent for fellow buccaneers.",
	"Set Sail for New Adventures - uncharted binary seas await.",
	"Venture to Treacherous Waters - upgrade with a Tor hull, risk the darkweb depths.",
	"Go Legit - leave piracy behind for a lawful, calmer digital life.",
}

type pirateBayModel struct {
	state  *GameState
	step   pirateBayStep
	menu   menu
	ending int
}

func newPirateBayModel(state *GameState) pirateBayModel {
	return pirateBayModel{state: state, step: stepBayIntro}
}

func (m pirateBayModel) Init() tea.Cmd {
	return nil
}

func (m pirateBayModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	switch m.step {
	case stepBayIntro:
		m.step = stepBayOutcome
	case stepBayOutcome:
		m.step = stepBayChoice
		m.menu = newMenu(nextPlanOptions)
	case stepBayChoice:
		if !m.menu.HandleKey(keyMsg) {
			return m, nil
		}
		m.ending = m.menu.SelectedIndex()
		m.step = stepBayEndingText
	case stepBayEndingText:
		m.step = stepBayClose
	case stepBayClose:
		return m, tea.Quit
	}

	return m, nil
}

// outcomeText mirrors the three branches in Encounter.pirateBay from the Java
// version, gated on the library/textbook choices made back in the email screen.
func (m pirateBayModel) outcomeText() string {
	if m.state.VisitedLibrary {
		if m.state.HasTextbook {
			return "Arr, ye've secured the coveted torrent ye sought, and ye've made a true pirate's promise to share the class tome with yer matey.\n\n" +
				"A fine haul and a loyal heart - the high binary seas have treated ye well, me heartie!"
		}
		return "Ye've got yer hands on the prized torrent, but ye missed the chance to honor the request of a fellow pirate.\n\n" +
			"Fear not, there be more chances ahead, and ye'll keep that promise in the next digital voyage."
	}
	return "Though the textbook slipped through yer grasp this time, ye've captured the digital treasure ye set out for.\n\n" +
		"With determination in yer heart, ye'll continue the search and ensure yer matey gets their hands on the class tome soon."
}

// endingText ports the five Ending.endingN methods from the Java version.
func (m pirateBayModel) endingText() string {
	switch m.ending {
	case 0:
		return "With a sense of curiosity and adventure, ye embark on a quest to explore the vast archives of The Pirate Bay.\n\n" +
			"The binary seas be teemin' with treasures, and ye discover torrents of other games, software, and digital wonders.\n\n" +
			"The Pirate Bay becomes yer home, and yer legend as a digital pirate grows with each new treasure ye uncover."
	case 1:
		return "Ye commit to the code of digital trade, downloadin' the coveted new game torrent and dedicatin' yer time to seedin' it for yer fellow pirates.\n\n" +
			"Yer generosity doesn't go unnoticed, and ye earn the respect and admiration of yer pirate peers."
	case 2:
		return "Yer heart longs for new adventures and uncharted waters, and ye set sail toward the distant horizon.\n\n" +
			"The binary seas stretch out before ye, vast and full of mysteries waitin' to be unraveled."
	case 3:
		return fmt.Sprintf("With darin' courage, ye decide to venture into the treacherous, shadowy waters of the dark web, upgradin' %s with a reinforced Tor hull.\n\n"+
			"Yer name strikes fear into the hearts of those who dwell in the dark, and ye become a legend of the high binary seas.", m.state.ShipName)
	default:
		return "Ye make the bold choice to leave the life of piracy behind and sail a more lawful path.\n\n" +
			"While ye may no longer be a pirate, ye navigate the digital realm with wisdom and integrity, carvin' out a respectable and lawful existence."
	}
}

func (m pirateBayModel) View() string {
	var content string
	switch m.step {
	case stepBayIntro:
		content = "Yer vessel finally drops anchor at The Pirate Bay, a fabled sanctuary fer buccaneers.\n\n(press any key to continue)\n"
	case stepBayOutcome:
		content = m.outcomeText() + "\n\n(press any key to continue)\n"
	case stepBayChoice:
		content = "What be your next plan?\n\n" + m.menu.View()
	case stepBayEndingText:
		content = m.endingText() + "\n\n(press any key to continue)\n"
	case stepBayClose:
		content = fmt.Sprintf("This be the outcomes of yer adventures as %s, Captain of %s!\n\n"+
			"Yer choices have shaped the course of yer digital piracy tale, and yer legacy be etched into the annals of the binary seas!\n\n"+
			"(press any key to exit)\n", m.state.PlayerName, m.state.ShipName)
	}
	return panel(content)
}
