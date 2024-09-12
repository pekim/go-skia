package generate

import (
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type methodOverload struct {
	Suffix     string `json:"suffix"`
	cName      string
	goFuncName string
	cFuncName  string
	record     *record
	doc        string
	isStatic   bool
	params     []param
	resultType clang.Type
	retrn      typ
}

type method struct {
	CName         string            `json:"name"`
	Overloads     []*methodOverload `json:"overloads"`
	enrichedCount int
}

func (m *method) enrich(record *record, cursor clang.Cursor) {
	var overload *methodOverload
	if len(m.Overloads) == 0 {
		m.Overloads = []*methodOverload{{}}
	}
	if m.enrichedCount >= len(m.Overloads) {
		fatalf("record %s, method %s, %d overloads configured, but more encountered",
			m.Overloads[0].record.CName, m.CName, len(m.Overloads))
	}
	overload = m.Overloads[m.enrichedCount]
	if overload == nil {
		m.enrichedCount++
		return
	}

	overload.cName = m.CName
	overload.record = record
	overload.isStatic = cursor.CXXMethod_IsStatic()
	if overload.isStatic {
		overload.goFuncName = fmt.Sprintf("%s%s%s", record.goName, m.CName, overload.Suffix)
	} else {
		overload.goFuncName = fmt.Sprintf("%s%s", goExportedName(m.CName), overload.Suffix)
	}
	overload.cFuncName = fmt.Sprintf("misk_%s_%s%s", record.goName, m.CName, overload.Suffix)
	overload.doc = cursor.RawCommentText()
	overload.resultType = cursor.ResultType()

	paramCount := int(cursor.NumArguments())
	overload.params = make([]param, paramCount)
	for i := 0; i < paramCount; i++ {
		arg := cursor.Argument(uint32(i))
		param := newParam(i, arg)
		overload.params[i] = param
	}

	m.enrichedCount++
}

func (m *method) enrich2(api api) {
	for i := range m.Overloads {
		overload := m.Overloads[i]
		for i := range overload.params {
			param := &overload.params[i]
			param.enrich2(api)
		}
		overload.retrn = mustTypFromClangType(overload.resultType, api)
	}
}

func (m method) generate(g generator) {
	if len(m.Overloads) == 0 {
		fatalf("method %s, 0 overloads configured, and none enriched", m.CName)
	}
	if m.enrichedCount < len(m.Overloads) {
		fatalf("record %s, method %s, only %d of %d overloads enriched",
			m.Overloads[0].record.CName, m.CName, m.enrichedCount, len(m.Overloads))
	}

	for _, method := range m.Overloads {
		if method != nil {
			method.generate(g)
		}
	}
}

func (m methodOverload) generate(g generator) {
	m.generateGo(g)
	m.generateHeader(g)
	m.generateCpp(g)
}

func (m methodOverload) generateGo(g generator) {
	f := g.goFile

	argsCount := len(m.params)
	argsOffset := 0
	if !m.isStatic {
		argsCount++
		argsOffset = 1
	}
	params := make([]string, len(m.params))
	cVars := make([]string, argsCount)
	cArgs := make([]string, argsCount)
	if !m.isStatic {
		cVars[0] = "c_obj := o.sk"
		cArgs[0] = "c_obj"
	}
	for i, param := range m.params {
		params[i] = fmt.Sprintf("%s %s", param.goName, param.typ.goName)
		cVars[argsOffset+i] = param.cgoVar
		cArgs[argsOffset+i] = param.cgoName
	}

	var returnDecl string
	if !m.retrn.isVoid {
		returnDecl = m.retrn.goName
	}

	var receiver string
	if !m.isStatic {
		receiver = fmt.Sprintf("(o %s)", m.record.goName)
	}
	call := fmt.Sprintf("C.%s(%s)", m.cFuncName, strings.Join(cArgs, ", "))

	f.writeDocComment(m.doc)
	f.writelnf("func %s %s(%s) %s {", receiver, m.goFuncName, strings.Join(params, ", "), returnDecl)
	if len(cVars) > 0 {
		f.writeln(strings.Join(cVars, "\n"))
	}
	if m.retrn.isVoid {
		f.writelnf("  %s", call)
	} else {
		f.writelnf("  retC := %s", call)
		if m.retrn.isPrimitive {
			f.writelnf("  return %s(retC)", m.retrn.goName)
		} else if m.retrn.enum != nil {
			f.writelnf("  return %s(retC)", m.retrn.enum.goName)
		} else if m.retrn.record != nil {
			if m.retrn.record.NoWrapper {
				f.writelnf("  return %s(retC)", m.retrn.record.goName)
			} else if m.retrn.isPointer || m.retrn.isSmartPointer {
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

func (m methodOverload) generateHeader(g generator) {
	f := g.headerFile

	paramsCount := len(m.params)
	paramOffset := 0
	if !m.isStatic {
		paramsCount++
		paramOffset = 1
	}
	params := make([]string, paramsCount)
	if !m.isStatic {
		params[0] = fmt.Sprintf("%s *c_obj", m.record.cStructName)
	}
	for i, param := range m.params {
		params[paramOffset+i] = param.cParam
	}

	returnDecl := m.retrn.cName
	returnPtr := ""
	if m.retrn.enum != nil {
		returnDecl = m.retrn.enum.cType.cgoName
	} else if m.retrn.record != nil {
		if m.retrn.isPointer || m.retrn.isSmartPointer {
			returnPtr = "*"
		}
		returnDecl = fmt.Sprintf("%s%s", m.retrn.record.cStructName, returnPtr)
	}

	f.writelnf("%s %s(%s);", returnDecl, m.cFuncName, strings.Join(params, ", "))
}

func (m methodOverload) generateCpp(g generator) {
	f := g.cppFile

	paramsCount := len(m.params)
	paramOffset := 0
	if !m.isStatic {
		paramsCount++
		paramOffset = 1
	}
	params := make([]string, paramsCount)
	args := make([]string, len(m.params))
	if !m.isStatic {
		params[0] = fmt.Sprintf("%s *c_obj", m.record.cStructName)
	}
	for i, param := range m.params {
		params[paramOffset+i] = param.cParam
		args[i] = param.cArg
	}

	returnDecl := m.retrn.cName
	returnPtr := ""
	if m.retrn.enum != nil {
		returnDecl = m.retrn.enum.cType.cgoName
	} else if m.retrn.record != nil {
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
		if m.isStatic {
			if m.retrn.isPointer || m.retrn.isSmartPointer {
				f.writelnf("  return reinterpret_cast<%s *> (%s::%s(%s)%s);",
					m.retrn.record.cStructName,
					m.record.CName,
					m.cName,
					strings.Join(args, ", "),
					skSpRelease)
			} else {
				f.writelnf("  auto ret = (%s::%s(%s)%s);",
					m.record.CName,
					m.cName,
					strings.Join(args, ", "),
					skSpRelease)
				f.writelnf("  return *(reinterpret_cast<%s *> (&ret));",
					m.retrn.record.cStructName,
				)
			}
		} else {
			// TODO

			if m.retrn.isPointer || m.retrn.isSmartPointer {
				f.writelnf("  return reinterpret_cast<%s *> (%s::%s(%s)%s);",
					m.retrn.record.cStructName,
					m.record.CName,
					m.cName,
					strings.Join(args, ", "),
					skSpRelease)
			} else {
				f.writelnf("  auto ret = (%s::%s(%s)%s);",
					m.record.CName,
					m.cName,
					strings.Join(args, ", "),
					skSpRelease)
				f.writelnf("  return *(reinterpret_cast<%s *> (&ret));",
					m.retrn.record.cStructName,
				)
			}
		}
	} else {
		if m.isStatic {
			f.writelnf("  return %s::%s(%s)%s;", m.record.CName, m.cName, strings.Join(args, ", "), skSpRelease)
		} else {
			f.writelnf("  return reinterpret_cast<%s*>(c_obj)->%s(%s)%s;",
				m.record.CName,
				m.cName,
				strings.Join(args, ", "),
				skSpRelease,
			)
		}
	}
	f.writeln("}")
	f.writeln()
}
