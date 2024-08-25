package generate

type fileHeader struct {
	*genFile
}

func newFileHeader() *fileHeader {
	f := &fileHeader{
		genFile: newGenFile("./api.h"),
	}

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
}
