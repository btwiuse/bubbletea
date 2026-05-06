package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
	boba "github.com/btwiuse/boba"
	"charm.land/lipgloss/v2"
)

// LayerHitMsg identifies which layer was hit by a mouse event.
type LayerHitMsg struct {
	ID    string
	Mouse tea.MouseMsg
}

const tickInterval = 160 * time.Millisecond // ms

var logoColors = []string{
	"#FF6B6B", // coral
	"#4ECDC4", // teal
	"#45B7D1", // sky blue
	"#96CEB4", // sage
	"#FFEAA7", // yellow
	"#DDA0DD", // plum
	"#FF9FF3", // pink
	"#54A0FF", // cornflower
}

const logoArt = `████   ██     ██  █████
██   ██  ██   ██   ██   ██
██   ██   █   █    ██   ██
██   ██    █ █     ██   ██
█████       █      █████ `

type model struct {
	x, y          int
	vx, vy        int
	width, height int
	dragging      bool
	dragOffX      int
	dragOffY      int
	paused        bool
	colorIndex    int
}

func (m model) logoText() string {
	color := lipgloss.Color(logoColors[m.colorIndex%len(logoColors)])
	return lipgloss.NewStyle().
		Width(36).
		Height(8).
		Padding(1, 3).
		Align(lipgloss.Center).
		AlignVertical(lipgloss.Center).
		Bold(true).
		Foreground(color).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(color).
		Render(logoArt)
}

func (m model) logoWidth() int  { return lipgloss.Width(m.logoText()) }
func (m model) logoHeight() int { return lipgloss.Height(m.logoText()) }

func (m model) Init() tea.Cmd {
	return tea.Tick(tickInterval, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}

type tickMsg struct{}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		if m.x == 0 && m.y == 0 {
			m.x = (m.width - m.logoWidth()) / 2
			m.y = (m.height - m.logoHeight()) / 2
		}

	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}

	case tickMsg:
		if !m.paused {
			lw, lh := m.logoWidth(), m.logoHeight()
			m.x += m.vx
			m.y += m.vy

			if m.x <= 0 || m.x+lw >= m.width {
				m.vx = -m.vx
				m.x = clamp(m.x, 0, m.width-lw)
				m.colorIndex++
			}
			if m.y <= 0 || m.y+lh >= m.height {
				m.vy = -m.vy
				m.y = clamp(m.y, 0, m.height-lh)
				m.colorIndex++
			}
		}
		return m, tea.Tick(tickInterval, func(t time.Time) tea.Msg {
			return tickMsg{}
		})

	case LayerHitMsg:
		mouse := msg.Mouse.Mouse()

		switch msg.Mouse.(type) {
		case tea.MouseClickMsg:
			if mouse.Button == tea.MouseLeft && msg.ID == "logo" {
				m.dragging = true
				m.paused = true
				m.dragOffX = mouse.X - m.x
				m.dragOffY = mouse.Y - m.y
			}

		case tea.MouseMotionMsg:
			if m.dragging {
				lw, lh := m.logoWidth(), m.logoHeight()
				m.x = clamp(mouse.X-m.dragOffX, 0, m.width-lw)
				m.y = clamp(mouse.Y-m.dragOffY, 0, m.height-lh)
			}

		case tea.MouseReleaseMsg:
			m.dragging = false
			m.paused = false
		}
	}

	return m, nil
}

func (m model) View() tea.View {
	var v tea.View

	bg := lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Top,
		lipgloss.Left,
		"",
		lipgloss.WithWhitespaceChars("."),
		lipgloss.WithWhitespaceStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("238"))),
	)

	root := lipgloss.NewLayer(bg).ID("bg")
	root.AddLayers(lipgloss.NewLayer(m.logoText()).ID("logo").X(m.x).Y(m.y))

	comp := lipgloss.NewCompositor(root)

	v.MouseMode = tea.MouseModeAllMotion
	v.AltScreen = true
	v.OnMouse = func(msg tea.MouseMsg) tea.Cmd {
		return func() tea.Msg {
			mouse := msg.Mouse()
			if id := comp.Hit(mouse.X, mouse.Y).ID(); id != "" {
				return LayerHitMsg{ID: id, Mouse: msg}
			}
			return nil
		}
	}
	v.SetContent(comp.Render())

	return v
}

func main() {
	dirs := []int{-1, 1}
	m := model{
		vx: dirs[rand.Intn(2)],
		vy: dirs[rand.Intn(2)],
	}

	if _, err := boba.NewProgram(m).Run(); err != nil {
		fmt.Println("Error while running program:", err)
		os.Exit(1)
	}
}

func clamp(n, min, max int) int {
	if n < min {
		return min
	}
	if n > max {
		return max
	}
	return n
}
