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
			// The two companion flags mean opposite things -- one supplies a
			// precondition, the other waives it -- and neither is guessable
			// from a bare bool beside a value. This is the file a caller reads.
			if c := patchFieldDoc(p); c != "" {
				m.Body = append(m.Body, ast.Comment(c))
			}
			m.Body = append(m.Body, f)

			if name := graph.PatchFlagName(p); name != "" {
				m.Body = append(m.Body, ast.Comment(patchFlagDoc(p)))
				m.Body = append(m.Body, ast.MessageField{
					Type:   "bool",
					Name:   name,
					Number: int(graph.PatchFlagNumber(p)),
				})
			}
		}
	})
}

// patchFieldDoc documents a value slot whose meaning is not "write this".
func patchFieldDoc(p graph.Prop) string {
	if graph.PatchFlagOf(p) != graph.PatchFlagForce {
		return ""
	}
	return "The version this update requires the stored " + p.Name() + " to be.\n" +
		"It is a precondition, not a write: the update applies only if the row\n" +
		"still holds this value, and the server stamps the new version itself.\n" +
		"Setting it together with " + graph.PatchFlagName(p) + " is an error --\n" +
		"the version is the token every client's compare-and-swap is measured\n" +
		"against, so it is not the caller's to choose."
}

// patchFlagDoc documents a companion bool. Both of them read as ordinary
// optional fields and neither behaves like one.
func patchFlagDoc(p graph.Prop) string {
	switch graph.PatchFlagOf(p) {
	case graph.PatchFlagForce:
		return "Update whatever the stored " + p.Name() + " is, with no precondition.\n" +
			"The server still stamps a new version, so other clients' tokens are\n" +
			"invalidated as usual; this declines the check for THIS update only.\n" +
			"One of " + p.Name() + " or this must be set. An omitted version is\n" +
			"refused rather than assumed, because an unset field cannot be told\n" +
			"apart from a caller who never considered locking at all."

	case graph.PatchFlagNull:
		return "Clear " + p.Name() + " instead of writing it.\n" +
			"It takes a field of its own because an unset value already means\n" +
			"\"leave it alone\", so no value could have meant NULL. It wins\n" +
			"outright: setting both this and " + p.Name() + " clears."

	default:
		return ""
	}
}
