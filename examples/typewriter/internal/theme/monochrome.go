package theme

import "charm.land/lipgloss/v2"

var Monochrome = Palette{
	Name:       "monochrome",
	Background: lipgloss.Color("#0a0a0a"),
	Foreground: lipgloss.Color("#505050"),
	Typed:      lipgloss.Color("#d0d0d0"),
	Error:      lipgloss.Color("#ff6b6b"),
	Cursor:     lipgloss.Color("#ffffff"),
	Accent:     lipgloss.Color("#ffffff"),
	Success:    lipgloss.Color("#a0a0a0"),
}
