## use of `panic`

Most errors result in a panic, as generation is all or nothing.
If something has gone wrong there's little point in continuing.

## `ast.dump`

The `ast.dump` file was generated with this command.

```sh
clang -Xclang -ast-dump -fsyntax-only -Iskia/skia -xc++-header -fno-color-diagnostics skia.h > generate/ast.dump
```
