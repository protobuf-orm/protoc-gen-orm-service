# protoc-gen-orm-service

A `protoc`/`buf` plugin that generates **gRPC service definitions** (`.proto`)
from messages annotated with [protobuf-orm](https://github.com/protobuf-orm/protobuf-orm)
options.

You annotate an entity once; this plugin emits its service with the verbs it
names -- `Add` / `AddMany` / `Get` / `Patch` / `Apply` / `Erase` -- and the
request/response messages those RPCs need. The output is
itself a `.proto` file, so it feeds back into your normal protoc pipeline (e.g.
`protoc-gen-go-grpc`, or [protoc-gen-orm-go](../protoc-gen-orm-go) /
[protoc-gen-orm-ent](../protoc-gen-orm-ent) for the implementation).

## What it generates

Given an ORM entity:

```proto
message User {
  bytes id = 1 [(orm.field) = {type: TYPE_UUID, key: true}];
  Tenant tenant = 2;
  string alias = 4 [(orm.field) = {default: ""}];
  string name = 5 [(orm.field) = {default: ""}];
  map<string, string> labels = 7;
  google.protobuf.Timestamp date_updated = 14 [(orm.field) = {version: {}}];
  google.protobuf.Timestamp date_created = 15 [(orm.field) = {immutable: true, default: ""}];

  option (orm.message) = {
    rpc: {add: {}, get: {}, patch: {}, apply: {}, erase: {}}
    indexes: [{name: "alias", unique: true, refs: [{name: "alias", number: 4}, {name: "tenant", number: 2}]}]
  };
}
```

it produces `user_svc.g.proto`:

```proto
service UserService {
  rpc Add(UserAddRequest) returns (User);
  rpc Get(UserGetRequest) returns (User);
  rpc Patch(UserPatchRequest) returns (User);
  rpc Apply(UserApplyRequest) returns (User);
  rpc Erase(UserRef) returns (UserEraseResponse);
}

message UserAddRequest { ... }        // creatable fields (no key, no version)
message UserRefByAlias { ... }        // one message per unique index
message UserRef { oneof key { ... } } // key OR any unique index
message UserSelect { ... }            // field mask (bool per prop)
message UserGetRequest { UserRef ref = 1; UserSelect select = 2; }
message UserPatchRequest { ... }      // nullable fields get a *_null flag; version gets *_force
message UserApplyRequest { UserRef ref = 1; patch.Patch patch = 2; }
```

Each verb is there when `(orm.message).rpc` names it and does not disable it,
and no verb brings another: `patch: {}` is a Patch and no Apply, and `apply: {}`
an Apply and no Patch. The one exception is AddMany, which is asked for inside
the Add. An `rpc: {}` that names no verb is a service with none.

Highlights of the mapping:

- **`UserRef`** is a `oneof` over the primary key and every **unique index**
  (each unique index also gets its own `UserRefBy<Index>` message).
- **`UserAddRequest`** contains only the props you can supply at creation time
  (the key and `version` fields are excluded).
- **`UserSelect`** is a field mask used by `Get`.
- **`UserPatchRequest`** adds a `<field>_null` boolean for nullable fields and a
  `<field>_force` boolean for the optimistic-locking `version` field.
- **`UserApplyRequest`**, where the entity asks for `apply`, carries a
  [`patch.Patch`](https://github.com/lesomnus/protobuf-patch) document instead of
  a field per prop, for the updates a fixed request shape cannot express — edit
  one map key, splice a list, or test-then-set. See [Apply](#apply).
- **`UserAddManyRequest`** is there when the entity asks for AddMany. See
  [AddMany](#addmany).

## Field types in requests

A declared `(orm.field).type` follows a field into `AddRequest`, `PatchRequest`,
scalar reference keys and composite-index messages. For example, an entity's
`bytes subject = 8 [(orm.field) = {type: TYPE_UUID, unique: true}];` produces
request and reference fields that still carry `type: TYPE_UUID`. Descriptor
consumers can recognize UUIDs without depending on an `id` naming convention.
An explicit `TYPE_BYTES` stays bytes too.

Only the type is copied. Key ownership, uniqueness, defaults, nullability and
version behavior belong to the entity; copying those storage options into a
request would give them a second meaning. The generated field's wire type,
number and presence stay the same. Regenerate the contracts and client code
after updating the plugin; no database migration is needed for this metadata.

## AddMany

`add: {many: {}}` beside an entity's Add asks for AddMany, which adds many rows
in one call; an Add alone does not:

```proto
option (orm.message) = {rpc: {add: {many: {max: 200}}, get: {}}};
```

```proto
rpc AddMany(UserAddManyRequest) returns (UserAddManyResponse);

message UserAddManyRequest { repeated UserAddRequest items = 1; }
message UserAddManyResponse { repeated User items = 1; }
```

An item is the entity's own `UserAddRequest`, so whatever reads an Add's request
reads an item the same way. The rows come back in the order of the items, and
the call is all of them or none. `max` is the most items one call may carry,
1000 when unset; the comment on the RPC says which.

## Apply

`Patch` spends one request field per prop, so it can only say "set this prop to
this value". `Apply` takes a
[`patch.Patch`](https://github.com/lesomnus/protobuf-patch) document instead —
the operations travel inside it, so the request only has to name the entity:

```proto
message UserApplyRequest {
  UserRef     ref = 1;
  patch.Patch patch = 2;
}
```

The document addresses fields of `User`, not of the request. That buys the edits
`UserPatchRequest` has no shape for: removing a single map key rather than
replacing `labels` wholesale, inserting into a list at an index, or asserting a
current value before writing.

There is no `<field>_force` counterpart either. A document states optimistic
locking as a `test` entry on the version field, and omitting that entry is the
same statement `_force` makes. Apply is atomic in the same way the format is: if
any entry fails — a `test` that does not hold, a field that has moved — nothing
is written.

`Apply` is there where the entity asks for it with `apply: {}`, and not because
there is a `Patch`. Either is there without the other: an entity with `apply`
and no `patch` has an `ApplyRequest` and no `PatchRequest`.

```proto
option (orm.message) = {rpc: {get: {}, apply: {}}};
```

Where there is an Apply, the generated `.proto` imports `patch/patch.proto`, so
add the module to your `buf.yaml`:

```yaml
deps:
  - buf.build/orm/orm
  - buf.build/patch/patch
```

> The patch schema is **not frozen**. It is changed in place, so pin it with
> `buf.lock` and re-read its
> [format.md](https://github.com/lesomnus/protobuf-patch/blob/main/docs/format.md)
> when you move.

## Usage

Add the plugin to your `buf.gen.yaml`:

```yaml
version: v2
plugins:
  - local: [go, run, github.com/protobuf-orm/protoc-gen-orm-service]
    out: ./proto
```

Options (`opt`):

| Option   | Default                    | Meaning                                            |
| -------- | -------------------------- | -------------------------------------------------- |
| `namer`  | `{{ .Name }}_svc.g.proto`  | Go text/template for the output filename per file. |

## Structure

```
main.go            flag parsing; wires protogen → Handler
handler.go         parses files into a protobuf-orm graph.Graph, runs app
app/
  app.go           per-file driver: one ast.File per source, one service per entity
  work.go          builds the AST (messages, refs, selects, requests) from the graph
  type.go          proto type-string mapping for a field (incl. map<>)
  x-*.go            one file per RPC (add/get/patch/apply/erase/ref/select)
internal/ast/      a small proto-syntax AST + printer used to emit the .proto text
internal/apptest/  fixtures (generated, gitignored) + a test that drives the plugin
```

The generator never writes `.proto` text directly: it builds an
[`internal/ast`](internal/ast) tree (`File` → `Service` / `Message` / `Enum` …)
and prints it through `ast.Printer`, which handles indentation, package-relative
type names, and import ordering.

## Development

This repo uses a `go.work` that points at the sibling `protobuf-orm` checkout, so
local changes to the library are picked up immediately.

```sh
buf generate     # regenerate apptest fixtures under proto/ and internal/apptest/
go test ./...    # internal/ast, plus internal/apptest end-to-end
go build ./...
```

`internal/apptest` is generated and gitignored, so `buf generate` has to have run
at least once before `go test ./...` will build. Its test runs the plugin over
those fixtures by reading their descriptors out of the protobuf registry, which
means it needs no `protoc` or `buf` of its own — and it asserts on the same text
that lands in `proto/apptest/*_svc.g.proto`.
