package generate

import (
	"fmt"

	"github.com/go-clang/clang-v15/clang"
)

type callable struct {
	CppName       string              `json:"name"`
	Overloads     []*callableOverload `json:"overloads"`
	enrichedCount int
}

// enrich1 enriches one of the callable's methods.
func (c *callable) enrich1(record *record, cursor clang.Cursor) {
	var overload *callableOverload
	if len(c.Overloads) == 0 {
		// If no overloads were configured in JSON it implies that there is just one.
		c.Overloads = []*callableOverload{{}}
	}
	if c.enrichedCount >= len(c.Overloads) {
		if record != nil {
			fatalf("record %s, method %s, %d overloads configured, but more encountered",
				c.Overloads[0].record.CppName, c.CppName, len(c.Overloads))
		} else {
			fatalf("function %s, %d overloads configured, but more encountered", c.CppName, len(c.Overloads))
		}
	}

	// Overloads configured in JSON are expected to be in the same order as they are encountered
	// in the AST for a translation unit.
	overload = c.Overloads[c.enrichedCount]
	if overload == nil {
		// The overload was configured as 'null' in JSON. That indicates that is it not to be
		// generated. So don't enrich it, skip it.
		c.enrichedCount++
		return
	}

	overload.cppName = c.CppName
	overload.record = record
	overload.isStatic = cursor.CXXMethod_IsStatic()
	if overload.isStatic {
		overload.goFuncName = fmt.Sprintf("%s%s%s", record.goName, c.CppName, overload.Suffix)
	} else {
		overload.goFuncName = fmt.Sprintf("%s%s", goExportedName(c.CppName), overload.Suffix)
	}
	overload.cFuncName = fmt.Sprintf("misk_%s_%s%s", record.goName, c.CppName, overload.Suffix)
	overload.doc = cursor.RawCommentText()
	overload.resultType = cursor.ResultType()

	paramCount := int(cursor.NumArguments())
	overload.params = make([]param, paramCount)
	for i := 0; i < paramCount; i++ {
		arg := cursor.Argument(uint32(i))
		param := newParam(i, arg)
		overload.params[i] = param
	}

	c.enrichedCount++
}

func (c *callable) enrich2(api api) {
	for i := range c.Overloads {
		overload := c.Overloads[i]
		if overload == nil {
			continue
		}
		for i := range overload.params {
			param := &overload.params[i]
			param.enrich2(api)
		}
		overload.retrn = mustTypFromClangType(overload.resultType, api)
	}
}

func (c callable) generate(g generator) {
	if len(c.Overloads) == 0 {
		fatalf("method %s, 0 overloads configured, and none enriched", c.CppName)
	}
	if c.enrichedCount < len(c.Overloads) {
		fatalf("record %s, method %s, only %d of %d overloads enriched",
			c.Overloads[0].record.CppName, c.CppName, c.enrichedCount, len(c.Overloads))
	}

	for _, method := range c.Overloads {
		if method != nil {
			method.generate(g)
		}
	}
}
