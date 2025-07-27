package app

import (
	"context"
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/protobuf-orm/protobuf-orm/graph"
	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
	"google.golang.org/protobuf/compiler/protogen"
)

type work struct {
	// entity fullname -> generated filepath
	paths map[string]string

	// type -> path
	// e.g. "google.protobuf.Timestamp" -> "google/protobuf/timestamp.proto"
	imports map[string]string
	msgs    map[string]ast.Message
}

func newWork() *work {
	return &work{
		paths: map[string]string{},

		imports: map[string]string{},
		msgs:    map[string]ast.Message{},
	}
}

type fileWork struct {
	root   *work
	entity graph.Entity

	imports map[string]string
	msgs    map[string]ast.Message

	names []string
	rpcs  []ast.ServiceBody
}

func (w *work) newFileWork(entity graph.Entity) *fileWork {
	fw := &fileWork{
		root:   w,
		entity: entity,

		imports: map[string]string{},
		msgs:    map[string]ast.Message{},

		names: []string{},
		rpcs:  []ast.ServiceBody{},
	}

	w.imports[string(entity.FullName())] = entity.Path()
	w.msgs[string(entity.FullName())] = ast.Message{}

	return fw
}

func (w *work) mustGetPath(v graph.Entity) string {
	name := string(v.FullName())
	p, ok := w.paths[name]
	if !ok {
		panic(fmt.Sprintf("generated filepath to the entity not found: %s", name))
	}

	return p
}

func (w *fileWork) path() string {
	return w.root.mustGetPath(w.entity)
}

// withEntity references message which defined by the given entity.
func (w *fileWork) withEntity(v graph.Entity) *fileWork {
	name := string(v.FullName())
	p := w.root.mustGetPath(v)
	w.imports["_"+name] = p
	return w.root.newFileWork(v)
}

func (w *fileWork) useEntityType(v graph.Entity) string {
	t := string(v.FullName())
	w.root.imports[t] = v.Path()
	w.imports[t] = v.Path()

	return t
}

func (w *fileWork) useFieldType(v graph.Field) string {
	t, p := ProtoType(v)
	if len(p) > 0 {
		w.root.imports[t] = p
		w.imports[t] = p
	}
	return t
}

func (w *fileWork) useType(p string, fullname string) string {
	w.root.imports[fullname] = p
	w.imports[fullname] = p
	return fullname
}

func (w *fileWork) defineMsg(name string, f func(m *ast.Message)) ast.Message {
	name = nameMsg(w.entity, name)
	m, ok := w.msgs[name]
	if ok {
		return m
	}

	m = ast.Message{
		Name: name,
		Body: []ast.MessageBody{},
	}

	f(&m)
	if m.Name != name {
		panic("do not alter tht name")
	}

	w.root.imports[name] = w.path()
	w.root.msgs[name] = m
	w.msgs[name] = m
	w.names = append(w.names, name)

	return m
}

func (w *fileWork) defineRpc(v ast.Rpc) {
	if _, ok := w.root.imports[v.RequestType]; !ok {
		panic(fmt.Sprintf("request type not found: %s", v.RequestType))
	}
	if _, ok := w.root.imports[v.ResponseType]; !ok {
		panic(fmt.Sprintf("response type not found: %s", v.ResponseType))
	}

	w.rpcs = append(w.rpcs, v)
}

func (w *work) run(ctx context.Context, gf *protogen.GeneratedFile, entity graph.Entity) error {
	fw := w.newFileWork(entity)

	rpcs := entity.Rpcs()
	if rpcs.HasAdd() {
		fw.xRpcAdd()
	}
	if rpcs.HasGet() {
		fw.xRpcGet()
	}
	if rpcs.HasErase() {
		fw.xRpcErase()
	}

	f := ast.File{
		Edition: ast.Edition2023,
		Package: string(entity.Package()),
		Defs: []ast.TopLevelDef{
			ast.Service{
				Name: string(entity.FullName().Name()) + "Service",
				Body: fw.rpcs,
			},
		},
	}

	imports := slices.Collect(maps.Values(fw.imports))
	slices.SortFunc(imports, strings.Compare)
	imports = slices.Compact(imports)
	for _, v := range imports {
		f.Imports = append(f.Imports, ast.Import{
			Name: v,
		})
	}

	for _, name := range fw.names {
		v, ok := fw.msgs[name]
		if !ok {
			panic(fmt.Sprintf("message not found: %s", name))
		}
		f.Defs = append(f.Defs, v)
	}

	p := ast.NewPrinter(gf, f.Package)
	f.PrintTo(p)

	return nil
}

func nameMsg(v graph.Entity, name string) string {
	pkg := v.FullName().Parent()
	return string(pkg.Append(v.FullName().Name())) + name
}
