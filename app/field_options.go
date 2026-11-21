package app

import (
	"fmt"

	"github.com/protobuf-orm/protobuf-orm/graph"
	"github.com/protobuf-orm/protobuf-orm/ormpb"
	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
	"google.golang.org/protobuf/proto"
)

// A request field has the entity field's semantic type, but none of its
// storage policy. Copying the whole option would incorrectly make a request
// own a key, default or version; dropping it loses the difference between a
// UUID and arbitrary bytes. Reading the declaration also preserves an explicit
// TYPE_BYTES without annotating every scalar whose protobuf type is sufficient.
func (w *fileWork) fieldTypeOptions(f graph.Field) []ast.FieldOption {
	opts, _ := proto.GetExtension(f.Descriptor().Options(), ormpb.E_Field).(*ormpb.FieldOptions)
	if opts.GetType() == ormpb.Type_TYPE_UNSPECIFIED {
		return nil
	}
	w.imports["orm.field"] = "orm.proto"
	return []ast.FieldOption{{Name: "orm.field", Value: ast.Unsafe(fmt.Sprintf("{type: %s}", opts.GetType()))}}
}
