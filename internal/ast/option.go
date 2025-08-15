package ast

import (
	"fmt"
)

var (
	FeaturesFieldPresenceImplicit = knownOption("features.field_presence", Value("IMPLICIT"))
	FeaturesFieldPresenceExplicit = knownOption("features.field_presence", Value("EXPLICIT"))
)

func knownOption(name string, value Constant) Option {
	return Option{
		Known: true,
		Name:  name,
		Value: value,
	}
}

type Option struct {
	Known bool
	Name  string
	Value Constant

	tagTopLevelDef
	tagMessageBody
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

func (o Option) WithinField() FieldOption {
	return FieldOption{
		Known: o.Known,
		Name:  o.Name,
		Value: o.Value,
	}
}

type FieldOption struct {
	Known bool
	Name  string
	Value Constant

	tagEnumBody
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
