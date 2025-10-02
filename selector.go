package gcss

import "strings"

type Selector string

// Multiple selectors comma seperated
func SelectorGroup(selectors ...Selector) Selector {
	strs := make([]string, len(selectors))

	for i, s := range selectors {
		strs[i] = string(s)
	}

	return Selector(strings.Join(strs, ", "))
}

// Element selector
func SelectorElement(element string) Selector {
	return Selector(element)
}

// Id selector
func SelectorId(id string) Selector {
	return Selector("#" + id)
}

// Class selector
func SelectorClass(class string) Selector {
	return Selector("." + class)
}

// Custom selector, such as psuedo classes
func SelectorCustom(s string) Selector {
	return Selector(s)
}

func (s Selector) String() string {
	return string(s)
}
