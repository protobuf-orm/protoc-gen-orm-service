package ast

import (
	"fmt"
)

type Constant interface {
	PrinterTo
	constant()
}

type Int int

func (v Int) constant() {}
func (v Int) PrintTo(p Printer) {
	fmt.Fprintf(p, "%d", v)
}

type Float float64

func (v Float) constant() {}
func (v Float) PrintTo(p Printer) {
	fmt.Fprintf(p, "%f", v)
}

type String string

func (v String) constant() {}
func (v String) PrintTo(p Printer) {
	fmt.Fprintf(p, "%q", v)
}

type Bool bool

func (v Bool) constant() {}
func (v Bool) PrintTo(p Printer) {
	if v {
		p.Write([]byte("true"))
	} else {
		p.Write([]byte("false"))
	}
}

// TODO:
// type Message struct {}
