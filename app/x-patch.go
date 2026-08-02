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

// xMsgPatch emits the PatchRequest.
//
// Where each field goes is graph's to say, not this generator's: the server
// that converts a request back into a patch document has to read the same
// layout, and it lives in another repository. Every number and every companion
// name here comes from graph.Patch*, so the two cannot drift apart.
func (w *fileWork) xMsgPatch() ast.Message {
	return w.defineMsg("PatchRequest", func(m *ast.Message) {
		m.Body = append(m.Body, ast.MessageField{
			Type:   w.xMsgRef().Name,
			Name:   "ref",
			Number: int(graph.PatchRefNumber(w.entity)),
		})

		for p := range graph.PatchProps(w.entity) {
			f := ast.MessageField{
				Name:   p.Name(),
				Number: int(graph.PatchValueNumber(p)),
			}
			if p.IsList() {
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

			if name := graph.PatchFlagName(p); name != "" {
				m.Body = append(m.Body, ast.MessageField{
					Type:   "bool",
					Name:   name,
					Number: int(graph.PatchFlagNumber(p)),
				})
			}
		}
	})
}
