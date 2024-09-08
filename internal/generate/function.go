package generate

import (
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type functionOverload struct {
	Suffix     string `json:"suffix"`
	cName      string
	goFuncName string
	cFuncName  string
	doc        string
	params     []param
	resultType clang.Type
	retrn      typ
}

type function struct {
	CName         string              `json:"name"`
	Overloads     []*functionOverload `json:"overloads"`
	cursor        clang.Cursor
	enrichedCount int
}

func (f *function) enrich(cursor clang.Cursor) {
	var overload *functionOverload
	if len(f.Overloads) == 0 {
		f.Overloads = []*functionOverload{{}}
	}
	if f.enrichedCount >= len(f.Overloads) {
		fatalf("function %s, %d overloads configured, but more encountered", f.CName, len(f.Overloads))
	}
	overload = f.Overloads[f.enrichedCount]
	if overload == nil {
		f.enrichedCount++
		return
	}

	overload.cName = f.CName
	overload.goFuncName = fmt.Sprintf("%s%s", f.CName, overload.Suffix)
	overload.cFuncName = fmt.Sprintf("misk_%s%s", f.CName, overload.Suffix)
	overload.doc = cursor.RawCommentText()
	overload.resultType = cursor.ResultType()

	paramCount := int(cursor.NumArguments())
	overload.params = make([]param, paramCount)
	for i := 0; i < paramCount; i++ {
		arg := cursor.Argument(uint32(i))
		param := newParam(i, arg)
		overload.params[i] = param
	}

	f.enrichedCount++
}

func (f *function) enrich2(api api) {
	for i := range f.Overloads {
		overload := f.Overloads[i]
		overload.retrn = mustTypFromClangType(overload.resultType, api)
	}
}

func (f function) generate(g generator) {
	if len(f.Overloads) == 0 {
		fatalf("function %s, 0 overloads configured, and none enriched", f.CName)
	}
	if f.enrichedCount < len(f.Overloads) {
		fatalf("function %s, only %d of %d overloads enriched", f.CName, f.enrichedCount, len(f.Overloads))
	}

	for _, function := range f.Overloads {
		if function != nil {
			function.generate(g)
		}
	}
}

func (f functionOverload) generate(g generator) {
	f.generateGo(g)
	f.generateHeader(g)
	f.generateCpp(g)
}

func (f functionOverload) generateGo(g generator) {
	file := g.goFile

	params := make([]string, len(f.params))
	cVars := make([]string, len(f.params))
	cArgs := make([]string, len(f.params))
	for i, param := range f.params {
		params[i] = fmt.Sprintf("%s %s", param.goName, param.typ.goName)
		cVars[i] = param.cgoVar
		cArgs[i] = param.cgoName
	}

	var returnDecl string
	if !f.retrn.isVoid {
		returnDecl = f.retrn.goName
	}

	call := fmt.Sprintf("C.%s(%s)", f.cFuncName, strings.Join(cArgs, ", "))

	file.writeDocComment(f.doc)
	file.writelnf("func %s(%s) %s {", f.goFuncName, strings.Join(params, ", "), returnDecl)
	if len(cVars) > 0 {
		file.writeln(strings.Join(cVars, "\n"))
	}
	if f.retrn.isVoid {
		file.writelnf("  %s", call)
	} else {
		file.writelnf("  retC := %s", call)
		if f.retrn.isPrimitive {
			file.writelnf("  return %s(retC)", f.retrn.goName)
		} else if f.retrn.record != nil {
			if f.retrn.isPointer || f.retrn.isSmartPointer {
				file.writelnf("  return %s{sk: retC}", f.retrn.record.goName)
			} else {
				file.writelnf("  return %s{sk: &retC}", f.retrn.record.goName)
			}
		} else {
			fatalf("return type '%s' not supported", f.retrn.goName)
		}
	}
	file.writeln("}")
	file.writeln()
}

func (f functionOverload) generateHeader(g generator) {
	file := g.headerFile

	params := make([]string, len(f.params))
	for i, param := range f.params {
		params[i] = param.cParam
	}

	returnDecl := f.retrn.cName
	returnPtr := ""
	if f.retrn.record != nil {
		if f.retrn.isPointer || f.retrn.isSmartPointer {
			returnPtr = "*"
		}
		returnDecl = fmt.Sprintf("%s%s", f.retrn.record.cStructName, returnPtr)
	}

	file.writelnf("%s %s(%s);", returnDecl, f.cFuncName, strings.Join(params, ", "))
}

func (f functionOverload) generateCpp(g generator) {
	file := g.cppFile

	params := make([]string, len(f.params))
	args := make([]string, len(f.params))
	for i, param := range f.params {
		params[i] = param.cParam
		args[i] = param.cArg
	}

	returnDecl := f.retrn.cName
	returnPtr := ""
	if f.retrn.record != nil {
		if f.retrn.isPointer || f.retrn.isSmartPointer {
			returnPtr = "*"
		}
		returnDecl = fmt.Sprintf("%s%s", f.retrn.record.cStructName, returnPtr)
	}

	var returnConst string
	var constCastStart string
	var constCastEnd string
	if f.retrn.isConst {
		returnConst = "const"
		constCastStart = fmt.Sprintf("const_cast<%s *>(", f.retrn.record.cStructName)
		constCastEnd = ")"
	}

	skSpRelease := ""
	if f.retrn.isSmartPointer {
		skSpRelease = ".release()"
	}

	file.writelnf("%s %s(%s) {", returnDecl, f.cFuncName, strings.Join(params, ", "))
	if f.retrn.record != nil {
		if f.retrn.isPointer || f.retrn.isSmartPointer {
			file.writelnf("  return %sreinterpret_cast<%s %s *> (%s(%s)%s)%s;",
				constCastStart,
				returnConst,
				f.retrn.record.cStructName,
				f.cName,
				strings.Join(args, ", "),
				skSpRelease,
				constCastEnd)
		} else {
			file.writelnf("  auto ret = (%s(%s)%s);",
				f.cName,
				strings.Join(args, ", "),
				skSpRelease)
			file.writelnf("  return *(reinterpret_cast<%s %s *> (&ret));",
				returnConst,
				f.retrn.record.cStructName,
			)
		}
	} else {
		file.writelnf("  return %s(%s)%s;", f.cName, strings.Join(args, ", "), skSpRelease)
	}
	file.writeln("}")
	file.writeln()
}
