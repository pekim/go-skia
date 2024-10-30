#! /usr/bin/env bash
set -eo pipefail

# vars for rust-skia repo
SKIA_MILESTONE=m130
SKIA_TAG=0.78.1
SKIA_DIR=_skia
# vars for skia-binaries repo
SKIA_BINARIES_COMMITHASH=ec00cf219c4901d785ed
SKIA_BINARIES_TAG=0.78.2

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
    "https://github.com/rust-skia/skia-binaries/releases/download/$SKIA_BINARIES_TAG/skia-binaries-$SKIA_BINARIES_COMMITHASH-$FILE_VARIATION.tar.gz" | \
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

# checkout required tag
pushd $SKIA_DIR > /dev/null
if [[ $(git describe --always --tags) != "$SKIA_MILESTONE-$SKIA_TAG" ]];then
  git fetch --tags
  git checkout "$SKIA_MILESTONE-$SKIA_TAG"
fi
popd > /dev/null

# copy headers
rm -fr header
mkdir header && cp -r _skia/include _skia/modules header
mkdir -p header/src/base && cp -r _skia/src/base/SkTLazy.h _skia/src/base/SkMathPriv.h header/src/base
mkdir -p header/src/core && cp -r _skia/src/core/SkTHash.h _skia/src/core/SkChecksum.h header/src/core

# get pre-built binaries for multiple OSes and architectures
get_binaries "darwin" "arm64" "aarch64-apple-darwin-gl-svg-textlayout"
get_binaries "darwin" "amd64" "x86_64-apple-darwin-gl-svg-textlayout"
get_binaries "linux" "amd64" "x86_64-unknown-linux-gnu-gl-svg-textlayout-x11"

# Generate & verify
go run internal/generate/cmd/main.go
echo test
go test github.com/pekim/go-skia
