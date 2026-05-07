package main

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	termimg "github.com/blacktop/go-termimg"
)

// renderResultMsg 承载异步渲染结果
type renderResultMsg struct {
	rendered string
	err      error
}

// protocolCount 协议总数，避免硬编码魔法数字
const protocolCount termimg.Protocol = 5

type model struct {
	widget        *termimg.ImageWidget
	protocol      termimg.Protocol
	renderedImage string
}

// renderCmd 将渲染逻辑封装为 Cmd，避免在 View() 中产生副作用
func renderCmd(w *termimg.ImageWidget, p termimg.Protocol) tea.Cmd {
	return func() tea.Msg {
		w.SetProtocol(p)
		rendered, err := w.Render()
		return renderResultMsg{rendered: rendered, err: err}
	}
}

func (m model) Init() tea.Cmd {
	// 程序启动时触发首次渲染
	return renderCmd(m.widget, m.protocol)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case renderResultMsg:
		// 接收渲染结果，更新缓存，View() 直接读取，无副作用
		m.renderedImage = msg.rendered
		return m, nil

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "space":
			// 切换协议，基于 protocolCount 常量，不再硬编码
			m.protocol = (m.protocol % protocolCount) + 1
			// 协议变更后触发重新渲染，同时清屏
			return m, renderCmd(m.widget, m.protocol)

		case "ctrl+z":
			return m, tea.Suspend
		}
	}

	return m, nil
}

// View() 是纯函数：只读取 model 状态，不产生任何副作用
func (m model) View() tea.View {
	var content string

	line := "Current protocol: " + m.protocol.String()
	content = strings.Join([]string{line, m.renderedImage}, "\n")

	v := tea.NewView(content)
	v.AltScreen = true
	return v
}
