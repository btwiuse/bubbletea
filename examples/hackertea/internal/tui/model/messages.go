package model

import (
	"charm.land/bubbles/v2/list"
)

type initMsg struct{}

type next struct {
	items []list.Item
}
