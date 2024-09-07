package generate

import (
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type method struct {
	CName      string `json:"name"`
	goFuncName string
	cFuncName  string
	record     *record
	doc        string
	isStatic   bool
	params     []param
	retrn      typ
	enriched   bool
}

func (m *method) enrich(api api, record *record, cursor clang.Cursor) {
	m.record = record
	m.goFuncName = fmt.Sprintf("%s%s", m.record.goName, m.CName)
	m.cFuncName = fmt.Sprintf("misk_%s_%s", m.record.goName, m.CName)
	m.doc = cursor.RawCommentText()
	m.isStatic = cursor.CXXMethod_IsStatic()

	if !m.isStatic {
		panic("TODO non-static methods")
	}

	paramCount := int(cursor.NumArguments())
	m.params = make([]param, paramCount)
	for i := 0; i < paramCount; i++ {
		arg := cursor.Argument(uint32(i))
		param := newParam(i, arg, api)
		m.params[i] = param
	}

	m.retrn = mustTypFromClangType(cursor.ResultType(), api)

	m.enriched = true
}

func (m method) generate(g generator) {
	m.generateGo(g)
	m.generateHeader(g)
	m.generateCpp(g)
}

func (m method) generateGo(g generator) {
	if !m.enriched {
		fatalf("method %s has not been enriched", m.CName)
	}

	f := g.goFile

	params := make([]string, len(m.params))
	cVars := make([]string, len(m.params))
	cArgs := make([]string, len(m.params))
	for i, param := range m.params {
		params[i] = fmt.Sprintf("%s %s", param.goName, param.typ.goName)
		cVars[i] = param.cgoVar
		cArgs[i] = param.cgoName
	}

	var returnDecl string
	if !m.retrn.isVoid {
		returnDecl = m.retrn.goName
	}

	call := fmt.Sprintf("C.%s(%s)", m.cFuncName, strings.Join(cArgs, ", "))

	f.writeDocComment(m.doc)
	f.writelnf("func %s(%s) %s {", m.goFuncName, strings.Join(params, ", "), returnDecl)
	f.writeln(strings.Join(cVars, "\n"))
	if m.retrn.isVoid {
		f.writelnf("  %s", call)
	} else {
		f.writelnf("  retC := %s", call)
		if m.retrn.isPrimitive {
			f.writelnf("  return %s(retC)", m.retrn.goName)
		} else if m.retrn.record != nil {
			if m.retrn.isPointer || m.retrn.isSmartPointer {
				f.writelnf("  return %s{sk: retC}", m.retrn.record.goName)
			} else {
				f.writelnf("  return %s{sk: &retC}", m.retrn.record.goName)
			}
		} else {
			fatalf("return type '%s' not supported", m.retrn.goName)
		}
	}
	f.writeln("}")
	f.writeln()
}

func (m method) generateHeader(g generator) {
	f := g.headerFile

	params := make([]string, len(m.params))
	for i, param := range m.params {
		params[i] = param.cParam
	}

	returnDecl := m.retrn.cName
	returnPtr := ""
	if m.retrn.record != nil {
		if m.retrn.isPointer || m.retrn.isSmartPointer {
			returnPtr = "*"
		}
		returnDecl = fmt.Sprintf("%s%s", m.retrn.record.cStructName, returnPtr)
	}

	f.writelnf("%s %s(%s);", returnDecl, m.cFuncName, strings.Join(params, ", "))
}

func (m method) generateCpp(g generator) {
	f := g.cppFile

	params := make([]string, len(m.params))
	args := make([]string, len(m.params))
	for i, param := range m.params {
		params[i] = param.cParam
		args[i] = param.cArg
	}

	returnDecl := m.retrn.cName
	returnPtr := ""
	if m.retrn.record != nil {
		if m.retrn.isPointer || m.retrn.isSmartPointer {
			returnPtr = "*"
		}
		returnDecl = fmt.Sprintf("%s%s", m.retrn.record.cStructName, returnPtr)
	}

	skSpRelease := ""
	if m.retrn.isSmartPointer {
		skSpRelease = ".release()"
	}

	f.writelnf("%s %s(%s) {", returnDecl, m.cFuncName, strings.Join(params, ", "))
	if m.retrn.record != nil {
		if m.retrn.isPointer || m.retrn.isSmartPointer {
			f.writelnf("  return reinterpret_cast<%s *> (%s::%s(%s)%s);",
				m.retrn.record.cStructName,
				m.record.CName,
				m.CName,
				strings.Join(args, ", "),
				skSpRelease)
		} else {
			f.writelnf("  auto ret = (%s::%s(%s)%s);",
				m.record.CName,
				m.CName,
				strings.Join(args, ", "),
				skSpRelease)
			f.writelnf("  return *(reinterpret_cast<%s *> (&ret));",
				m.retrn.record.cStructName,
			)
		}
	} else {
		f.writelnf("  return %s::%s(%s)%s;", m.record.CName, m.CName, strings.Join(args, ", "), skSpRelease)
	}
	f.writeln("}")
	f.writeln()
}
