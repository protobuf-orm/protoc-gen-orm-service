package ast

import (
	"fmt"
)

type Edition struct {
	Keyword string
	Value   string
}

var (
	SyntaxProto2 = Edition{"syntax", "proto2"}
	SyntaxProto3 = Edition{"syntax", "proto3"}
	Edition2023  = Edition{"edition", "2023"}
)

func (v Edition) PrintTo(p Printer) {
	fmt.Fprintf(p, "%s = %q;", v.Keyword, v.Value)
	p.Newline()
}
