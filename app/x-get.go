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

	w.defineRpc(v)
	return v
}

func (w *fileWork) xMsgGet() ast.Message {
	name := nameMsg(w.entity, "GetRequest")
	if v, ok := w.msgs[name]; ok {
		return v
	}

	v := ast.Message{
		Name: name,
		Body: []ast.MessageBody{
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
		},
	}

	w.defineMsg(v)
	return v
}
