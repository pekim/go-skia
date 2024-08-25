package generate

import (
	"fmt"
	"os"
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
	file     *os.File
}

func newGenFile(filename string) genFile {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	return genFile{
		filename: filename,
		file:     file,
	}
}

func (f genFile) close() {
	err := f.file.Close()
	if err != nil {
		panic(err)
	}
}

func (f genFile) write(text string) {
	_, err := f.file.WriteString(text)
	if err != nil {
		panic(err)
	}
}

func (f genFile) writeln(line string) {
	f.write(line)
	f.write("\n")
}

func (f genFile) writef(format string, args ...any) {
	f.write(fmt.Sprintf(format, args...))
}

func (f genFile) writelnf(format string, args ...any) {
	f.writef(format, args...)
	f.write("\n")
}
