package app

import (
	"github.com/protobuf-orm/protobuf-orm/graph"
	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
)

func (w *fileWork) xRpcAdd() ast.Rpc {
	return w.defineRpc(
		ast.Comment("Add creates a new "+w.entity.Name()),
		ast.Rpc{
			Name:         "Add",
			RequestType:  w.xMsgAddRequest().Name,
			ResponseType: w.useEntityType(w.entity),
		},
	)
}

func (w *fileWork) xMsgAddRequest() ast.Message {
	return w.defineMsg("AddRequest", func(m *ast.Message) {
		for p := range w.entity.Props() {
			f := ast.MessageField{
				Name:   p.Name(),
				Number: int(p.Number()),
			}
			if p.IsList() {
				f.Label = ast.LabelRepeated
			}

			switch p := p.(type) {
			case graph.Field:
				if p.IsVersion() {
					continue
				}

				f.Type = w.useFieldType(p)
				if !p.IsList() && p.Type().IsScalar() && !p.IsOptional() {
					f.Opts = append(f.Opts, ast.FeaturesFieldPresenceImplicit.WithinField())
				}

			case graph.Edge:
				f.Type = w.withEntity(p.Target()).xMsgRef().Name

			default:
				panic(errUnknownPropType)
			}
			m.Body = append(m.Body, f)
		}
	})
}
