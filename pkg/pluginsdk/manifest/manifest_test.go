package manifest_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ContinuumApp/continuum-plugin-sdk/pkg/pluginsdk/manifest"
)

func TestLoadFromDisk(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "manifest.json")

	data := []byte(`{
  "plugin_id": "example.plugin",
  "version": "1.0.0",
  "capabilities": [
    {
      "type": "scheduled_task.v1",
      "id": "hello"
    }
  ]
}`)
	if err := os.WriteFile(path, data, 0644); err != nil {
		t.Fatalf("os.WriteFile(%q) returned error: %v", path, err)
	}

	loaded, err := manifest.LoadFromDisk(path)
	if err != nil {
		t.Fatalf("LoadFromDisk(%q) returned error: %v", path, err)
	}

	if got := loaded.GetPluginId(); got != "example.plugin" {
		t.Fatalf("plugin_id = %q, want example.plugin", got)
	}
	if got := loaded.GetCapabilities()[0].GetId(); got != "hello" {
		t.Fatalf("capability id = %q, want hello", got)
	}
}
