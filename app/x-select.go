package app

import (
	"github.com/protobuf-orm/protobuf-orm/graph"
	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
)

func (w *fileWork) xMsgSelect() ast.Message {
	return w.defineMsg("Select", func(m *ast.Message) {
		for p := range w.entity.Props() {
			t := "bool"
			if f, ok := p.(graph.Edge); ok {
				t = w.withEntity(f.Target()).xMsgSelect().Name
			}

			m.Body = append(m.Body, ast.MessageField{
				Type:   t,
				Name:   p.Name(),
				Number: int(p.Number()),
			})
		}

		k := w.entity.Key()
		for i, f := range m.Body {
			f := f.(ast.MessageField)
			if f.Number != int(k.Number()) {
				continue
			}

			f.Name = "all"
			m.Body[i] = f
			break
		}
	})
}
