package ast_test

import (
	"testing"

	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
)

func TestMessage(t *testing.T) {
	t.Run("fields", PrintedTo(ast.Message{
		Name: "Foo",
		Body: []ast.MessageBody{
			ast.MessageField{Type: "bytes", Name: "id", Number: 0},
			ast.MessageField{Label: "repeated", Type: "int32", Name: "values", Number: 1},
		},
	},
		`message Foo {
	bytes id = 0;
	repeated int32 values = 1;
}
`))
	t.Run("options", PrintedTo(ast.Message{
		Name: "Foo",
		Body: []ast.MessageBody{
			ast.MessageField{Type: "bytes", Name: "id", Number: 0},
			ast.MessageField{
				Type:   "string",
				Name:   "name",
				Number: 2,
				Opts: []ast.FieldOption{
					ast.FeaturesFieldPresenceImplicit,
					{
						Name:  "my_option",
						Value: ast.String("Django"),
					},
				},
			},
			ast.MessageField{Label: "repeated", Type: "int32", Name: "values", Number: 1},
		},
	},
		`message Foo {
	bytes id = 0;
	string name = 2 [
		features.field_presence = IMPLICIT
		, (my_option) = "Django"
	];
	repeated int32 values = 1;
}
`))
	t.Run("oneof field", PrintedTo(ast.Message{
		Name: "Foo",
		Body: []ast.MessageBody{
			ast.MessageField{Type: "bytes", Name: "id", Number: 0},
			ast.MessageOneof{
				Name: "kind",
				Body: []ast.MessageOneofBody{
					ast.MessageOneofField{Type: "bool", Name: "a", Number: 2},
					ast.MessageOneofField{Type: "string", Name: "b", Number: 3},
				},
			},
			ast.MessageField{Label: "repeated", Type: "int32", Name: "values", Number: 1},
		},
	},
		`message Foo {
	bytes id = 0;
	oneof kind {
		bool a = 2;
		string b = 3;
	}
	repeated int32 values = 1;
}
`))
}
