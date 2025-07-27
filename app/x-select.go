package app

import (
	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
)

func (w *fileWork) xMsgSelect() ast.Message {
	name := nameMsg(w.entity, "Select")
	if v, ok := w.msgs[name]; ok {
		return v
	}

	v := ast.Message{
		Name: name,
		Body: []ast.MessageBody{},
	}
	for p := range w.entity.Props() {
		v.Body = append(v.Body, ast.MessageField{
			Type:   "bool",
			Name:   string(p.FullName().Name()),
			Number: int(p.Number()),
		})
	}

	k := w.entity.Key()
	for i, f_ := range v.Body {
		f := f_.(ast.MessageField)
		if f.Number != int(k.Number()) {
			continue
		}

		f.Name = "all"
		v.Body[i] = f
		break
	}

	w.defineMsg(v)
	return v
}
