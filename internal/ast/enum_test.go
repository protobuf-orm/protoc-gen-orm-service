package ast_test

import (
	"testing"

	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
)

func TestEnum(t *testing.T) {
	t.Run("fields", PrintedTo(
		ast.Enum{
			Name: "Level",
			Body: []ast.EnumBody{
				ast.EnumField{Name: "Low", Number: 0},
				ast.EnumField{Name: "Mid", Number: 1},
				ast.EnumField{Name: "High", Number: 2},
			},
		},
		`enum Level {
	Low = 0;
	Mid = 1;
	High = 2;
}
`,
	))
	t.Run("field with option", PrintedTo(
		ast.Enum{
			Name: "Level",
			Body: []ast.EnumBody{
				ast.EnumField{Name: "Low", Number: 0},
				ast.EnumField{Name: "Mid", Number: 1, Options: []ast.FieldOption{
					{Name: "foo.bar", Value: ast.String("baz")},
				}},
				ast.EnumField{Name: "High", Number: 2},
			},
		},
		`enum Level {
	Low = 0;
	Mid = 1 [
		(foo.bar) = "baz"
	];
	High = 2;
}
`,
	))
	t.Run("with option", PrintedTo(
		ast.Enum{
			Name: "Level",
			Body: []ast.EnumBody{
				ast.Option{Name: "foo", Value: ast.String("bar")},
				ast.EnumField{Name: "Low", Number: 0},
				ast.Option{Known: true, Name: "bar", Value: ast.String("baz")},
				ast.EnumField{Name: "Mid", Number: 1},
				ast.Option{Name: "baz", Value: ast.String("qux")},
			},
		},
		`enum Level {
	option (foo) = "bar";
	Low = 0;
	option bar = "baz";
	Mid = 1;
	option (baz) = "qux";
}
`,
	))
}
