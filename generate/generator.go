package generate

import (
	"fmt"

	"github.com/go-clang/clang-v15/clang"
)

type generator struct {
	transUnit clang.TranslationUnit
}

func Generate() {
	// Put some space after C deprecation warning from clang package.
	fmt.Println("")
	fmt.Println("")

	g := generator{}
	g.parse()
	g.visit()
	g.generate()
}
