package props

type Content string

// Content custom value
func ContentCustom(c string) Content {
	return Content(c)
}

func (c Content) String() string {
	return string(c)
}
