package style

import (
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"

	"examples/hackertea/internal/constants"
)

func DocStyle() lipgloss.Style {

	return lipgloss.NewStyle().
		Padding(1, 2, 1, 2)
}

func WindowStyle(light, dark, border string) lipgloss.Style {
	highlightColor := compat.AdaptiveColor{
		Light: lipgloss.Color(light),
		Dark:  lipgloss.Color(dark),
	}

	borderStyle := lipgloss.NormalBorder()

	switch border {
	case constants.RoundedBorder:
		borderStyle = lipgloss.RoundedBorder()
	case constants.ThickBorder:
		borderStyle = lipgloss.ThickBorder()
	case constants.DoubleBorder:
		borderStyle = lipgloss.DoubleBorder()
	}

	return lipgloss.NewStyle().
		BorderForeground(highlightColor).
		Padding(2, 2).
		Border(borderStyle).UnsetBorderTop()
}

func TabGapStyle(light, dark, border string) lipgloss.Style {
	highlightColor := compat.AdaptiveColor{
		Light: lipgloss.Color(light),
		Dark:  lipgloss.Color(dark),
	}
	tabGapBorder := tabGapBorderWithBottom("┴", border)

	return lipgloss.NewStyle().Border(tabGapBorder, true).
		BorderForeground(highlightColor).
		Padding(0, 0).
		BorderTop(false).
		BorderLeft(false)
}

func TitleTabStyle(light, dark, border string) lipgloss.Style {
	activeTabBorder := tabBorderWithBottom("├", "─", "┴", border)

	highlightColor := compat.AdaptiveColor{
		Light: lipgloss.Color(light),
		Dark:  lipgloss.Color(dark),
	}

	return lipgloss.NewStyle().
		Border(activeTabBorder, true).
		BorderForeground(highlightColor).
		Foreground(highlightColor).
		Padding(0, 1)
}

func ActiveTabStyle(light, dark, border string) lipgloss.Style {
	activeTabBorder := tabBorderWithBottom("┘", " ", "└", border)
	highlightColor := compat.AdaptiveColor{
		Light: lipgloss.Color(light),
		Dark:  lipgloss.Color(dark),
	}

	return lipgloss.NewStyle().
		Border(activeTabBorder, true).
		BorderForeground(highlightColor).
		Padding(0, 1)
}

func InActiveTabStyle(light, dark, border string) lipgloss.Style {
	inActiveTabBorder := tabBorderWithBottom("┴", "─", "┴", border)
	highlightColor := compat.AdaptiveColor{
		Light: lipgloss.Color(light),
		Dark:  lipgloss.Color(dark),
	}

	return lipgloss.NewStyle().
		Border(inActiveTabBorder, true).
		BorderForeground(highlightColor).
		Padding(0, 1)
}

func ItemNormalTitleStyle(light, dark string) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(compat.AdaptiveColor{Light: lipgloss.Color(light), Dark: lipgloss.Color(dark)}).
		Padding(0, 0, 0, 2)
}

func ItemNormalDescStyle(light, dark string) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(compat.AdaptiveColor{Light: lipgloss.Color(light), Dark: lipgloss.Color(dark)}).
		Padding(0, 0, 0, 2)
}

func ItemSelectedTitleStyle(bLight, bDark, light, dark string) lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(compat.AdaptiveColor{Light: lipgloss.Color(bLight), Dark: lipgloss.Color(bDark)}).
		Foreground(compat.AdaptiveColor{Light: lipgloss.Color(light), Dark: lipgloss.Color(dark)}).
		Padding(0, 0, 0, 1)
}

func ItemSelectedDescStyle(bLight, bDark, light, dark string) lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(compat.AdaptiveColor{Light: lipgloss.Color(bLight), Dark: lipgloss.Color(bDark)}).
		Foreground(compat.AdaptiveColor{Light: lipgloss.Color(light), Dark: lipgloss.Color(dark)}).
		Padding(0, 0, 0, 1)
}

func ItemDimmedTitleStyle(light, dark string) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(compat.AdaptiveColor{Light: lipgloss.Color(light), Dark: lipgloss.Color(dark)}).
		Padding(0, 0, 0, 2)
}

func ItemDimmedDescStyle(light, dark string) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(compat.AdaptiveColor{Light: lipgloss.Color(light), Dark: lipgloss.Color(dark)}).
		Padding(0, 0, 0, 2)
}

func VisitedStyle(light, dark string) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(compat.AdaptiveColor{Light: lipgloss.Color(light), Dark: lipgloss.Color(dark)})
}

func FilterMatchedStyle(bLight, bDark, light, dark string) lipgloss.Style {
	return lipgloss.NewStyle().
		Background(compat.AdaptiveColor{Light: lipgloss.Color(bLight), Dark: lipgloss.Color(bDark)}).
		Foreground(compat.AdaptiveColor{Light: lipgloss.Color(light), Dark: lipgloss.Color(dark)}).
		Underline(true)
}

func tabBorderWithBottom(left, middle, right, border string) lipgloss.Border {
	borderStyle := lipgloss.RoundedBorder()

	switch border {
	case constants.RoundedBorder:
		borderStyle = lipgloss.RoundedBorder()
	case constants.ThickBorder:
		borderStyle = lipgloss.ThickBorder()
	case constants.DoubleBorder:
		borderStyle = lipgloss.DoubleBorder()
	}

	borderStyle.BottomLeft = left
	borderStyle.Bottom = middle
	borderStyle.BottomRight = right

	return borderStyle
}

func tabGapBorderWithBottom(left, border string) lipgloss.Border {
	borderStyle := lipgloss.RoundedBorder()

	switch border {
	case constants.RoundedBorder:
		borderStyle = lipgloss.RoundedBorder()
	case constants.ThickBorder:
		borderStyle = lipgloss.ThickBorder()
	case constants.DoubleBorder:
		borderStyle = lipgloss.DoubleBorder()
	}

	borderStyle.BottomLeft = left
	borderStyle.BottomRight = borderStyle.TopRight
	borderStyle.Right = ""

	return borderStyle
}
