package generate

type fileHeader struct {
	*genFile
}

func newFileHeader() *fileHeader {
	f := &fileHeader{
		genFile: newGenFile("./api.h"),
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
	f.writelnf(`
		#ifdef __cplusplus
		}
		#endif // __cplusplus
	`)

	f.close()
	f.clangFormat()
}
