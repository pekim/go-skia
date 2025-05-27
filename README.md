# github.com/pekim/go-skia

[![PkgGoDev](https://pkg.go.dev/badge/github.com/pekim/go-skia)](https://pkg.go.dev/github.com/pekim/go-skia)

`github.com/pekim/go-skia` is a package that provides generated Go bindings
for some of [skia](https://skia.org/)'s C++ api.

To add it to a Go module, `go get github.com/pekim/go-skia`.

## development

### pre-requisites

- version 1.23.0 or later of `Go`
- version 15 or later of `clang`

### build

The build script does the following

- [`rust-skia`](https://github.com/rust-skia/skia) repo
  - clone repo (if not present) to `_skia` dir
  - check out desired tag (if not currently checked out)
  - copy required headers to `header` dir
- [`skia-binaries`](https://github.com/rust-skia/skia-binaries) repo
  - if desired release not present in `lib` dir
    - download release
    - extract in to `lib` dir
- generate Go bindings
- run tests to verify that the generated bindings result in a package that builds

```sh
./build.sh
```

### upgrade skia

To upgrade skia, update some variables in `build.sh`.

- https://github.com/rust-skia/skia/tags
  - `SKIA_TAG` - set to a tag, such as `m136-0.84.2`
- https://github.com/rust-skia/skia-binaries/releases
  - Expand the assets for the latest available minor release that `SKIA_TAG` references.
    For example if `SKIA_TAG` is `m136-0.84.2`, the release to use will be `0.84.0`.
    Do not try to mix minor releases, as it's unlikely to work.
  - `SKIA_BINARIES_COMMITHASH` - set to the hash used in all of the release's downloadable assets
  - `SKIA_BINARIES_TAG` - the release's tag, such as `0.84.0`

### \_skia dir

The leading underscore in the dir name is to prevent the directory appearing in godoc output
in the "Subdirectories" section.

### pre-commit hook

There are configuration files for linting and other checks.
To use a git pre-commit hook for the checks

- install `goimports` if not already installed
  - https://pkg.go.dev/golang.org/x/tools/cmd/goimports
- install `golangci-lint` if not already installed
  - https://golangci-lint.run/usage/install/#local-installation
- install the `pre-commit` application if not already installed
  - https://pre-commit.com/index.html#install
- install a git pre-commit hook in this repo's workspace
  - `pre-commit install`
