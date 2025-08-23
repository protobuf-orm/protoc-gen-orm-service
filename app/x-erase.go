package app

import (
	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
)

func (w *fileWork) xRpcErase() ast.Rpc {
	return w.defineRpc(
		ast.Comment("Erase deletes a "+w.entity.Name()),
		ast.Rpc{
			Name:         "Erase",
			RequestType:  w.xMsgRef().Name,
			ResponseType: w.useType("google/protobuf/empty.proto", "google.protobuf.Empty"),
		},
	)
}
