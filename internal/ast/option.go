package ast

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type Option struct {
	Name  protoreflect.FullName
	Value Constant

	tagTopLevelDef
	tagEnumBody
}

func (v Option) PrintTo(p Printer) {
	p.Write([]byte("option "))
	if v.Name.Parent() == "" {
		p.Write([]byte(v.Name))
	} else {
		fmt.Fprintf(p, "(%s)", v.Name)
	}
	fmt.Fprintf(p, " = ")
	v.Value.PrintTo(p)
	p.Write([]byte(";"))
	p.Newline()
}

type FieldOption struct {
	Name     string
	Constant Constant

	tagEnumBody
}

func (v FieldOption) PrintTo(p Printer) {
	if strings.Contains(v.Name, ".") {
		fmt.Fprintf(p, "(%s)", v.Name)
	} else {
		p.Write([]byte(v.Name))
	}
	fmt.Fprintf(p, " = ")
	v.Constant.PrintTo(p)
	p.Newline()
}
