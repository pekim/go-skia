package generate

import (
	"fmt"

	"github.com/go-clang/clang-v15/clang"
)

type generator struct {
	transUnit  clang.TranslationUnit
	classes    []class
	enums      []enum
	goFile     *fileGo
	headerFile *fileHeader
	cppFile    *fileCpp
}

func Generate() {
	// Put some space after ninja output (from build.sh).
	fmt.Println("")

	g := generator{}
	fmt.Println("parse")
	g.parse()
	fmt.Println("visit")
	g.visit()
	fmt.Println("generate")
	g.generate()
}
