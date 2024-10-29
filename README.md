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

### \_skia dir

The leading underscore in the dir name is to prevent the directory appearing in godoc output
in the "Subdirectories" section.

### pre-commit hook

There is configuration for a git pre-commit hook
that performs some linting.

- install the `pre-commit` application if not already installed
  - https://pre-commit.com/index.html#install
- install pre-commit hook in this repo's workspace
  - `pre-commit install`
