# Compatibility and Versioning

## Scope

`continuum-plugin-sdk` is the public build-time contract for Go plugin authors.

This repository is released as a semver-governed Go module. Third-party plugins and first-party consumers should depend on tagged releases, not on sibling repo checkouts or workspace-only overrides.

The compatibility boundary includes:

- protobuf messages and gRPC services under `pkg/pluginproto/continuum/plugin/v1`
- runtime bootstrap behavior in `pkg/pluginsdk/runtime`
- manifest helpers in `pkg/pluginsdk/manifest`
- config validation helpers in `pkg/pluginsdk/config`
- generic capability metadata conversion helpers in `pkg/pluginsdk/convert`

## Versioning Rules

- Treat the SDK as a semver boundary.
- Publish semver tags from this repository and consume those tags from downstream repos.
- Prefer additive protobuf evolution.
- Avoid renaming or removing protobuf fields, services, or enum values in `v1`.
- Keep plugin capability expansion additive: new functionality should arrive as new capability families or additive fields, not breaking rewrites of existing ones.
- First-party consumers should not merge code that depends on new SDK packages or symbols until the required SDK tag exists.

## Consumer Rules

- `Continuum`, `continuum-plugin-tvdb`, and `continuum-plugin-tmdb` should pin released SDK tags in `go.mod`.
- CI and release pipelines should build with `GOWORK=off` and without checking out this repo as a sibling source dependency.
- Local `go.work` files and temporary `replace` directives are acceptable for development, but they must not be committed as the release path.

## Runtime Compatibility

- `continuum_api_version` is the coarse runtime compatibility gate between Continuum and a plugin binary.
- Host installs should reject incompatible API versions before runtime startup.
- A plugin binary should return the same manifest shape that Continuum installs, except that binaries may compute their checksum dynamically at runtime.

## Metadata Language Compatibility

- For `metadata_provider.v1`, Continuum sends request `language` values in ISO 639-1 form.
- Plugins own translation from that host value to provider-specific language formats such as BCP 47 or ISO 639-3.
- Plugins may tolerate region-qualified input for compatibility, but the public Continuum contract does not require region-qualified tags.

## Go Support

The supported public authoring path today is Go-only.

The protobuf and gRPC contracts are the long-term compatibility source of truth, but non-Go authoring is not an official support target in this release.

## Self-Describing Binary Guidance

If a plugin should be installable by direct binary upload:

- embed a manifest template in the binary
- compute the executable checksum at runtime
- return that populated manifest from `Runtime.GetManifest`

This keeps the plugin installable without requiring external repo state at upload time.
