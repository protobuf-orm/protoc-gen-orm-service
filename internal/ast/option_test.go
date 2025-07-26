package ast_test

import (
	"strings"
	"testing"

	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
	"github.com/stretchr/testify/require"
)

func TestOption(t *testing.T) {
	t.Run("protobuf defined option", WithPrinter(func(x *require.Assertions, b *strings.Builder, p ast.Printer) {
		v := ast.Option{
			Name:  "go_package",
			Value: ast.String("foo"),
		}
		v.PrintTo(p)

		x.Equal(b.String(), `option go_package = "foo";
`)
	}))
	t.Run("custom option", WithPrinter(func(x *require.Assertions, b *strings.Builder, p ast.Printer) {
		v := ast.Option{
			Name:  "orm.message",
			Value: ast.String("foo"),
		}
		v.PrintTo(p)

		x.Equal(b.String(), `option (orm.message) = "foo";
`)
	}))
}
