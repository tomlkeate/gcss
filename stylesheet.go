package gcss

import (
	"context"
	"fmt"
	"io"
)

type CSSComponent interface {
	CSS(w io.Writer) error
}

// Stylesheet is a collection of styles.
type Stylesheet []CSSComponent

// Render writes the CSS representation of the stylesheet to the writer.
// this is to ensure that it implements a templ.Component.
func (ss Stylesheet) Render(ctx context.Context, w io.Writer) error {
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

type MediaBlock struct {
	MediaQuery string
	Styles     Stylesheet
}

func (mb *MediaBlock) CSS(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s{", mb.MediaQuery); err != nil {
		return err
	}
	for _, style := range mb.Styles {
		if err := style.CSS(w); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprint(w, "}"); err != nil {
		return err
	}
	return nil
}
