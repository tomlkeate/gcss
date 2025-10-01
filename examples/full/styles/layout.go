package styles

import (
	"github.com/tomlkeate/gcss"
	"github.com/tomlkeate/gcss/props"
	"github.com/tomlkeate/gcss/variables"
)

// Layout returns the styles for the layout for the base stylesheet.
func (ss *StyleSheet) Layout() Styles {
	return Styles{
		{
			Selector: "body",
			Props: gcss.Props{
				MinHeight: variables.FullScreenHeight,
			},
		},
		{
			Selector: "main",
			Props: gcss.Props{
				Display: props.DisplayGrid,
			},
		},
	}
}

// Layout returns the styles for the layout for the media.
func (m *Media) Layout() Styles {
	return Styles{
		{
			Selector: "main",
			Props: gcss.Props{
				Padding: m.Padding,
				RowGap:  m.VerticalGap,
			},
		},
	}
}

// Layout returns the styles for the layout for the theme.
func (t *Theme) Layout() Styles {
	return Styles{
		{
			Selector: "body",
			Props: gcss.Props{
				BackgroundColor: t.Background,
				Color:           t.Foreground,
			},
		},
	}
}
