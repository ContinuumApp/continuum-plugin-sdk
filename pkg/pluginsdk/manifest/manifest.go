package manifest

import (
	"fmt"
	"io/fs"

	"google.golang.org/protobuf/encoding/protojson"

	pluginv1 "github.com/ContinuumApp/continuum-plugin-sdk/pkg/pluginproto/continuum/plugin/v1"
)

func Load(data []byte) (*pluginv1.PluginManifest, error) {
	var manifest pluginv1.PluginManifest
	if err := protojson.Unmarshal(data, &manifest); err != nil {
		return nil, fmt.Errorf("decode plugin manifest: %w", err)
	}
	if err := Validate(&manifest); err != nil {
		return nil, err
	}
	return &manifest, nil
}

func MustLoad(data []byte) *pluginv1.PluginManifest {
	manifest, err := Load(data)
	if err != nil {
		panic(err)
	}
	return manifest
}

func Validate(manifest *pluginv1.PluginManifest) error {
	if manifest == nil {
		return fmt.Errorf("plugin manifest is required")
	}
	if manifest.PluginId == "" {
		return fmt.Errorf("plugin manifest plugin_id is required")
	}
	if manifest.Version == "" {
		return fmt.Errorf("plugin manifest version is required")
	}
	for _, capability := range manifest.Capabilities {
		if capability.Type == "" {
			return fmt.Errorf("plugin capability type is required")
		}
		if capability.Id == "" {
			return fmt.Errorf("plugin capability id is required")
		}
	}
	return nil
}

func RegisterHTTPRoutes(manifest *pluginv1.PluginManifest, routes ...*pluginv1.HttpRouteDescriptor) error {
	if err := Validate(manifest); err != nil {
		return err
	}
	manifest.HttpRoutes = append(manifest.HttpRoutes, routes...)
	return nil
}

func RegisterAssets(manifest *pluginv1.PluginManifest, assets ...*pluginv1.PackagedAsset) error {
	if err := Validate(manifest); err != nil {
		return err
	}
	manifest.Assets = append(manifest.Assets, assets...)
	return nil
}

func Asset(path string, filesystem fs.FS) (*pluginv1.PackagedAsset, error) {
	if _, err := fs.Stat(filesystem, path); err != nil {
		return nil, fmt.Errorf("plugin asset %q: %w", path, err)
	}
	return &pluginv1.PackagedAsset{Path: path}, nil
}
