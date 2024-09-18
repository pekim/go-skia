package generate

import (
	"github.com/go-clang/clang-v15/clang"
)

type method struct {
	callable
}

func (m *method) enrich1(record *record, cursor clang.Cursor) {
	m.callable.enrich1(record, cursor)
}
