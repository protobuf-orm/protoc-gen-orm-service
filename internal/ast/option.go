package ast

import (
	"fmt"
)

type Option struct {
	Known bool
	Name  string
	Value Constant

	tagTopLevelDef
	tagEnumBody
}

func (v Option) PrintTo(p Printer) {
	p.Write([]byte("option "))
	if v.Known {
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
	Known bool
	Name  string
	Value Constant

	tagEnumBody
}

var (
	FeaturesFieldPresenceImplicit = knownFieldOption("features.field_presence", Value("IMPLICIT"))
	FeaturesFieldPresenceExplicit = knownFieldOption("features.field_presence", Value("EXPLICIT"))
)

func knownFieldOption(name string, value Constant) FieldOption {
	return FieldOption{
		Known: true,
		Name:  name,
		Value: value,
	}
}

func (v FieldOption) PrintTo(p Printer) {
	if v.Known {
		p.Write([]byte(v.Name))
	} else {
		fmt.Fprintf(p, "(%s)", v.Name)
	}
	fmt.Fprintf(p, " = ")
	v.Value.PrintTo(p)
	p.Newline()
}
