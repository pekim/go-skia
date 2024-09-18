package generate

import (
	"github.com/go-clang/clang-v15/clang"
)

type callable struct {
	CppName       string              `json:"name"`
	Overloads     []*callableOverload `json:"overloads"`
	record        *record
	enrichedCount int
}

// enrich1 enriches one of the callable's methods.
func (c *callable) enrich1(record *record, cursor clang.Cursor) {
	c.record = record

	var overload *callableOverload
	if len(c.Overloads) == 0 {
		// If no overloads were configured in JSON it implies that there is just one.
		c.Overloads = []*callableOverload{{}}
	}
	if c.enrichedCount >= len(c.Overloads) {
		if record != nil {
			fatalf("record %s, method %s, %d overloads configured, but more encountered",
				c.record.CppName, c.CppName, len(c.Overloads))
		} else {
			fatalf("function %s, %d overloads configured, but more encountered", c.CppName, len(c.Overloads))
		}
	}

	// Overloads configured in JSON are expected to be in the same order as they are encountered
	// in the AST for a translation unit.
	//
	// If an overload is configured as 'null' in JSON that indicates that is it not to be
	// generated. So it won't be enriched.
	overload = c.Overloads[c.enrichedCount]
	if overload != nil {
		overload.enrich1(c, record, cursor)
	}
	c.enrichedCount++
}

func (c *callable) enrich2(api api) {
	for i := range c.Overloads {
		overload := c.Overloads[i]
		if overload != nil {
			overload.enrich2(api)
		}
	}
}

func (c callable) generate(g generator) {
	if len(c.Overloads) == 0 {
		if c.record != nil {
			fatalf("record %s, method %s, has 0 overloads configured, and none enriched", c.record.CppName, c.CppName)
		} else {
			fatalf("function %s has 0 overloads configured, and none enriched", c.CppName)
		}
	}
	if c.enrichedCount < len(c.Overloads) {
		if c.record != nil {
			fatalf("record %s, method %s, only %d of %d overloads enriched",
				c.Overloads[0].record.CppName, c.CppName, c.enrichedCount, len(c.Overloads))
		} else {
			fatalf("function %s, only %d of %d overloads enriched",
				c.CppName, c.enrichedCount, len(c.Overloads))
		}
	}

	for _, overload := range c.Overloads {
		if overload != nil {
			overload.generate(g)
		}
	}
}
