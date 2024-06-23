package logger

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	ErrorStyle = lipgloss.NewStyle().Foreground(
		lipgloss.AdaptiveColor{
			Light: "#ff6666",
			Dark:  "#ff2222",
		},
	)
	WarnStyle = lipgloss.NewStyle().Foreground(
		lipgloss.AdaptiveColor{
			Light: "#ffcc66",
			Dark:  "#ffaa22",
		},
	)
	InfoStyle = lipgloss.NewStyle().Foreground(
		lipgloss.AdaptiveColor{
			Light: "#6666ff",
			Dark:  "#2222ff",
		},
	)
	SuccessStyle = lipgloss.NewStyle().Foreground(
		lipgloss.AdaptiveColor{
			Light: "#66ff66",
			Dark:  "#22ff22",
		},
	)
	ErrorMark = "✘"
	WarnMark  = "⚠"
	InfoMark  = "(i)"
	SuccessMark = "✔"
)