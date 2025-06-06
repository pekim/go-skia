package generate

import (
	"fmt"
	"strings"
)

func (o callableOverload) generateGo(g generator) {
	f := g.goFile

	// make function's params declarations
	var params []string
	for _, param := range o.Params {
		if param.ValueNil {
			// no more params
			break
		}
		params = append(params, fmt.Sprintf("%s %s", param.goName, param.typGoName))
	}

	// make return declaration
	var returnDecl string
	if !o.retrn.isVoid {
		returnDecl = o.retrn.goName
	}
	// make receiver if a non-static class method
	var receiver string
	if o.isNonStaticMethod {
		if o.record.noWrapper {
			receiver = fmt.Sprintf("(o *%s)", o.record.goName)
		} else {
			receiver = fmt.Sprintf("(o %s)", o.record.goName)
		}
	}

	f.write(o.doc)
	f.writelnf("func %s %s(%s) %s {", receiver, o.goFuncName, strings.Join(params, ", "), returnDecl)
	o.generateGoBody(g)
	f.writeln("}")
	f.writeln()
}

func (o callableOverload) generateGoBody(g generator) {
	f := g.goFile

	var cVars []string
	var cArgs []string
	firstCVar, firstCArg, haveExtraCArg := o.firstGoCArg()
	if haveExtraCArg {
		cVars = append(cVars, firstCVar)
		cArgs = append(cArgs, firstCArg)
	}
	for _, param := range o.Params {
		if param.ValueNil {
			// no more params
			break
		}
		cVars = append(cVars, param.cgoVar)
		cArgs = append(cArgs, param.cName)
	}

	// build the call to the function/method
	call := fmt.Sprintf("C.%s(%s)", o.cFuncName, strings.Join(cArgs, ", "))

	// write the creation of the vars that will be passed as args to the C function
	if len(cVars) > 0 {
		f.writeln(strings.Join(cVars, "\n"))
	}

	if o.retrn.isVoid {
		f.writelnf("  %s", call)
	} else {
		f.writelnf("  retC := %s", call)
		if o.retrn.isPrimitive {
			f.writelnf("  return %s(retC)", o.retrn.goName)
		} else if o.retrn.enum != nil {
			f.writelnf("  return %s(retC)", o.retrn.enum.goName)
		} else if o.retrn.record != nil {
			if o.retrn.record.noWrapper {
				f.writelnf("  return %s(retC)", o.retrn.record.goName)
			} else if o.retrn.isPointer || o.retrn.isSmartPointer {
				f.writelnf("  return %s{sk: retC}", o.retrn.record.goName)
			} else {
				f.writelnf("  return %s{sk: &retC}", o.retrn.record.goName)
			}
		} else if (o.retrn.isPointer || o.retrn.isSmartPointer) && o.retrn.subTyp.record != nil {
			f.writelnf("  return %s{sk: retC}", o.retrn.subTyp.record.goName)
		} else if o.retrn.isLValueReference {
			if o.retrn.subTyp.record.noWrapper {
				f.writelnf("  return %s(retC)", o.retrn.subTyp.record.goName)
			} else {
				f.writelnf("  return %s{sk: &retC}", o.retrn.subTyp.record.goName)
			}
		} else if o.retrn.goName == "string" {
			f.writelnf("  return C.GoString( retC)")
		} else if o.retrn.typedef != nil {
			f.writelnf("  return %s(retC)", o.retrn.goName)
		} else {
			fatalf("return type '%s' not supported", o.retrn.goName)
		}
	}
}

// firstCArg provides an (extra) first C arg if the overload is a non-static
// class function.
func (o callableOverload) firstGoCArg() (string, string, bool) {
	if !o.isNonStaticMethod {
		return "", "", false
	}

	cVar := "c_obj := o.sk"
	if o.record.noWrapper {
		cVar = fmt.Sprintf("c_obj := (*C.%s)(o)", o.record.cStructName)
	}
	cArg := "c_obj"
	return cVar, cArg, true
}
