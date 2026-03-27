# Continuum Plugin SDK

Public Go SDK for building Continuum plugins.

## Packages

- `github.com/ContinuumApp/continuum-plugin-sdk/pkg/pluginproto/continuum/plugin/v1`
- `github.com/ContinuumApp/continuum-plugin-sdk/pkg/pluginsdk/manifest`
- `github.com/ContinuumApp/continuum-plugin-sdk/pkg/pluginsdk/runtime`

## Development

Regenerate protobuf code with:

```sh
make proto
```

Run tests with:

```sh
go test ./...
```
