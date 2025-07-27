package app

import (
	"github.com/protobuf-orm/protobuf-orm/graph"
	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
)

func (w *fileWork) xRpcPatch() ast.Rpc {
	req := w.xMsgPatch()
	v := ast.Rpc{
		Name:         "Patch",
		RequestType:  req.Name,
		ResponseType: w.useEntityType(w.entity),
	}

	w.defineRpc(v)
	return v
}

func (w *fileWork) xMsgPatch() ast.Message {
	return w.defineMsg("PatchRequest", func(m *ast.Message) {
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
			m.Body = append(m.Body, f)
		}
	})
}
