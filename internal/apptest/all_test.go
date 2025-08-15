package apptest_test

import (
	"testing"

	"github.com/protobuf-orm/protoc-gen-orm-service/internal/apptest"
	"github.com/stretchr/testify/require"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

func TestTypes(t *testing.T) {
	const RANGE = 0x1E - 0x10 + 1

	t.Run("XxxAdd", func(t *testing.T) {
		d := apptest.File_apptest_all_svc_g_proto.Messages().ByName("AllAddRequest")
		require.NotNil(t, d)

		t.Run("fields with implicit presence", func(t *testing.T) {
			for i := range RANGE {
				j := i + 0x10
				fd := d.Fields().ByNumber(protoreflect.FieldNumber(j))
				require.False(t, fd.HasPresence())
			}
		})
		t.Run("fields with explicit presence", func(t *testing.T) {
			for i := range RANGE {
				j := i + 0x30
				fd := d.Fields().ByNumber(protoreflect.FieldNumber(j))
				require.True(t, fd.HasPresence())
			}
		})
		t.Run("fields with repeated type", func(t *testing.T) {
			for i := range RANGE {
				j := i + 0x50
				fd := d.Fields().ByNumber(protoreflect.FieldNumber(j))
				require.True(t, fd.IsList())
			}
		})
		t.Run("nullable fields", func(t *testing.T) {
			for i := range RANGE {
				j := i + 0x70
				fd := d.Fields().ByNumber(protoreflect.FieldNumber(j))
				require.True(t, fd.HasPresence())
			}
		})
		t.Run("fields with default value and implicit presence", func(t *testing.T) {
			for i := range RANGE {
				j := i + 0x90
				fd := d.Fields().ByNumber(protoreflect.FieldNumber(j))
				require.True(t, fd.HasPresence())
			}
		})
		t.Run("fields with default value and explicit presence", func(t *testing.T) {
			for i := range RANGE {
				j := i + 0xB0
				fd := d.Fields().ByNumber(protoreflect.FieldNumber(j))
				require.True(t, fd.HasPresence())
			}
		})
		t.Run("immutable fields with implicit presence", func(t *testing.T) {
			for i := range RANGE {
				j := i + 0xD0
				fd := d.Fields().ByNumber(protoreflect.FieldNumber(j))
				require.False(t, fd.HasPresence())
			}
		})
		t.Run("immutable fields with explicit presence", func(t *testing.T) {
			for i := range RANGE {
				j := i + 0xF0
				fd := d.Fields().ByNumber(protoreflect.FieldNumber(j))
				require.True(t, fd.HasPresence())
			}
		})
	})
	t.Run("XxxPatch", func(t *testing.T) {
		d := apptest.File_apptest_all_svc_g_proto.Messages().ByName("AllPatchRequest")
		require.NotNil(t, d)

		t.Run("fields with implicit presence", func(t *testing.T) {
			for i := range RANGE {
				j := (i + 0x10) * 2
				require.NotNil(t, d.Fields().ByNumber(protoreflect.FieldNumber(j-1)))
				require.Nil(t, d.Fields().ByNumber(protoreflect.FieldNumber(j)))
			}
		})
		t.Run("fields with explicit presence", func(t *testing.T) {
			for i := range RANGE {
				j := (i + 0x30) * 2
				require.NotNil(t, d.Fields().ByNumber(protoreflect.FieldNumber(j-1)))
				require.NotNil(t, d.Fields().ByNumber(protoreflect.FieldNumber(j)))
			}
		})
		t.Run("fields with repeated type", func(t *testing.T) {
			for i := range RANGE {
				j := (i + 0x50) * 2
				require.NotNil(t, d.Fields().ByNumber(protoreflect.FieldNumber(j-1)))
				require.Nil(t, d.Fields().ByNumber(protoreflect.FieldNumber(j)))
			}
		})
		t.Run("nullable fields", func(t *testing.T) {
			for i := range RANGE {
				j := (i + 0x70) * 2
				require.NotNil(t, d.Fields().ByNumber(protoreflect.FieldNumber(j-1)))
				require.NotNil(t, d.Fields().ByNumber(protoreflect.FieldNumber(j)))
			}
		})
		t.Run("fields with default value and implicit presence", func(t *testing.T) {
			for i := range RANGE {
				j := (i + 0x90) * 2
				require.NotNil(t, d.Fields().ByNumber(protoreflect.FieldNumber(j-1)))
				require.Nil(t, d.Fields().ByNumber(protoreflect.FieldNumber(j)))
			}
		})
		t.Run("fields with default value and explicit presence", func(t *testing.T) {
			for i := range RANGE {
				j := (i + 0xB0) * 2
				require.NotNil(t, d.Fields().ByNumber(protoreflect.FieldNumber(j-1)))
				require.NotNil(t, d.Fields().ByNumber(protoreflect.FieldNumber(j)))
			}
		})
		t.Run("immutable fields with implicit presence", func(t *testing.T) {
			for i := range RANGE {
				j := (i + 0xD0) * 2
				require.Nil(t, d.Fields().ByNumber(protoreflect.FieldNumber(j-1)))
				require.Nil(t, d.Fields().ByNumber(protoreflect.FieldNumber(j)))
			}
		})
		t.Run("immutable fields with explicit presence", func(t *testing.T) {
			for i := range RANGE {
				j := (i + 0xF0) * 2
				require.Nil(t, d.Fields().ByNumber(protoreflect.FieldNumber(j-1)))
				require.Nil(t, d.Fields().ByNumber(protoreflect.FieldNumber(j)))
			}
		})
	})
}
