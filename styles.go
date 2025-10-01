package gcss

import (
	"context"
	"io"
)

type CSSComponent interface {
	CSS(w io.Writer) error
}

// Styles is a collection of styles.
type Styles []CSSComponent

// Render writes the CSS representation of the stylesheet to the writer.
// this is to ensure that it implements a templ.Component.
func (ss Styles) Render(ctx context.Context, w io.Writer) error {
	//TODO: Get a CSP nonce from the context.
	if _, err := io.WriteString(w, "<style type=\"text/css\">\n"); err != nil {
		return err
	}
	for _, s := range ss {
		if err := s.CSS(w); err != nil {
			return err
		}
	}
	if _, err := io.WriteString(w, "\n</style>"); err != nil {
		return err
	}
	return nil
}

func (ss Styles) CSS(w io.Writer) error {
	for _, s := range ss {
		if err := s.CSS(w); err != nil {
			return err
		}
	}
	return nil
}
