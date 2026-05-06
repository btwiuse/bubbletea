package theme

import "image/color"

// Palette holds all the colors for a theme
type Palette struct {
	Name       string
	Background color.Color
	Foreground color.Color // untyped text, hints
	Typed      color.Color // correctly typed
	Error      color.Color // mistakes
	Cursor     color.Color // current character
	Accent     color.Color // highlights, timer, active elements
	Success    color.Color // personal best, positive feedback
}

var All = []Palette{TokyoNight, Gruvbox, Sakura, Monkeytype, Monochrome, Forest, Espresso, Lumon, Mars, Void, Everforest, Chameleon}

var Current = TokyoNight

// Next cycles themes
func Next() {
	for i, t := range All {
		if t.Name == Current.Name {
			Current = All[(i+1)%len(All)]
			return
		}
	}
}

// ByName finds a theme, defaults to TokyoNight
func ByName(name string) Palette {
	for _, t := range All {
		if t.Name == name {
			return t
		}
	}
	return TokyoNight
}
