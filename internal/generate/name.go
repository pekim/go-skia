package generate

import (
	"strings"
)

func stripSkPrefix(name string) string {
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
