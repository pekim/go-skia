package generate

import "github.com/go-clang/clang-v15/clang"

type param struct {
	cName   string
	cgoName string
	goName  string
}

func newParam(cursor clang.Cursor) param {
	cName := cursor.DisplayName()
	return param{
		cName:   cName,
		cgoName: "c_" + cName,
		goName:  validGoName(cName),
	}
}
