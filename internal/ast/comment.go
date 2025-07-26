package ast

import (
	"strings"
)

type Comment string

func (v Comment) _topLevelDef()      {}
func (v Comment) _messageBody()      {}
func (v Comment) _messageOneofBody() {}
func (v Comment) _serviceBody()      {}
func (v Comment) PrintTo(p Printer) {
	lines := strings.SplitSeq(string(v), "\n")
	for line := range lines {
		p.Write([]byte("// "))
		p.Write([]byte(line))
		p.Newline()
	}
}

type MultilineComment string

func (v MultilineComment) _topLevelDef()      {}
func (v MultilineComment) _messageBody()      {}
func (v MultilineComment) _messageOneofBody() {}
func (v MultilineComment) _serviceBody()      {}
func (v MultilineComment) PrintTo(p Printer) {
	p.Write([]byte("/**"))
	p.Newline()
	lines := strings.SplitSeq(string(v), "\n")
	for line := range lines {
		p.Write([]byte(" * "))
		p.Write([]byte(line))
		p.Newline()
	}
	p.Write([]byte(" */"))
	p.Newline()
}
