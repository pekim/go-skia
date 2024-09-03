package generate

import "strings"

func stripSkPrefix(name string) string {
	return strings.TrimPrefix(name, "Sk")
}

func stripKPrefix(name string) string {
	return strings.TrimPrefix(name, "k")
}

func validGoName(name string) string {
	if name == "type" {
		return "typ"
	}
	return name
}
