package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const bannerArt = `
                    |    |    |
                   )_)  )_)  )_)
                  )___))___))___)\
                 )____)____)_____)\\
               _____|____|____|____\\\__
      ---------\                   /---------
        ^^^^^ ^^^^^^^^^^^^^^^^^^^^^
           ^^^^      ^^^^     ^^^
`

var (
	bannerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("37")).Bold(true)

	panelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("37")).
			Padding(1, 2).
			Width(74)
)

// panel wraps a screen's content in a consistent bordered box. Every screen's
// View should return panel(...) as its final step.
func panel(content string) string {
	return panelStyle.Render(content)
}

func banner() string {
	return bannerStyle.Render(bannerArt)
}

// healthBar renders a labeled, colored block-character bar - green when
// healthy, amber in the middle, red when critical - used for the player's
// health/bandwidth and the enemy's health during the fight screen.
func healthBar(label string, current, max, width int) string {
	if max <= 0 {
		max = 1
	}
	clamped := current
	if clamped < 0 {
		clamped = 0
	}
	if clamped > max {
		clamped = max
	}

	filled := int(float64(width) * float64(clamped) / float64(max))
	if filled > width {
		filled = width
	}

	var barColor lipgloss.Color
	switch ratio := float64(clamped) / float64(max); {
	case ratio > 0.6:
		barColor = lipgloss.Color("42")
	case ratio > 0.3:
		barColor = lipgloss.Color("214")
	default:
		barColor = lipgloss.Color("196")
	}

	bar := strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
	styledBar := lipgloss.NewStyle().Foreground(barColor).Render(bar)

	return fmt.Sprintf("%-10s[%s] %d/%d", label, styledBar, current, max)
}
