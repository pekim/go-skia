package generate

type fileHeader struct {
	*file
}

func newFileHeader() *fileHeader {
	f := &fileHeader{
		file: newFile("./api.h"),
	}

	f.writeln("#include <stdbool.h>")
	f.writeln("#include <sys/types.h>")
	f.writeln("typedef signed char schar;")
	f.writeln("typedef unsigned char uchar;")

	f.writelnf(`
		#ifdef __cplusplus
		extern "C" {
		#endif // __cplusplus
	`)

	return f
}

func (f fileHeader) finish() {
	f.writeln(`  sk_SkFontMgr *sk_fontmgr_ref_default (void);`)

	f.writelnf(`
		#ifdef __cplusplus
		}
		#endif // __cplusplus
	`)

	f.close()
	f.clangFormat()
}
