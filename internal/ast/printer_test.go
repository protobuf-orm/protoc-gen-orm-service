package ast_test

import (
	"strings"
	"testing"

	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
	"github.com/stretchr/testify/require"
)

// Enum is a top-level definition, like Message and Service.
var _ ast.TopLevelDef = ast.Enum{}

func TestPrintTypename(t *testing.T) {
	for _, tc := range []struct {
		name string
		pkg  string
		in   string
		want string
	}{
		{"strips local package", "orm.test", "orm.test.User", "User"},
		{"keeps nested name", "orm.test", "orm.test.User.Sub", "User.Sub"},
		{"keeps foreign type", "orm.test", "google.protobuf.Empty", "google.protobuf.Empty"},
		{"exact match does not panic", "orm.test", "orm.test", "orm.test"},
		{"empty package does not mangle", "", "User", "User"},
		{"partial prefix is not a boundary", "lib", "library.User", "library.User"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			x := require.New(t)
			b := &strings.Builder{}
			p := ast.NewPrinter(b, tc.pkg)
			p.PrintTypename(tc.in)
			x.Equal(tc.want, b.String())
		})
	}
}

func TestPrinter(t *testing.T) {
	WithPrinter(func(x *require.Assertions, b *strings.Builder, p ast.Printer) {
		p.Newline()
		p.Write([]byte("a"))
		p.Indent()
		p.Newline()
		p.Newline()
		p.Write([]byte("b"))
		p.Newline()
		p.Indent()
		p.Write([]byte("c"))
		p.Newline()
		p.Dedent()
		p.Write([]byte("d"))

		x.Equal(`
a

	b
		c
	d`, b.String())
	})
}

func WithPrinter(f func(x *require.Assertions, b *strings.Builder, p ast.Printer)) func(t *testing.T) {
	return func(t *testing.T) {
		x := require.New(t)
		b := &strings.Builder{}
		p := ast.NewPrinter(b, "orm.test")
		f(x, b, p)
	}
}

func PrintedTo(pt ast.PrinterTo, expected string) func(t *testing.T) {
	return WithPrinter(func(x *require.Assertions, b *strings.Builder, p ast.Printer) {
		pt.PrintTo(p)
		v := b.String()
		x.Equal(expected, v)
	})
}
