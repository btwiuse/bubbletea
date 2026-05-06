package config

type Config struct {
	Style   Style
	Workers int
}

type Style struct {
	ListItem ListItem
	Visited  AdaptiveColor
	Window   Window
	Tab      Tab
}

type ListItem struct {
	NormalTitle   AdaptiveColor
	NormalDesc    AdaptiveColor
	SelectedTitle ListItemStyle
	SelectedDesc  AdaptiveColor
	DimmedTitle   AdaptiveColor
	DimmedDesc    AdaptiveColor
	FilterMatch   ListItemStyle
}

type ListItemStyle struct {
	BorderForeground AdaptiveColor
	Foreground       AdaptiveColor
}

type Window struct {
	Border string
	Color  AdaptiveColor
}

type Tab struct {
	Color AdaptiveColor
}

type AdaptiveColor struct {
	Dark  string
	Light string
}

// LoadConfig returns the default configuration.
func LoadConfig() (*Config, error) {
	return basicConfig(), nil
}

// basicConfig returns the default configuration.
func basicConfig() *Config {
	return &Config{
		Style: Style{
			ListItem: ListItem{
				NormalTitle: AdaptiveColor{
					Dark:  "#E6EBE9",
					Light: "#7D56F4",
				},
				NormalDesc: AdaptiveColor{
					Dark:  "#4f4f4f",
					Light: "#7D56F4",
				},
				SelectedTitle: ListItemStyle{
					BorderForeground: AdaptiveColor{
						Dark:  "#2A8f69",
						Light: "#7D56F4",
					},
					Foreground: AdaptiveColor{
						Dark:  "#2A8f69",
						Light: "#7D56F4",
					},
				},
				SelectedDesc: AdaptiveColor{
					Dark:  "#4f4f4f",
					Light: "#7D56F4",
				},
				DimmedTitle: AdaptiveColor{
					Dark:  "#874BFD",
					Light: "#7D56F4",
				},
				DimmedDesc: AdaptiveColor{
					Dark:  "#874BFD",
					Light: "#7D56F4",
				},
				FilterMatch: ListItemStyle{
					BorderForeground: AdaptiveColor{
						Dark:  "#2A8f69",
						Light: "#7D56F4",
					},
					Foreground: AdaptiveColor{
						Dark:  "#2A8f69",
						Light: "#7D56F4",
					},
				},
			},
			Visited: AdaptiveColor{
				Dark:  "#777777",
				Light: "#777777",
			},
			Window: Window{
				Border: "normal",
				Color: AdaptiveColor{
					Dark:  "#2A8f69",
					Light: "#7D56F4",
				},
			},
			Tab: Tab{
				Color: AdaptiveColor{
					Dark:  "#2A8f69",
					Light: "#7D56F4",
				},
			},
		},
		Workers: 10,
	}
}
