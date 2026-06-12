package greet

import "strings"

func Hello(name string) string { // exported functions start with capital letters

	clean := normalizeName(name)
	return "Hello," + clean
	
}

func normalizeName (name string) string {

	n := strings.TrimSpace(name)

	if n == "" {
		return "Guest"
	} else {
		return strings.ToUpper(n)
	}
}