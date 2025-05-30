#! /usr/bin/env bash

set -eo pipefail

# vars for rust-skia repo
SKIA_TAG=m138-0.86.0
SKIA_DIR=_skia
# vars for skia-binaries repo
SKIA_BINARIES_HASH=cab569e6478958ca0783
SKIA_BINARIES_TAG=0.86.0

function get_binaries() {
  GOOS=$1
  GOARCH=$2
  FILE_VARIATION=$3
  LIB_DIR="lib/$GOOS-$GOARCH"
  GO_FILE="libpath_${GOOS}_${GOARCH}.go"

  TAG_FILE="$LIB_DIR/tag.txt"
  if test -f $TAG_FILE && grep -q $SKIA_BINARIES_TAG $TAG_FILE; then
    # already have the binaries for tag
    return
  fi

  echo "get binaries for $GOOS $GOARCH"

  # get the OS/arch specific libraries
  wget \
    -qO- \
    "https://github.com/rust-skia/skia-binaries/releases/download/$SKIA_BINARIES_TAG/skia-binaries-$SKIA_BINARIES_HASH-$FILE_VARIATION.tar.gz" | \
    tar xz
  rm -fr $LIB_DIR
  mkdir -p $LIB_DIR
  mv skia-binaries/* $LIB_DIR
  rm -r skia-binaries

  # write a file to use the OS/arch specific lib dir when linking
  cat << EOF > $GO_FILE
//go:build $GOOS && $GOARCH

// This is a generated file. DO NOT EDIT.

package skia

// #cgo LDFLAGS: -L \${SRCDIR}/$LIB_DIR
import "C"
EOF
}

# clone rust-skia repo if not present
if [ ! -e $SKIA_DIR ]; then
  git clone https://github.com/rust-skia/skia.git $SKIA_DIR
fi
# checkout required rust-skia repo tag
pushd $SKIA_DIR > /dev/null
if [[ $(git describe --always --tags) != "$SKIA_TAG" ]];then
  git fetch --tags
  git checkout "$SKIA_TAG"
fi
popd > /dev/null

# copy headers
rm -fr skia
mkdir skia
cp -r _skia/include _skia/modules _skia/src skia
find skia/ -type f | \
  grep -v ".*\.h$" | \
  grep -v LICENSE | \
  xargs rm # remove all files except header files and licences

# get pre-built binaries for multiple OSes and architectures
get_binaries "darwin" "arm64" "aarch64-apple-darwin-gl-pdf-svg-textlayout"
get_binaries "darwin" "amd64" "x86_64-apple-darwin-gl-pdf-svg-textlayout"
get_binaries "linux" "amd64" "x86_64-unknown-linux-gnu-gl-pdf-svg-textlayout-x11"
find lib -name "bindings.rs" | xargs rm -f

# Generate & verify
go run internal/generate/cmd/main.go
echo test
go test github.com/pekim/go-skia
