package apptest_test

import (
	"testing"

	"github.com/protobuf-orm/protoc-gen-orm-service/internal/apptest"
	"github.com/stretchr/testify/require"
)

func TestEdge(t *testing.T) {
	t.Run("XxxPatch", func(t *testing.T) {
		d := apptest.File_apptest_field_svc_g_proto.Messages().ByName("EdgePatchRequest")
		require.NotNil(t, d)
		t.Run("required edge", func(t *testing.T) {
			require.NotNil(t, d.Fields().ByName("required"))
			require.Nil(t, d.Fields().ByName("required_null"))
		})
		t.Run("nullable edge", func(t *testing.T) {
			require.NotNil(t, d.Fields().ByName("nullable"))
			require.NotNil(t, d.Fields().ByName("nullable_null"))
		})
		t.Run("required immutable edge", func(t *testing.T) {
			require.Nil(t, d.Fields().ByName("required_immutable"))
			require.Nil(t, d.Fields().ByName("required_immutable_null"))
		})
		t.Run("nullable immutable edge", func(t *testing.T) {
			require.Nil(t, d.Fields().ByName("nullable_immutable"))
			require.Nil(t, d.Fields().ByName("nullable_immutable_null"))
		})
	})
}
