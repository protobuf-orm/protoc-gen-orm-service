package apptest_test

import (
	"context"
	"maps"
	"slices"
	"strings"
	"testing"
	"text/template"

	"github.com/protobuf-orm/protobuf-orm/graph"
	"github.com/protobuf-orm/protoc-gen-orm-service/app"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"

	// Registers the fixtures the generator runs on.
	_ "github.com/protobuf-orm/protoc-gen-orm-service/internal/apptest"
)

// The fixtures under proto/apptest, named explicitly rather than swept out of
// the registry: this package is generated into an ignored directory, so a
// fixture that was removed can leave its .pb.go behind, and a sweep would feed
// the generator a source that no longer exists.
var fixtures = []string{
	"apptest/edge.proto",
	"apptest/field.proto",
	"apptest/tenant.proto",
	"apptest/user.proto",
}

// generate runs the plugin over the fixtures the way `buf generate` does, and
// returns the generated files by path. Descriptors come from the global
// registry, so this needs no toolchain beyond the Go one.
func generate(t *testing.T) map[string]string {
	t.Helper()
	x := require.New(t)

	// protogen requires a file to appear after everything it imports.
	seen := map[string]bool{}
	fds := []*descriptorpb.FileDescriptorProto{}
	var collect func(fd protoreflect.FileDescriptor)
	collect = func(fd protoreflect.FileDescriptor) {
		if seen[fd.Path()] {
			return
		}
		seen[fd.Path()] = true

		imports := fd.Imports()
		for i := range imports.Len() {
			collect(imports.Get(i).FileDescriptor)
		}
		fds = append(fds, protodesc.ToFileDescriptorProto(fd))
	}
	for _, path := range fixtures {
		fd, err := protoregistry.GlobalFiles.FindFileByPath(path)
		x.NoError(err, "run `buf generate` to build the fixtures")
		collect(fd)
	}

	p, err := protogen.Options{}.New(&pluginpb.CodeGeneratorRequest{
		FileToGenerate: fixtures,
		Parameter:      proto.String("namer={{.Name}}_svc.g.proto"),
		ProtoFile:      fds,
	})
	x.NoError(err)
	p.SupportedEditionsMinimum = descriptorpb.Edition_EDITION_PROTO2
	p.SupportedEditionsMaximum = descriptorpb.Edition_EDITION_MAX

	ctx := context.Background()
	g := graph.NewGraph()
	x.NoError(graph.ParseFiles(ctx, g, p.Files))

	a, err := app.New(app.WithNamer(template.Must(
		template.New("namer").Parse("{{.Name}}_svc.g.proto"),
	)))
	x.NoError(err)
	x.NoError(a.Run(ctx, p, g))

	res := p.Response()
	x.Nil(res.Error, res.GetError())

	files := map[string]string{}
	for _, f := range res.File {
		files[f.GetName()] = f.GetContent()
	}

	return files
}

func TestRpcApply(t *testing.T) {
	x := require.New(t)

	files := generate(t)
	src, ok := files["apptest/user_svc.g.proto"]
	x.True(ok, "generated: %v", slices.Sorted(maps.Keys(files)))

	x.Contains(src, `import "patch/patch.proto";`)
	x.Contains(src, "\trpc Apply(UserApplyRequest) returns (User);\n")

	// The document carries the operations, so the request only says which
	// entity they apply to.
	x.Contains(src, strings.Join([]string{
		"message UserApplyRequest {",
		"\tUserRef ref = 1;",
		"\tpatch.Patch patch = 2;",
		"}",
	}, "\n"))
}

// Apply is emitted exactly where Patch is: protobuf-orm has no `apply` option,
// so the two share the `patch` one.
func TestRpcApplyIsGatedOnPatch(t *testing.T) {
	x := require.New(t)

	for path, src := range generate(t) {
		x.Equal(
			strings.Contains(src, "rpc Patch("),
			strings.Contains(src, "rpc Apply("),
			"Patch and Apply disagree in %s", path,
		)
	}
}
