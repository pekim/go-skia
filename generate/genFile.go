package generate

import (
	"fmt"
	"go/format"
	"os"
	"os/exec"
	"strings"
)

type genFile struct {
	filename string
	content  strings.Builder
	goFmt    bool
}

func newGenFile(filename string) *genFile {
	f := &genFile{
		filename: filename,
	}
	f.generatedFileComment()

	return f
}

func (f *genFile) generatedFileComment() {
	f.writeln("// This is a generated file. DO NOT EDIT.")
	f.writeln("")
}

func (f *genFile) close() {
	unformatted := []byte(f.content.String())
	var formatted []byte

	if f.goFmt {
		var err error
		formatted, err = format.Source(unformatted)
		if err != nil {
			fmt.Printf("failed to format %s: %s\n", f.filename, err.Error())
			formatted = unformatted
		}
	} else {
		formatted = unformatted
	}

	err := os.WriteFile(f.filename, formatted, 0644)
	if err != nil {
		panic(err)
	}
}

func (f *genFile) clangFormat() {
	cmd := exec.Command("clang-format", "-style=llvm", "-i", f.filename)
	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		fmt.Printf("Output of clang-format for %s\n", f.filename)
		fmt.Println(string(output))
	}
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

func (f *genFile) writelnfTrim(format string, args ...any) {
	f.write(strings.TrimSpace(fmt.Sprintf(format, args...)))
	f.write("\n")
}
