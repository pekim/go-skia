package generate

import (
	"fmt"
	"go/format"
	"os"
	"strings"
)

func (g *generator) generate() {
	g.goFile = newGenFile("./skia.go")
	defer g.goFile.close()

	g.goFile.writeln("package skia")
	g.goFile.writeln("")

	for _, class := range g.classes {
		class.generate(g.goFile)
	}

	for _, enum := range g.enums {
		enum.generate(g.goFile)
	}
}

type genFile struct {
	filename string
	content  strings.Builder
}

func newGenFile(filename string) *genFile {
	f := &genFile{
		filename: filename,
	}
	f.writeln("// This is a generated file. DO NOT EDIT.")
	f.writeln("")

	return f
}

func (f *genFile) close() {
	unformatted := []byte(f.content.String())
	formatted, err := format.Source(unformatted)
	if err != nil {
		fmt.Printf("failed to format %s: %s\n", f.filename, err.Error())
		formatted = unformatted
	}
	err = os.WriteFile(f.filename, formatted, 0644)
	if err != nil {
		panic(err)
	}
}

func (f *genFile) write(content string) {
	f.content.WriteString(content)
}

func (f *genFile) writeln(line string) {
	f.write(line)
	f.write("\n")
}

func (f *genFile) writef(format string, args ...any) {
	f.write(fmt.Sprintf(format, args...))
}

func (f *genFile) writelnf(format string, args ...any) {
	f.writef(format, args...)
	f.write("\n")
}
