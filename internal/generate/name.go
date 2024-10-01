package generate

import (
	"strings"
	"unicode"
)

func stripSkPrefix(name string) string {
	if len(name) > 2 {
		rune2 := []rune(name)[2]
		if !unicode.IsUpper([]rune(name)[2]) && rune2 != '_' {
			return name
		}
	}

	name = strings.TrimPrefix(name, "Sk")
	name = strings.TrimPrefix(name, "SK_")
	return name
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

func goExportedName(name string) string {
	name = strings.ReplaceAll(name, "::", "")
	name = strings.ToUpper(name[:1]) + name[1:]
	return name
}
