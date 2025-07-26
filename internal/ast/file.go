package ast

import (
	"fmt"
)

type File struct {
	Edition Edition
	Package string
	Imports []Import
	Options []Option

	Defs []TopLevelDef
}

type TopLevelDef interface {
	Node
	_topLevelDef()
}

type tagTopLevelDef struct{}

func (a tagTopLevelDef) _topLevelDef() {}

func (v File) PrintTo(p Printer) {
	v.Edition.PrintTo(p)
	p.Newline()
	fmt.Fprintf(p, "package %s;", v.Package)
	p.Newline()
	p.Newline()
	for _, w := range v.Imports {
		w.PrintTo(p)
	}
	if len(v.Imports) > 0 {
		p.Newline()
	}
	for _, w := range v.Options {
		w.PrintTo(p)
	}
	if len(v.Options) > 0 {
		p.Newline()
	}
	for i, w := range v.Defs {
		w.PrintTo(p)
		if i+1 < len(v.Defs) {
			p.Newline()
		}
	}
}
