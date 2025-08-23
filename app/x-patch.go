package app

import (
	"github.com/protobuf-orm/protobuf-orm/graph"
	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
)

func (w *fileWork) xRpcPatch() ast.Rpc {
	return w.defineRpc(
		ast.Comment("Patch updates an existing "+w.entity.Name()),
		ast.Rpc{
			Name:         "Patch",
			RequestType:  w.xMsgPatch().Name,
			ResponseType: w.useEntityType(w.entity),
		},
	)
}

func (w *fileWork) xMsgPatch() ast.Message {
	return w.defineMsg("PatchRequest", func(m *ast.Message) {
		m.Body = append(m.Body, ast.MessageField{
			Type:   w.xMsgRef().Name,
			Name:   "target",
			Number: int(w.entity.Key().Number()),
		})

		for p := range w.entity.Props() {
			if p.IsImmutable() {
				continue
			}

			f := ast.MessageField{
				Name:   p.Name(),
				Number: int(p.Number())*2 - 1,
			}
			if p.Descriptor().IsList() {
				f.Label = ast.LabelRepeated
			}

			switch p := p.(type) {
			case graph.Field:
				f.Type = w.useFieldType(p)

			case graph.Edge:
				f.Type = w.withEntity(p.Target()).xMsgRef().Name

			default:
				panic(errUnknownPropType)
			}
			m.Body = append(m.Body, f)

			if p.IsNullable() {
				m.Body = append(m.Body, ast.MessageField{
					Type:   "bool",
					Name:   p.Name() + "_null",
					Number: int(p.Number()) * 2,
				})
			}
		}
	})
}
