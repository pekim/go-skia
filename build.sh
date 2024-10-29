#! /usr/bin/env bash
set -eo pipefail

# vars for rust-skia repo
SKIA_MILESTONE=m130
SKIA_TAG=0.78.1
SKIA_DIR=_skia
# vars for skia-binaries repo
SKIA_BINARIES_COMMITHASH=ec00cf219c4901d785ed
SKIA_BINARIES_TAG=0.78.2

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

# Update libraries
if grep -q -v $SKIA_BINARIES_TAG "lib/tag.txt"; then
  wget \
    -qO- \
    "https://github.com/rust-skia/skia-binaries/releases/download/$SKIA_BINARIES_TAG/skia-binaries-$SKIA_BINARIES_COMMITHASH-x86_64-unknown-linux-gnu-gl-svg-textlayout-x11.tar.gz" | \
    tar xvz
  rm lib/*
  mv skia-binaries/* lib
  rm -r skia-binaries
fi

# Generate & verify
go run internal/generate/cmd/main.go
echo test
go test github.com/pekim/go-skia
