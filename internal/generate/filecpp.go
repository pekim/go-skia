package generate

type fileCpp struct {
	*file
}

func newFileCpp() *fileCpp {
	f := &fileCpp{
		file: newFile("./api.cpp"),
	}

	for _, headerFile := range headerFiles {
		f.writelnf("#include <%s>", headerFile)
	}
	f.writeln()

	f.writeln(`
		#if defined(SKIA_MAC)
			#include "include/ports/SkFontMgr_mac_ct.h"
		#endif

		#if defined(SKIA_UNIX)
			#include "include/ports/SkFontMgr_fontconfig.h"
		#endif

		#if defined(SKIA_WINDOWS)
			#include "include/ports/SkTypeface_win.h"
		#endif
	`)

	f.writelnf(`
		extern "C"
  	{
		#include "api.h"

	`)

	return f
}

func (f fileCpp) finish() {
	f.writelnf(`
		sk_SkFontMgr *sk_fontmgr_ref_default(void)
		{
			#if defined(SKIA_MAC)
				return reinterpret_cast<sk_SkFontMgr *>(SkFontMgr_New_CoreText(nullptr).release());
			#elif defined(SKIA_UNIX)
				return reinterpret_cast<sk_SkFontMgr *>(SkFontMgr_New_FontConfig(nullptr).release());
			#elif defined(SKIA_WINDOWS)
				return reinterpret_cast<sk_SkFontMgr *>(SkFontMgr_New_DirectWrite().release());
			#else
				#error "No font manager available for this platform"
			#endif
		}
	}
	`)

	f.close()
	f.clangFormat()
}
