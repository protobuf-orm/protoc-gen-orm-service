package ast

import "fmt"

type Message struct {
	Name string
	Body []MessageBody

	tagTopLevelDef
}

func (v Message) PrintTo(p Printer) {
	fmt.Fprintf(p, "message %s {", getName(v.Name))
	ScopedPrint(p, v.Body, "}")
}

type MessageBody interface {
	Node
	_messageBody()
}

type tagMessageBody struct{}

func (tagMessageBody) _messageBody() {}

type MessageField struct {
	Label  string
	Type   string
	Name   string
	Number int

	tagMessageBody
}

func (v MessageField) PrintTo(p Printer) {
	if v.Label != "" {
		fmt.Fprintf(p, "%s ", v.Label)
	}
	p.PrintTypename(v.Type)
	fmt.Fprintf(p, " %s = %d;", v.Name, v.Number)
	p.Newline()
}

type MessageOneof struct {
	Name string
	Body []MessageOneofBody

	tagMessageBody
}

type MessageOneofBody interface {
	Node
	_messageOneofBody()
}

type tagMessageOneofBody struct{}

func (tagMessageOneofBody) _messageOneofBody() {}

func (v MessageOneof) PrintTo(p Printer) {
	fmt.Fprintf(p, "oneof %s {", v.Name)
	ScopedPrint(p, v.Body, "}")
}

type MessageOneofField struct {
	Type   string
	Name   string
	Number int

	tagMessageOneofBody
}

func (v MessageOneofField) PrintTo(p Printer) {
	p.PrintTypename(v.Type)
	fmt.Fprintf(p, " %s = %d;", v.Name, v.Number)
	p.Newline()
}
