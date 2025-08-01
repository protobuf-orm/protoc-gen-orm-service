package ast

import (
	"io"
	"strings"
)

type Printer interface {
	Write(b []byte) (int, error)
	PrintTypename(v string)

	Indent()
	Dedent()
	Newline()
}

type PrinterTo interface {
	PrintTo(p Printer)
}

func Scope(p Printer, f func(), closer string) {
	p.Newline()
	p.Indent()
	f()
	p.Dedent()
	p.Write([]byte(closer))
	p.Newline()
}

func ScopedPrint[T PrinterTo](p Printer, vs []T, closer string) {
	Scope(p, func() {
		for _, v := range vs {
			v.PrintTo(p)
		}
	}, closer)
}

type printer struct {
	io.Writer
	pkg string

	depth  int
	offset int
}

func NewPrinter(w io.Writer, pkg string) Printer {
	return &printer{Writer: w, pkg: pkg}
}

func (p *printer) Write(b []byte) (n int, err error) {
	if p.offset == 0 {
		tab := strings.Repeat("\t", p.depth)
		n, err = p.Writer.Write([]byte(tab))
		if err != nil {
			return
		}
	}

	n, err = p.Writer.Write(b)
	p.offset += n
	return
}

func (p *printer) PrintTypename(v string) {
	name, ok := strings.CutPrefix(v, p.pkg)
	if ok {
		name = name[1:]
	} else {
		name = v
	}
	p.Write([]byte(name))
}

func (p *printer) Indent() { p.depth++ }
func (p *printer) Dedent() { p.depth-- }
func (p *printer) Newline() {
	p.offset = 0
	p.Writer.Write([]byte("\n"))
}
