package generate

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Generate() {
	// Put some space after C deprecation warning
	fmt.Println("")
	fmt.Println("")

	g := generator{}
	g.parse()
	g.generate()
}

func clangResourceDir() (string, error) {
	out, err := exec.Command("clang", "-print-resource-dir").Output()
	if err != nil {
		log.Fatal(err)
	}

	resDir := strings.TrimSpace(string(out))
	parts := strings.Split(resDir, "\n")
	resDir = parts[0]

	if resDir == "" {
		return "", errors.New("no output when getting clang resource dir")
	}
	if !strings.HasPrefix(resDir, "/") {
		return "", fmt.Errorf("expected clang resource dir to start with '/', but it %s", resDir)
	}

	return resDir, nil
}
