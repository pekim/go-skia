package generate

import (
	"fmt"
	"strconv"

	"github.com/go-clang/clang-v15/clang"
)

type classCtor struct {
	class  *class
	params []param

	nameSuffix string
	cFuncName  string
}

func newClassCtor(class *class, cursor clang.Cursor) classCtor {
	ctor := classCtor{
		class: class,
	}

	numParam := int(cursor.NumArguments())
	ctor.params = make([]param, 0, numParam)
	for i := 0; i < numParam; i++ {
		arg := cursor.Argument(uint32(i))
		param := newParam(arg, i)
		ctor.params = append(ctor.params, param)
	}

	return ctor
}

func (c *classCtor) generate(g *generator) {
	for _, param := range c.params {
		if !param.supported() {
			return
		}
	}

	c.class.ctorN++
	if c.class.ctorN > 1 {
		c.nameSuffix = strconv.Itoa(c.class.ctorN)
	}
	c.cFuncName = fmt.Sprintf("skia_new_%s%s", c.class.cName, c.nameSuffix)

	c.generateGo(g)
	c.generateHeader(g)
	c.generateCpp(g)
}
