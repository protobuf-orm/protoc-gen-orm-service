package app

import (
	"slices"

	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
)

func (w *fileWork) xMsgRef() ast.Message {
	pkg := w.entity.FullName().Parent()
	name := string(pkg.Append(w.entity.FullName().Name() + "Ref"))
	if v, ok := w.msgs[name]; ok {
		return v
	}

	fs := []ast.MessageOneofField{}
	for v := range w.entity.Fields() {
		if !v.IsUnique() {
			continue
		}

		t := w.useFieldType(v)
		fs = append(fs, ast.MessageOneofField{
			Type:   t,
			Name:   string(v.FullName().Name()),
			Number: int(v.Number()),
		})
	}
	// for v := range w.entity.Indexes() {
	// 	if !v.IsUnique() {
	// 		continue
	// 	}

	// }

	u := make([]ast.MessageOneofBody, 0, len(fs))
	slices.SortFunc(fs, func(a, b ast.MessageOneofField) int {
		return a.Number - b.Number
	})
	for _, f := range fs {
		u = append(u, f)
	}

	v := ast.Message{
		Name: name,
		Body: []ast.MessageBody{
			ast.MessageOneof{
				Name: "kind",
				Body: u,
			},
		},
	}

	w.defineMsg(v)
	return v
}
