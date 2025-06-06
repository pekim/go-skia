/*
Package generate is used to generate Go bindings for the skia library.

Three files are generated, api.go, api.h, and api.cpp.
The C++ file wraps skia functions in C functions.
The header file declares the C functions.
And the Go file calls the C functions using cgo.

The generation happens in 3 phases.
  - An instance of type api is populated with details of what api entities ares to be generated
    (enums, classes, functions, and so on).
    This provides information regarding which entities (classes, structs, methods, enums etc.) are to be generated.
  - A number of skia header files are parsed using libclang, which provides an AST.
    The nodes in each header file's AST are visited.
    Where they correspond to api entities that are to be generated, the entities are enriched with
    information from the AST.
    Methods named enrich1 perform the enrichment.
  - Now that all entities have been enriched, references from one entity to another can be patched up.
    Methods named enrich2 are responsible for doing this.
*/
package generate
