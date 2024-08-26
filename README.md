# github.com/pekim/go-skia

`github.com/pekim/go-skia` is a package that provides generated Go bindings
to some of skia's C++ api.

To add it to a Go module, `go get github.com/pekim/go-skia`.

## development

### pre-requisites

Version 15 or later of `clang` is required.

### build

The build script

- downloads skia source (if not already downloaded)
- builds skia
- generates Go bindings
- verifies that the generated bindings result in a package that builds

```sh
./build.sh
```

A warning from `github.com/go-clang/clang-v15/clang` about use of the deprecated
`clang_getDiagnosticCategoryName`
function is annoying but harmless.

### pre-commit hook

There is configuration for a git pre-commit hook
that performs some linting.

- install the `pre-commit` application if not already installed
  - https://pre-commit.com/index.html#install
- install pre-commit hook in this repo's workspace
  - `pre-commit install`

## acknowledgement

Much of `build.sh` was copied from
https://github.com/richardwilkes/cskia/blob/main/build.sh.
