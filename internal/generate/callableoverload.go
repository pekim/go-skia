package generate

import (
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type callableOverload struct {
	Suffix            string `json:"suffix"`
	cppName           string
	goFuncName        string
	cFuncName         string
	record            *record
	doc               string
	isStatic          bool
	isNonStaticMethod bool    // Go function is a method, and C function has a receiver first parameter
	Params            []param `json:"params"`
	cParamsDecl       string
	cppArgs           string
	resultType        clang.Type
	retrn             typ
}

func (o *callableOverload) enrich1(callable *callable, record *record, cursor clang.Cursor) {
	o.cppName = callable.CppName
	o.record = record
	o.isStatic = cursor.CXXMethod_IsStatic()
	o.isNonStaticMethod = o.record != nil && !o.isStatic
	if o.isStatic {
		o.goFuncName = fmt.Sprintf("%s%s%s", record.goName, o.cppName, o.Suffix)
	} else {
		o.goFuncName = fmt.Sprintf("%s%s", goExportedName(o.cppName), o.Suffix)
	}
	if o.record != nil {
		o.cFuncName = fmt.Sprintf("misk_%s_%s%s", o.record.goName, o.cppName, o.Suffix)
	} else {
		cName := strings.ReplaceAll(o.cppName, "::", "") // remove any "::" in the name
		o.cFuncName = fmt.Sprintf("misk_%s%s", cName, o.Suffix)
	}
	o.doc = cursor.RawCommentText()
	o.resultType = cursor.ResultType()

	paramCount := int(cursor.NumArguments())
	if len(o.Params) > 0 {
		if len(o.Params) > 0 && len(o.Params) != paramCount {
			fatalf("function %s, expected %d params but have %d", o.cppName, len(o.Params), paramCount)
		}
		// Do not make o.Params, as it's already present from unmarshalled JSON.
	} else {
		o.Params = make([]param, paramCount)
	}
	for i := 0; i < paramCount; i++ {
		arg := cursor.Argument(uint32(i))
		param := newParam(i, arg, o.Params[i].ValueNil)
		o.Params[i] = param
	}
}

func (o *callableOverload) enrich2(api api) {
	for i := range o.Params {
		param := &o.Params[i]
		param.enrich2(api)
	}
	o.setCParamsDecl()
	o.retrn = mustTypFromClangType(o.resultType, api)
}

func (o *callableOverload) setCParamsDecl() {
	var params []string
	if o.isNonStaticMethod {
		params = append(params, fmt.Sprintf("%s *c_obj", o.record.cStructName))
	}
	for _, param := range o.Params {
		if param.ValueNil {
			// no more params
			break
		}
		params = append(params, param.cParam)
	}
	o.cParamsDecl = strings.Join(params, ", ")
}

func (o callableOverload) generate(g generator) {
	o.generateGo(g)
	o.generateHeader(g)
	o.generateCpp(g)
}
