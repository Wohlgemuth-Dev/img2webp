package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Lipgloss Styles
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 2)

	InfoStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575")).Bold(true)
	SuccessStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#A8CC8C")).Bold(true)
	ErrorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#E88388")).Bold(true)
	WarnStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#DDB6F2"))
	Highlight    = lipgloss.NewStyle().Foreground(lipgloss.Color("#E2E1ED")).Bold(true)
	DimStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262"))
)
