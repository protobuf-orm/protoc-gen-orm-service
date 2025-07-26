package ast_test

import (
	"testing"

	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
)

func TestComment(t *testing.T) {
	t.Run("single line", PrintedTo(
		ast.Comment("foo"),
		`// foo
`,
	))
	t.Run("multiline", PrintedTo(
		ast.Comment("foo\nbar"),
		`// foo
// bar
`,
	))
}

func TestMultilineComment(t *testing.T) {
	t.Run("single line", PrintedTo(
		ast.MultilineComment("foo"),
		`/**
 * foo
 */
`,
	))
	t.Run("multiline", PrintedTo(
		ast.MultilineComment("foo\nbar"),
		`/**
 * foo
 * bar
 */
`,
	))

}
