package generate

type class struct {
	cName  string
	goName string
	enums  []enum
}

func (c class) generate(fileGo *genFile) {
	for _, enum := range c.enums {
		enum.generate(fileGo)
	}
}
