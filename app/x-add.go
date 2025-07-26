package app

import (
	"github.com/protobuf-orm/protobuf-orm/graph"
	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
)

func (w *fileWork) xRpcAdd() ast.Rpc {
	req := w.xMsgAddRequest()
	v := ast.Rpc{
		Name:         "Add",
		RequestType:  req.Name,
		ResponseType: string(w.entity.FullName()),
	}

	w.defineRpc(v)
	return v
}

func (w *fileWork) xMsgAddRequest() ast.Message {
	pkg := w.entity.FullName().Parent()
	name := string(pkg.Append(w.entity.FullName().Name() + "AddRequest"))
	if v, ok := w.msgs[name]; ok {
		return v
	}

	v := ast.Message{
		Name: name,
		Body: []ast.MessageBody{},
	}
	for p := range w.entity.Props() {
		f := ast.MessageField{
			Name:   string(p.FullName().Name()),
			Number: int(p.Number()),
		}
		switch u := p.(type) {
		case graph.Field:
			t := w.useFieldType(u)
			f.Type = t
		case graph.Edge:
			t := w.withEntity(u.Target()).xMsgRef()
			f.Type = t.Name

		default:
			panic("unknown type of graph prop")
		}
		v.Body = append(v.Body, f)
	}

	w.defineMsg(v)
	return v
}
