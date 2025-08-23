package app

import (
	"fmt"

	"github.com/protobuf-orm/protobuf-orm/graph"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// ProtoType returns string representation of protobuf type of given field
// and the filepath to the type if the type is Enum or Message.
func ProtoType(f graph.Field) (string, string) {
	return protoType(f.Descriptor())
}

func protoType(d protoreflect.FieldDescriptor) (string, string) {
	kind := d.Kind()
	switch kind {
	case protoreflect.EnumKind:
		return string(d.Enum().FullName()), d.ParentFile().Path()
	case protoreflect.MessageKind:
		if d.IsMap() {
			k := d.MapKey().Kind().String()
			v, p := protoType(d.MapValue())
			return fmt.Sprintf("map<%s, %s>", k, v), p
		}

		d := d.Message()
		return string(d.FullName()), d.ParentFile().Path()
	default:
		return kind.String(), ""
	}
}
