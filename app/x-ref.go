package app

import (
	"slices"

	"github.com/iancoleman/strcase"
	"github.com/protobuf-orm/protobuf-orm/graph"
	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
)

func (w *fileWork) xMsgRef() ast.Message {
	return w.defineMsg("Ref", func(m *ast.Message) {
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
		for v := range w.entity.Indexes() {
			if !v.IsUnique() {
				continue
			}

			f := ast.MessageOneofField{
				Type: w.xMsgRefByIndex(v).Name,
				Name: v.Name(),
			}
			v.Props()(func(v graph.Prop) bool {
				f.Number = int(v.Number())
				return false
			})

			fs = append(fs, f)
		}

		u := make([]ast.MessageOneofBody, 0, len(fs))
		slices.SortFunc(fs, func(a, b ast.MessageOneofField) int {
			return a.Number - b.Number
		})
		for _, f := range fs {
			u = append(u, f)
		}

		m.Body = append(m.Body, ast.MessageOneof{
			Name: "key",
			Body: u,
		})
	})
}

func (w *fileWork) xMsgRefByIndex(index graph.Index) ast.Message {
	return w.defineMsg("RefBy"+strcase.ToCamel(index.Name()), func(m *ast.Message) {
		for p := range index.Props() {
			f := ast.MessageField{
				Name:   string(p.FullName().Name()),
				Number: int(p.Number()),
			}
			switch u := p.(type) {
			case graph.Field:
				t := w.useFieldType(u)
				f.Type = t
			case graph.Edge:
				t := w.withEntity(u.Target()).xMsgRef()
				f.Type = t.Name

			default:
				panic("unknown type of graph prop")
			}
			m.Body = append(m.Body, f)
		}
	})
}
