package ast

import (
	"fmt"
)

type Rpc struct {
	Name string

	RequestStream  bool
	RequestType    string
	ResponseStream bool
	ResponseType   string

	tagServiceBody
}

func (v Rpc) PrintTo(p Printer) {
	fmt.Fprintf(p, "rpc %s(", v.Name)
	if v.RequestStream {
		p.Write([]byte("stream "))
	}
	p.PrintTypename(v.RequestType)
	p.Write([]byte(") returns ("))
	if v.ResponseStream {
		p.Write([]byte("stream "))
	}
	p.PrintTypename(v.ResponseType)
	p.Write([]byte(");"))
	p.Newline()
}
