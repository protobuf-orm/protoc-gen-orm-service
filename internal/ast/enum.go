package ast

import "fmt"

type Enum struct {
	Name string
	Body []EnumBody
}

func (v Enum) PrintTo(p Printer) {
	fmt.Fprintf(p, "enum %s {", v.Name)
	ScopedPrint(p, v.Body, "}")
}

type EnumBody interface {
	Node
	_enumBody()
}

type tagEnumBody struct{}

func (tagEnumBody) _enumBody() {}

type EnumField struct {
	Name    string
	Number  int
	Options []FieldOption

	tagEnumBody
}

func (v EnumField) PrintTo(p Printer) {
	fmt.Fprintf(p, "%s = %d", v.Name, v.Number)
	if len(v.Options) > 0 {
		p.Write([]byte(" ["))
		ScopedPrint(p, v.Options, "];")
	} else {
		p.Write([]byte(";"))
		p.Newline()
	}
}
