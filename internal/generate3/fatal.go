package generate

import "fmt"

func fatal(v any) {
	panic(v)
}

func fatalOnError(err error) {
	if err != nil {
		fatal(err)
	}
}

func fatalf(format string, args ...any) {
	fatal(fmt.Sprintf(format, args...))
}
