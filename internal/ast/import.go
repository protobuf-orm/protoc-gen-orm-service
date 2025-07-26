package ast

import (
	"fmt"
)

type Visibility string

const (
	VisibilityWeak   Visibility = "weak"
	VisibilityPublic Visibility = "public"
)

type Import struct {
	Name       string
	Visibility Visibility
}

func (v Import) PrintTo(p Printer) {
	p.Write([]byte("import "))
	if v.Visibility != "" {
		fmt.Fprintf(p, "%s ", string(v.Visibility))
	}
	fmt.Fprintf(p, "%q;", v.Name)
	p.Newline()
}
