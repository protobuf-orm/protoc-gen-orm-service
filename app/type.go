package app

import (
	"fmt"
	"reflect"

	"github.com/protobuf-orm/protobuf-orm/graph"
	"github.com/protobuf-orm/protobuf-orm/ormpb"
)

func protoTypeScalar(t ormpb.Type) string {
	switch t {
	case ormpb.Type_TYPE_BOOL:
		return "bool"
	case ormpb.Type_TYPE_INT32:
		return "int32"
	case ormpb.Type_TYPE_SINT32:
		return "sint32"
	case ormpb.Type_TYPE_UINT32:
		return "uint32"
	case ormpb.Type_TYPE_INT64:
		return "int64"
	case ormpb.Type_TYPE_SINT64:
		return "sint64"
	case ormpb.Type_TYPE_UINT64:
		return "uint64"
	case ormpb.Type_TYPE_SFIXED32:
		return "sfixed32"
	case ormpb.Type_TYPE_FIXED32:
		return "fixed32"
	case ormpb.Type_TYPE_FLOAT:
		return "float"
	case ormpb.Type_TYPE_SFIXED64:
		return "sfixed64"
	case ormpb.Type_TYPE_FIXED64:
		return "fixed64"
	case ormpb.Type_TYPE_DOUBLE:
		return "double"
	case ormpb.Type_TYPE_STRING:
		return "string"
	case ormpb.Type_TYPE_BYTES:
		return "bytes"
	case ormpb.Type_TYPE_UUID:
		return "bytes"
	case ormpb.Type_TYPE_TIME:
		return "google.protobuf.Timestamp"
	}

	panic(fmt.Sprintf("must be a scalar type: %v", t.String()))
}

func protoType(t ormpb.Type, s graph.Shape) (string, string) {
	if t == ormpb.Type_TYPE_GROUP {
		panic("TODO")
	}
	if t.IsScalar() {
		return protoTypeScalar(t), ""
	}

	switch s_ := s.(type) {
	case graph.ScalarShape:
		panic("it must not be a scalar")
	case graph.MapShape:
		t, p := protoType(s_.V, s_.S)
		return fmt.Sprintf("map<%s,%s>", protoTypeScalar(s_.K), t), p
	case graph.MessageShape:
		return string(s_.FullName), s_.Filepath
	default:
		panic(fmt.Sprintf("unknown shape: %s", reflect.TypeOf(s).Name()))
	}
}

// ProtoType returns string representation of protobuf type of given field
// and the filepath to the type if the type is Enum or Message.
func ProtoType(f graph.Field) (string, string) {
	return protoType(f.Type(), f.Shape())
}
