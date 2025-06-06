package generate

func (api *api) configEnums() {
	api.enums = []enum{
		{cppName: "GrSurfaceOrigin"},
		{cppName: "GrSyncCpu"},
		{cppName: "GrSemaphoresSubmitted"},
		{cppName: "SkAlphaType"},
		{cppName: "SkApplyPerspectiveClip"},
		{cppName: "SkBlendMode"},
		{cppName: "SkClipOp"},
		{cppName: "SkColorType"},
		{cppName: "SkFontHinting"},
		{cppName: "SkFilterMode"},
		{cppName: "SkPathDirection"},
		{cppName: "SkPathFillType"},
		{cppName: "SkPathOp"},
		{cppName: "SkPixelGeometry"},
		{cppName: "SkTextEncoding"},
		{cppName: "skgpu::Protected"},
	}
}

func (api *api) configFunctions() {
	api.functions = []callable{
		{cppName: "GrBackendRenderTargets::MakeGL"},
		{
			cppName: "GrDirectContexts::MakeGL",
			overloads: []*callableOverload{
				{Suffix: "InterfaceOptions"},
				{Suffix: "Interface"},
				{Suffix: "Options"},
				{},
			},
		},
		{cppName: "GrGLMakeNativeInterface"},
		{cppName: "SkColorSetARGB"},
		{cppName: "SkColorSetA"},
		{cppName: "SkImages::DeferredFromEncodedData"},
		{cppName: "SkImages::RasterFromData"},
		{cppName: "SkPreMultiplyARGB"},
		{cppName: "SkPreMultiplyColor"},
		{
			cppName: "SkSurfaces::WrapBackendRenderTarget",
			overloads: []*callableOverload{
				{
					Params: []param{
						{},
						{},
						{},
						{},
						{},
						{},
						{ValueNil: true},
						{ValueNil: true},
					},
				},
			},
		},

		{cppName: "Op"},
		{cppName: "Simplify"},
		{cppName: "AsWinding"},
	}
}

func (api *api) configTypedefs() {
	api.typedefs = []typedef{
		{cppName: "SkAlpha"},
		{cppName: "SkColor"},
		{cppName: "SkPMColor"},
	}
}
