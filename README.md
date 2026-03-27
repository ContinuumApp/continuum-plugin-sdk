# Continuum Plugin SDK

Public Go SDK for building Continuum plugins.

`continuum-plugin-sdk` is the source of truth for the public plugin authoring contract. First-party consumers such as `Continuum`, `continuum-plugin-tvdb`, and `continuum-plugin-tmdb` should pin tagged semver releases from this repository in `go.mod`. Local multi-repo workspaces may use `go.work` or a temporary `replace`, but CI and release builds must resolve the SDK from a published module tag.

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

## Releases

SDK releases are cut from semver tags such as `v0.1.0` and published through GitHub Actions.

- Additive public API changes belong in a new minor release.
- Compatible fixes and documentation updates belong in a patch release.
- Breaking public API, protobuf, or manifest contract changes require a new major version.

Before downstream repos stop using local workspace overrides, the required SDK commit must be pushed and tagged here first.

## Development

Regenerate protobuf code with:

```sh
make proto
```

Run tests with:

```sh
go test ./...
```

## License

`continuum-plugin-sdk` is licensed under `Apache-2.0`. See [LICENSE](LICENSE).
