package app

import (
	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
)

func (w *fileWork) xRpcErase() ast.Rpc {
	req := w.xMsgRef()
	v := ast.Rpc{
		Name:         "Erase",
		RequestType:  req.Name,
		ResponseType: w.useType("google/protobuf/empty.proto", "google.protobuf.Empty"),
	}

	w.defineRpc(v, ast.Comment("Erase deletes a "+w.entity.Name()))
	return v
}
