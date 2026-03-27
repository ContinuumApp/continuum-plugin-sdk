# Continuum Plugin SDK

Public Go SDK for building Continuum plugins.

## Packages

- `github.com/ContinuumApp/continuum-plugin-sdk/pkg/pluginproto/continuum/plugin/v1`
- `github.com/ContinuumApp/continuum-plugin-sdk/pkg/pluginsdk/config`
- `github.com/ContinuumApp/continuum-plugin-sdk/pkg/pluginsdk/convert`
- `github.com/ContinuumApp/continuum-plugin-sdk/pkg/pluginsdk/manifest`
- `github.com/ContinuumApp/continuum-plugin-sdk/pkg/pluginsdk/runtime`

## Supported Capability Families

- `metadata_provider.v1`
- `media_analyzer.v1`
- `scheduled_task.v1`
- `event_consumer.v1`
- `auth_provider.v1`
- `http_routes.v1`

## Author Workflow

The supported authoring path is Go plugins that:

1. define a plugin manifest using the protobuf contract
2. expose a `Runtime` server plus one or more capability servers
3. support the `manifest` subcommand via `pkg/pluginsdk/runtime`
4. can be installed either from a catalog or by uploading a trusted binary to a Continuum server

For a minimal self-describing plugin, see [`examples/hello-scheduled-task`](examples/hello-scheduled-task).

## Self-Describing Binaries

Direct binary upload works best when the plugin embeds a manifest template and computes its own executable checksum at runtime before returning `Runtime.GetManifest`. That keeps the plugin installable without requiring a checked-out Continuum repository or a sibling `manifest.json` file at upload time.

The example plugin shows this pattern.

## Compatibility

Compatibility and versioning expectations are documented in [`docs/compatibility.md`](docs/compatibility.md).

## Development

Regenerate protobuf code with:

```sh
make proto
```

Run tests with:

```sh
go test ./...
```
