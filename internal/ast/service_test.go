package ast_test

import (
	"testing"

	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
)

func TestService(t *testing.T) {
	t.Run("rpcs", PrintedTo(ast.Service{
		Name: "Foo",
		Body: []ast.ServiceBody{
			ast.Rpc{
				Name:         "Unary",
				RequestType:  "Bar",
				ResponseType: "Baz",
			},
			ast.Rpc{
				Name:          "ClientStream",
				RequestStream: true,
				RequestType:   "Bar",
				ResponseType:  "Baz",
			},
			ast.Rpc{
				Name:           "ServerStream",
				RequestType:    "Bar",
				ResponseStream: true,
				ResponseType:   "Baz",
			},
			ast.Rpc{
				Name:           "BidiStream",
				RequestStream:  true,
				RequestType:    "Bar",
				ResponseStream: true,
				ResponseType:   "Baz",
			},
		},
	},
		`service Foo {
	rpc Unary(Bar) returns (Baz);
	rpc ClientStream(stream Bar) returns (Baz);
	rpc ServerStream(Bar) returns (stream Baz);
	rpc BidiStream(stream Bar) returns (stream Baz);
}
`))
}
