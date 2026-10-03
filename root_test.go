package main

import (
	"bytes"
	"io"
	"regexp"
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/exp/teatest"
)

var ansiEscape = regexp.MustCompile(`\x1b\[[0-9;?]*[a-zA-Z]`)

// normalizeOutput strips ANSI escape codes and the panel's rounded-border
// glyphs, then collapses all whitespace (including the newlines a long line
// gets wrapped across inside the fixed-width panel). Matching against this
// instead of the raw bytes means test phrases don't break just because a
// sentence happens to wrap across a border line at the panel's text width -
// which is correct, desirable rendering behavior, not something worth
// hand-picking "safe" phrases to dodge.
func normalizeOutput(b []byte) string {
	s := ansiEscape.ReplaceAllString(string(b), "")
	s = strings.NewReplacer("│", " ", "─", " ", "╭", " ", "╮", " ", "╰", " ", "╯", " ").Replace(s)
	return strings.Join(strings.Fields(s), " ")
}

// TestFullSlice drives the entire flow built so far - intro, setup, ship
// selection (Sloop, for the shortest customize path), customization, the
// email/library branch (skipping the library for a shorter path), and all
// four storm events through to the end - with synthetic key messages, and
// asserts on the final rendered output. This is the only practical way to
// verify a Bubble Tea program end-to-end without a live TTY.
//
// Waits for each screen's expected text before sending the next key, rather
// than firing input rapid-fire: screen transitions happen via an async
// tea.Cmd, so racing ahead of it is how the customize-screen index-out-of-
// range panic surfaced during development.
//
// Uses a hand-rolled poll instead of teatest.WaitFor: that helper drains the
// same underlying buffer it reads from on every call, so once one waitFor
// call consumes the bytes containing its expected text, a later call has no
// way to see that same text again - only bytes written since the last read
// by anyone are visible. Accumulating everything read across the whole test
// into one running buffer sidesteps that.
func TestFullSlice(t *testing.T) {
	tm := teatest.NewTestModel(t, NewRootModel(), teatest.WithInitialTermSize(100, 40))

	var seen bytes.Buffer
	waitFor := func(text string) {
		t.Helper()
		deadline := time.Now().Add(2 * time.Second)
		for {
			b, _ := io.ReadAll(tm.Output())
			seen.Write(b)
			if strings.Contains(normalizeOutput(seen.Bytes()), text) {
				return
			}
			if time.Now().After(deadline) {
				t.Fatalf("waitFor %q: condition not met after 2s. Seen so far:\n%s", text, seen.String())
			}
			time.Sleep(20 * time.Millisecond)
		}
	}

	enter := func() { tm.Send(tea.KeyMsg{Type: tea.KeyEnter}) }
	runes := func(s string) { tm.Send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(s)}) }

	waitFor("Ready to embark")
	runes("y") // intro -> setup

	waitFor("what be yer name")
	runes("TestCaptain")
	enter() // -> ship name step

	waitFor("name of our ship")
	runes("TestShip")
	enter() // -> ship select

	waitFor("What be the type o' yer ship")
	enter() // select Sloop (index 0) -> summary

	waitFor("Take TestShip ahead full")
	enter() // acknowledge summary -> customize

	waitFor("broadside cannons")
	enter() // broadside choice

	// Fiber Optic Sails / Tune-up Utility Rigging (index 1 for both) are the
	// speed-boosting options, keeping Sloop's speed advantage over the enemy
	// intact for the flee checks in the fight below - VPN Sails/Anti-virus
	// Rigging (index 0) both subtract speed instead, which brought a test
	// Sloop down to parity with the enemy and made fleeing an unreliable
	// coin flip, risking death during the retry loop below.
	waitFor("sails be the lifeblood")
	tm.Send(tea.KeyMsg{Type: tea.KeyDown})
	enter() // Fiber Optic Sails

	waitFor("proper riggin'")
	tm.Send(tea.KeyMsg{Type: tea.KeyDown})
	enter() // Tune-up Utility Rigging -> email

	waitFor("secret passphrase")
	enter() // submit passphrase -> lib-or-sea

	waitFor("Go to the Library")
	tm.Send(tea.KeyMsg{Type: tea.KeyDown}) // move to "Pirate Bay"
	enter()                                // confirm Pirate Bay -> storm

	waitFor("Bitstorm be upon us")
	enter() // storm intro -> lightning choice

	waitFor("Lightning Strike")
	enter() // lightning choice -> outcome

	waitFor("continue")
	enter() // outcome -> whirlwind choice

	waitFor("whirlwind be upon us")
	enter() // whirlwind choice -> outcome

	waitFor("continue")
	enter() // outcome -> eruption choice

	waitFor("firewall eruption")
	enter() // eruption choice -> outcome

	waitFor("continue")
	enter() // outcome -> wave choice

	waitFor("data packet wave")
	enter() // wave choice -> outcome

	waitFor("continue")
	enter() // outcome -> storm end (summary)

	// The storm-end summary includes the player/ship names from GameState,
	// proving they threaded correctly through every screen since setup.
	waitFor("Bitstorm be behind us")
	waitFor("TestCaptain")
	waitFor("TestShip")
	enter() // storm end -> fight

	waitFor("ominous privateer")
	enter() // fight intro -> choosing

	waitFor("What be yer move")

	// Fleeing is genuinely probabilistic (an initiative/speed contest), so
	// retry until it succeeds rather than assuming a fixed number of tries.
	// Each attempt's result is checked only against output written since
	// that attempt started - waitFor's forever-accumulating buffer can't be
	// reused here, since "No luck this time" from a failed attempt would
	// otherwise still match on every later attempt too.
	fleeUntilEscaped := func(maxAttempts int) {
		t.Helper()
		for i := 0; i < maxAttempts; i++ {
			startLen := seen.Len()
			tm.Send(tea.KeyMsg{Type: tea.KeyDown}) // Sloop's menu is [Broadside, Flee]
			enter()

			deadline := time.Now().Add(2 * time.Second)
			for {
				b, _ := io.ReadAll(tm.Output())
				seen.Write(b)
				fresh := normalizeOutput(seen.Bytes()[startLen:])
				if strings.Contains(fresh, "made yer escape") || strings.Contains(fresh, "unopposed") {
					return
				}
				if strings.Contains(fresh, "No luck this time") {
					enter() // back to choosing for the next attempt
					break
				}
				if strings.Contains(fresh, "You are dead") {
					t.Fatalf("flee attempt %d: player died before escaping. Fresh output:\n%s", i, fresh)
				}
				if time.Now().After(deadline) {
					t.Fatalf("flee attempt %d: no result text appeared. Fresh output:\n%s", i, fresh)
				}
				time.Sleep(20 * time.Millisecond)
			}
		}
		t.Fatalf("failed to escape after %d attempts", maxAttempts)
	}
	fleeUntilEscaped(40)

	enter() // round result -> fled screen
	enter() // fled screen -> pirate bay

	waitFor("drops anchor")
	enter() // bay intro -> outcome

	waitFor("textbook slipped through")
	enter() // outcome -> next-plan choice

	waitFor("What be your next plan")
	enter() // select "Explore the Pirate Bay" (index 0) -> ending text

	waitFor("curiosity and adventure")
	enter() // ending text -> close

	waitFor("This be the outcomes of yer adventures as TestCaptain")
	enter() // close -> quit

	tm.WaitFinished(t, teatest.WithFinalTimeout(2*time.Second))
}

// TestFightDeathPath verifies the flee-or-die promise's other branch,
// deterministically: rigged stats (dodge far below the enemy's accuracy,
// zero armor, 1 HP) guarantee the enemy's first hit is lethal regardless of
// initiative order or which of the six attacks gets rolled, so this doesn't
// depend on getting lucky/unlucky RNG like the happy-path flee loop above.
func TestFightDeathPath(t *testing.T) {
	state := NewGameState()
	state.PlayerName = "Doomed"
	state.ShipName = "Wreck"
	state.Player = NewShip(Sloop)
	state.Player.Health = 1
	state.Player.Armor = 0
	state.Player.Dodge = -100
	state.Player.BroadsideType = 2
	state.Enemy.Accuracy = 100

	tm := teatest.NewTestModel(t, newFightModel(state), teatest.WithInitialTermSize(100, 40))

	var seen bytes.Buffer
	waitFor := func(text string) {
		t.Helper()
		deadline := time.Now().Add(2 * time.Second)
		for {
			b, _ := io.ReadAll(tm.Output())
			seen.Write(b)
			if strings.Contains(normalizeOutput(seen.Bytes()), text) {
				return
			}
			if time.Now().After(deadline) {
				t.Fatalf("waitFor %q: condition not met after 2s. Seen so far:\n%s", text, seen.String())
			}
			time.Sleep(20 * time.Millisecond)
		}
	}
	enter := func() { tm.Send(tea.KeyMsg{Type: tea.KeyEnter}) }

	waitFor("ominous privateer")
	enter() // fight intro -> choosing

	waitFor("What be yer move")
	enter() // choose Broadside (index 0); the enemy's guaranteed hit kills regardless

	waitFor("It strikes true")
	enter() // round result -> dead screen

	waitFor("You are dead")
	enter() // -> quit

	tm.WaitFinished(t, teatest.WithFinalTimeout(2*time.Second))
}
