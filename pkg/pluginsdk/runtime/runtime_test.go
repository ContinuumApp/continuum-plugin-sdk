package runtime_test

import (
	"testing"

	pluginv1 "github.com/ContinuumApp/continuum-plugin-sdk/pkg/pluginproto/continuum/plugin/v1"
	runtime "github.com/ContinuumApp/continuum-plugin-sdk/pkg/pluginsdk/runtime"
)

func TestRuntimeBootstrapCompiles(t *testing.T) {
	t.Helper()

	_ = runtime.ProtocolVersion
	_ = runtime.HandshakeConfig()
	_ = runtime.ServeConfig{}
	_ = &pluginv1.PluginManifest{}
}
