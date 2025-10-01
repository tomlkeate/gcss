package styles

import (
	"github.com/tomlkeate/gcss"
	"github.com/tomlkeate/gcss/props"
	"github.com/tomlkeate/gcss/variables"
)

// Resets returns the styles for resets for the base stylesheet.
func (ss *StyleSheet) Resets() Styles {
	return Styles{
		{
			Selector: "*,::after,::before,::backdrop,::file-selector-button",
			Props: gcss.Props{
				Border: props.Border{
					Width: variables.Size0,
					Style: props.BorderStyleSolid,
				},
				BoxSizing: props.BoxSizingBorderBox,
				Margin:    variables.Size0,
				Padding:   variables.Size0,
			},
		},
		{
			Selector: "html,:host",
			Props: gcss.Props{
				FontFamily: props.FontFamilySans,
				LineHeight: props.UnitRaw(1.5),
			},
			CustomProps: []gcss.CustomProp{
				{Attr: "-webkit-text-size-adjust", Value: "100%"},
				{Attr: "tab-size", Value: "4"},
			},
		},
		{
			Selector: "body",
			Props: gcss.Props{
				LineHeight: props.UnitInherit(),
			},
		},
		// more resets
	}
}
