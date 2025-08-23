package app

import (
	"slices"

	"github.com/ettle/strcase"
	"github.com/protobuf-orm/protobuf-orm/graph"
	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
)

func (w *fileWork) xMsgRef() ast.Message {
	return w.defineMsg("Ref", func(m *ast.Message) {
		fs := []ast.MessageOneofField{}
		for p := range w.entity.Keys() {
			f := ast.MessageOneofField{
				Name:   p.Name(),
				Number: int(p.Number()),
			}

			switch p := p.(type) {
			case graph.Field:
				f.Type = w.useFieldType(p)

			case graph.Edge:
				panic("edge key not implemented")

			case graph.Index:
				f.Type = w.xMsgRefByIndex(p).Name

			default:
				panic(errUnknownPropType)
			}
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
	return w.defineMsg("RefBy"+strcase.ToPascal(index.Name()), func(m *ast.Message) {
		for p := range index.Props() {
			f := ast.MessageField{
				Name:   p.Name(),
				Number: int(p.Number()),
			}

			switch u := p.(type) {
			case graph.Field:
				f.Type = w.useFieldType(u)

			case graph.Edge:
				f.Type = w.withEntity(u.Target()).xMsgRef().Name

			default:
				panic(errUnknownPropType)
			}
			m.Body = append(m.Body, f)
		}
	})
}
