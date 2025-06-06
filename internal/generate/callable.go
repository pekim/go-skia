package generate

import (
	"github.com/go-clang/clang-v15/clang"
)

type callable struct {
	cppName       string
	overloads     []*callableOverload
	record        *record
	enrichedCount int
}

// enrich1 enriches one of the callable's methods.
func (c *callable) enrich1(record *record, cursor clang.Cursor) {
	c.record = record

	var overload *callableOverload
	if len(c.overloads) == 0 {
		// If no overloads were configured in JSON it implies that there is just one.
		c.overloads = []*callableOverload{{}}
	}
	if c.enrichedCount >= len(c.overloads) {
		if record != nil {
			fatalf("record %s, method %s, %d overloads configured, but more encountered",
				c.record.cppName, c.cppName, len(c.overloads))
		} else {
			fatalf("function %s, %d overloads configured, but more encountered", c.cppName, len(c.overloads))
		}
	}

	// Overloads configured in JSON are expected to be in the same order as they are encountered
	// in the AST for a translation unit.
	//
	// If an overload is configured as 'null' in JSON that indicates that is it not to be
	// generated. So it won't be enriched.
	overload = c.overloads[c.enrichedCount]
	if overload != nil {
		overload.enrich1(c, record, cursor)
	}
	c.enrichedCount++
}

func (c *callable) enrich2(record *record, api api) {
	c.record = record
	if len(c.overloads) == 0 {
		if c.record != nil {
			fatalf("record %s, method %s, has 0 overloads configured, and none enriched", c.record.cppName, c.cppName)
		} else {
			fatalf("function %s has 0 overloads configured, and none enriched", c.cppName)
		}
	}

	for i := range c.overloads {
		overload := c.overloads[i]
		if overload != nil {
			overload.enrich2(api)
		}
	}
}

func (c callable) generate(g generator) {
	if len(c.overloads) == 0 {
		if c.record != nil {
			fatalf("record %s, method %s, has 0 overloads configured, and none enriched", c.record.cppName, c.cppName)
		} else {
			fatalf("function %s has 0 overloads configured, and none enriched", c.cppName)
		}
	}
	if c.enrichedCount < len(c.overloads) {
		if c.record != nil {
			fatalf("record %s, method %s, only %d of %d overloads enriched",
				c.overloads[0].record.cppName, c.cppName, c.enrichedCount, len(c.overloads))
		} else {
			fatalf("function %s, only %d of %d overloads enriched",
				c.cppName, c.enrichedCount, len(c.overloads))
		}
	}

	for _, overload := range c.overloads {
		if overload != nil {
			overload.generate(g)
		}
	}
}
