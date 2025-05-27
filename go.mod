module github.com/pekim/go-skia

go 1.23.0

// Address deprecation of clang_getDiagnosticCategoryName.
// Avoids a noisy message when generating api.
replace github.com/go-clang/clang-v15 => github.com/pekim/clang-v15 v0.0.0-20240830114552-c0d27ccce9ec

require (
	github.com/go-clang/clang-v15 v0.0.0-20230222085438-ee3102fa0c71
	github.com/go-gl/gl v0.0.0-20231021071112-07e5d0ea2e71
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20250301202403-da16c1255728
	github.com/stretchr/testify v1.9.0
	golang.org/x/sync v0.14.0
	muzzammil.xyz/jsonc v1.0.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
