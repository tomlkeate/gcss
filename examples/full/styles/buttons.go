package styles

import (
	"github.com/tomlkeate/gcss"
	"github.com/tomlkeate/gcss/props"
	"github.com/tomlkeate/gcss/variables"
)

// Buttons returns the styles for buttons for the base stylesheet.
func (ss *StyleSheet) Buttons() Styles {
	return Styles{
		{
			Selector: ".button",
			Props: gcss.Props{
				AlignItems:     props.AlignItemsCenter,
				BorderRadius:   variables.Size1H,
				Display:        props.DisplayInlineFlex,
				FontSize:       variables.Size3H,
				FontWeight:     props.FontWeightMedium,
				Height:         variables.Size10,
				JustifyContent: props.JustifyContentCenter,
				LineHeight:     variables.Size5,
				PaddingTop:     variables.Size2,
				PaddingRight:   variables.Size4,
				PaddingBottom:  variables.Size2,
				PaddingLeft:    variables.Size4,
			},
		},
	}
}

// Buttons returns the styles for buttons for the theme.
func (t *Theme) Buttons() Styles {
	return Styles{
		{
			Selector: ".button-primary",
			Props: gcss.Props{
				BackgroundColor: t.Primary,
				Color:           t.PrimaryForeground,
			},
		},
	}
}
