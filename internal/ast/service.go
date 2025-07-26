package ast

import "fmt"

type Service struct {
	Name string
	Body []ServiceBody

	tagTopLevelDef
}

type ServiceBody interface {
	Node
	_serviceBody()
}

type tagServiceBody struct{}

func (tagServiceBody) _serviceBody() {}

func (v Service) PrintTo(p Printer) {
	fmt.Fprintf(p, "service %s {", v.Name)
	ScopedPrint(p, v.Body, "}")
}
