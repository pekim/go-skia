package generate

import (
	"errors"
	"fmt"
	"os/exec"
	"path"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type translationUnit struct {
	clang.TranslationUnit
}

func newTranslationUnit(headerFile string) translationUnit {
	resourcesDir, err := clangResourceDir()
	if err != nil {
		panic(err)
	}

	parseArgs := []string{
		"-I", path.Join(resourcesDir, "include"),
		"-I", "./skia/skia/",
		"-x", "c++-header",
	}

	var transUnit translationUnit
	index := clang.NewIndex(0, 1)
	errCode := index.ParseTranslationUnit2(headerFile, parseArgs, nil,
		clang.TranslationUnit_SkipFunctionBodies, &transUnit.TranslationUnit)
	if errCode != clang.Error_Success {
		panic(errCode)
	}
	return transUnit
}

func clangResourceDir() (string, error) {
	out, err := exec.Command("clang", "-print-resource-dir").Output()
	if err != nil {
		panic(err)
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
