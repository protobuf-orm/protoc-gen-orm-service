package ast_test

import (
	"testing"

	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
)

func TestFile(t *testing.T) {
	t.Run("valid", PrintedTo(ast.File{
		Edition: ast.Edition2023,
		Package: "orm.example",
		Imports: []ast.Import{
			{Name: "google/protobuf/timestamp.proto"},
			{Name: "orm.proto"},
		},
		Options: []ast.Option{
			{
				Known: true,
				Name:  "go_package",
				Value: ast.String("github.com/protobuf-orm"),
			},
		},
		Defs: []ast.TopLevelDef{
			ast.Service{
				Name: "UserService",
				Body: []ast.ServiceBody{
					ast.Rpc{Name: "Add", RequestType: "UserAddRequest", ResponseType: "User"},
					ast.Rpc{Name: "Get", RequestType: "UserRef", ResponseType: "User"},
				},
			},
			ast.Message{
				Name: "UserRef",
				Body: []ast.MessageBody{
					ast.MessageOneof{
						Name: "key",
						Body: []ast.MessageOneofBody{
							ast.MessageOneofField{Type: "bytes", Name: "id", Number: 1},
							ast.MessageOneofField{Type: "string", Name: "alias", Number: 2},
						},
					},
				},
			},
		},
	}, `edition = "2023";

package orm.example;

import "google/protobuf/timestamp.proto";
import "orm.proto";

option go_package = "github.com/protobuf-orm";

service UserService {
	rpc Add(UserAddRequest) returns (User);
	rpc Get(UserRef) returns (User);
}

message UserRef {
	oneof key {
		bytes id = 1;
		string alias = 2;
	}
}
`))
}
