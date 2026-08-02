package app

import (
	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
)

// Path and fullname of the patch document schema.
// See https://github.com/lesomnus/protobuf-patch.
const (
	patchImportPath = "patch/patch.proto"
	patchTypeName   = "patch.Patch"
)

func (w *fileWork) xRpcApply() ast.Rpc {
	return w.defineRpc(
		ast.Comment("Apply applies a patch document to an existing "+w.entity.Name()),
		ast.Rpc{
			Name:         "Apply",
			RequestType:  w.xMsgApply().Name,
			ResponseType: w.useEntityType(w.entity),
		},
	)
}

// xMsgApply defines the request of the Apply RPC.
//
// Unlike PatchRequest, which enumerates the props it can update, the operations
// are carried by the patch document itself, so the request only has to say
// which entity they are applied to. Note that the document addresses fields of
// the entity, not of this request.
//
// There is no counterpart to PatchRequest's `<version>_force` either: a
// document expresses optimistic locking as a `test` entry on the version field,
// and its absence is the same statement `_force` makes.
func (w *fileWork) xMsgApply() ast.Message {
	return w.defineMsg("ApplyRequest", func(m *ast.Message) {
		m.Body = []ast.MessageBody{
			ast.MessageField{
				Type:   w.xMsgRef().Name,
				Name:   "ref",
				Number: 1,
			},
			ast.MessageField{
				Type:   w.useType(patchImportPath, patchTypeName),
				Name:   "patch",
				Number: 2,
			},
		}
	})
}
