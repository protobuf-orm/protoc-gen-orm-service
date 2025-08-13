package app

import (
	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
)

func (w *fileWork) xRpcGet() ast.Rpc {
	req := w.xMsgGet()
	v := ast.Rpc{
		Name:         "Get",
		RequestType:  req.Name,
		ResponseType: w.useEntityType(w.entity),
	}

	w.defineRpc(v, ast.Comment("Get retrieves a "+w.entity.Name()))
	return v
}

func (w *fileWork) xMsgGet() ast.Message {
	return w.defineMsg("GetRequest", func(m *ast.Message) {
		m.Body = []ast.MessageBody{
			ast.MessageField{
				Type:   w.xMsgRef().Name,
				Name:   "ref",
				Number: 1,
			},
			ast.MessageField{
				Type:   w.xMsgSelect().Name,
				Name:   "select",
				Number: 2,
			},
		}
	})
}
