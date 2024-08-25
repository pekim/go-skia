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
	formatted, err := format.Source([]byte(f.content.String()))
	if err != nil {
		panic(err)
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
