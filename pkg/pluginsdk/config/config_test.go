package config_test

import (
	"testing"

	pluginv1 "github.com/ContinuumApp/continuum-plugin-sdk/pkg/pluginproto/continuum/plugin/v1"
	"github.com/ContinuumApp/continuum-plugin-sdk/pkg/pluginsdk/config"
)

func TestValidateManifestGlobalValue(t *testing.T) {
	manifest := &pluginv1.PluginManifest{
		GlobalConfigSchema: []*pluginv1.ConfigSchema{
			{
				Key:        "connection",
				Title:      "Connection",
				JsonSchema: `{"type":"object","properties":{"url":{"type":"string"}},"required":["url"],"additionalProperties":false}`,
			},
		},
	}

	t.Run("accepts valid payload", func(t *testing.T) {
		if err := config.ValidateManifestGlobalValue(manifest, "connection", map[string]any{
			"url": "https://api.example.com",
		}); err != nil {
			t.Fatalf("ValidateManifestGlobalValue() returned error: %v", err)
		}
	})

	t.Run("rejects undeclared keys", func(t *testing.T) {
		if err := config.ValidateManifestGlobalValue(manifest, "missing", map[string]any{}); err == nil {
			t.Fatal("expected undeclared config key to be rejected")
		}
	})

	t.Run("rejects invalid payload", func(t *testing.T) {
		if err := config.ValidateManifestGlobalValue(manifest, "connection", map[string]any{
			"url": 42,
		}); err == nil {
			t.Fatal("expected invalid config payload to be rejected")
		}
	})
}

func TestValidateManifestUserValue(t *testing.T) {
	manifest := &pluginv1.PluginManifest{
		UserConfigSchema: []*pluginv1.ConfigSchema{
			{
				Key:        "preferences",
				Title:      "Preferences",
				JsonSchema: `{"type":"object","properties":{"theme":{"type":"string"}},"additionalProperties":false}`,
			},
		},
	}

	if err := config.ValidateManifestUserValue(manifest, "preferences", map[string]any{
		"theme": "midnight",
	}); err != nil {
		t.Fatalf("ValidateManifestUserValue() returned error: %v", err)
	}
}
