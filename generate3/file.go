package generate

import (
	"fmt"
	"go/format"
	"os"
	"os/exec"
	"strings"
)

type file struct {
	filename string
	content  strings.Builder
	goFmt    bool
}

func newFile(filename string) *file {
	f := &file{
		filename: filename,
	}
	f.generatedFileComment()

	return f
}

func (f *file) generatedFileComment() {
	f.writeln("// This is a generated file. DO NOT EDIT.")
	f.writeln()
}

func (f *file) close() {
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

func (f *file) clangFormat() {
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

func (f *file) write(content string) {
	f.content.WriteString(content)
}

func (f *file) writeln(args ...any) {
	line := fmt.Sprint(args...)
	f.write(line)
	f.write("\n")
}

func (f *file) writef(format string, args ...any) {
	f.write(fmt.Sprintf(format, args...))
}

func (f *file) writelnf(format string, args ...any) {
	f.writef(format, args...)
	f.write("\n")
}
