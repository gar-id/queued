package tools

func DefaultString(a, b string) string {
	var empty string
	if a != empty {
		return a
	}
	return b
}

func DefaultValue(a, b interface{}) interface{} {
	var empty string
	if a != empty {
		return a
	}
	return b
}

func DefaultInt(a, b int) int {
	var empty int
	if a != empty {
		return a
	}
	return b
}
