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
				// The version is the server's to stamp, so there is nothing for
				// a request to say about it.
				if p.IsVersion() {
					continue
				}
				// And a row is added alive. No value the erased field could
				// carry here means anything: a date on it says the row is gone,
				// and a row that arrives gone is a delete with extra steps --
				// one that skips whatever Erase does besides writing a column.
				// `graph.PatchProps` leaves it out of the PatchRequest for the
				// same reason.
				if p.IsErased() {
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
