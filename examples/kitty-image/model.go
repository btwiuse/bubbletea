package main

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	termimg "github.com/blacktop/go-termimg"
)

type model struct {
	widget   *termimg.ImageWidget
	protocol termimg.Protocol
}

func (m model) View() tea.View {
	line := "Current protocol: " + m.protocol.String()

	m.widget.SetProtocol(m.protocol)
	rendered, _ := m.widget.Render()

	v := tea.NewView(strings.Join([]string{line, rendered}, "\n"))
	v.AltScreen = true

	return v
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "space":
			m.protocol = (m.protocol % 5) + 1
			return m, tea.ClearScreen
		case "ctrl+z":
			return m, tea.Suspend
		}
	}
	return m, nil
}
