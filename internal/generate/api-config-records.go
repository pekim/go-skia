package generate

func (api *api) configRecords() {
	api.records = []record{
		{cppName: "GrGLFramebufferInfo", noWrapper: true},

		{
			cppName: "GrBackendRenderTarget",
			ctors: []*recordCtor{
				{}, nil, {Suffix: "Copy"},
			},
			methods: []callable{
				{cppName: "dimensions"},
				{cppName: "width"},
				{cppName: "height"},
				{cppName: "sampleCnt"},
				{cppName: "stencilBits"},
				{cppName: "isFramebufferOnly"},
			},
		},

		{
			cppName: "GrDirectContext",
			as:      []string{"GrRecordingContext"},
			methods: []callable{
				{
					cppName:   "flushAndSubmit",
					overloads: []*callableOverload{{}, nil, nil},
				},
				{
					cppName:   "flush",
					overloads: []*callableOverload{{}, nil, nil, nil, nil, nil, nil},
				},
				{
					cppName:   "submit",
					overloads: []*callableOverload{nil, {}},
				},
			},
		},
		{cppName: "GrFlushInfo"},
		{cppName: "GrRecordingContext"},

		{cppName: "GrGLInterface", ctors: []*recordCtor{nil}},
		{
			cppName:   "GrContextOptions",
			noWrapper: true,
			ctors:     []*recordCtor{{}},
			enums:     []enum{{cppName: "Enable"}, {cppName: "ShaderCacheStrategy"}},
		},

		{
			cppName: "SkArc",
			ctors:   []*recordCtor{{}, {Suffix: "Copy"}},
		},

		{
			cppName: "SkBitmap",
			ctors:   []*recordCtor{{}, {Suffix: "Copy"}, nil},
			methods: []callable{
				{cppName: "ComputeIsOpaque"},
				{cppName: "bytesPerPixel"},
				{cppName: "rowBytesAsPixels"},
				{cppName: "shiftPerPixel"},
				{cppName: "drawsNothing"},
				{cppName: "rowBytes"},
				{cppName: "isImmutable"},
				{cppName: "setImmutable"},
				{cppName: "bounds"},
				{cppName: "dimensions"},
			},
		},

		{
			cppName: "SkCanvas",
			ctors: []*recordCtor{
				{},
				{Suffix: "WithDimensions"},
				nil,
				{Suffix: "FromBitmap"},
				{Suffix: "FromBitmapSurfaceProps"},
			},
			enums: []enum{
				{cppName: "ClipEdgeStyle"},
				{cppName: "PointMode"},
				{cppName: "QuadAAFlags"},
				{cppName: "SaveLayerFlagsSet"},
				{cppName: "SaveLayerStrategy"},
				{cppName: "SrcRectConstraint"},
			},
			methods: []callable{
				{cppName: "getProps"},
				{cppName: "getBaseProps"},
				{cppName: "getTopProps"},
				{cppName: "getBaseLayerSize"},
				{cppName: "makeSurface"},
				{cppName: "getSurface"},
				{cppName: "recordingContext"},
				{cppName: "peekPixels"},
				{
					cppName: "readPixels",
					overloads: []*callableOverload{
						{Suffix: "ImageInfo"},
						{Suffix: "Pixmap"},
						{Suffix: "Bitmap"},
					},
				},
				{
					cppName:   "writePixels",
					overloads: []*callableOverload{{Suffix: "ImageInfo"}, {Suffix: "Bitmap"}},
				},
				{cppName: "save"},
				{
					cppName:   "saveLayer",
					overloads: []*callableOverload{{}, nil, nil},
				},
				{cppName: "saveLayerAlpha"},
				{cppName: "restore"},
				{cppName: "getSaveCount"},
				{cppName: "restoreToCount"},
				{cppName: "translate"},
				{cppName: "scale"},
				{cppName: "rotate", overloads: []*callableOverload{{}, {Suffix: "AboutPoint"}}},
				{cppName: "skew"},
				{
					cppName:   "concat",
					overloads: []*callableOverload{{Suffix: "Matrix"}, {Suffix: "M44"}},
				},
				{
					cppName:   "setMatrix",
					overloads: []*callableOverload{{Suffix: "M44"}, {}},
				},
				{cppName: "resetMatrix"},
				{cppName: "clipRect", overloads: []*callableOverload{{}, nil, nil}},
				{cppName: "clipRRect", overloads: []*callableOverload{{}, nil, nil}},
				{cppName: "clipPath", overloads: []*callableOverload{{}, nil, nil}},
				{cppName: "clipRegion"},
				{
					cppName:   "quickReject",
					overloads: []*callableOverload{{Suffix: "Rect"}, {Suffix: "Path"}},
				},
				{
					cppName: "getLocalClipBounds",
					overloads: []*callableOverload{
						{Suffix: "Rect"},
						{Suffix: "Path", Params: []param{{Out: true}}},
					},
				},
				{
					cppName:   "getDeviceClipBounds",
					overloads: []*callableOverload{{}, {Suffix: "Rect", Params: []param{{Out: true}}}},
				},
				{
					cppName:   "drawColor",
					overloads: []*callableOverload{{}, {Suffix: "4f"}},
				},
				{
					cppName:   "clear",
					overloads: []*callableOverload{{}, {Suffix: "4f"}},
				},
				{cppName: "discard"},
				{cppName: "drawPaint"},
				{
					cppName:   "drawPoint",
					overloads: []*callableOverload{{Suffix: "Scalars"}, {}},
				},
				{
					cppName:   "drawLine",
					overloads: []*callableOverload{{Suffix: "Scalars"}, {Suffix: "Points"}},
				},
				{cppName: "drawRect"},
				{cppName: "drawIRect"},
				{cppName: "drawRegion"},
				{cppName: "drawOval"},
				{cppName: "drawRRect"},
				{cppName: "drawDRRect"},
				{
					cppName:   "drawCircle",
					overloads: []*callableOverload{{Suffix: "Scalars"}, {Suffix: "Point"}},
				},
				{cppName: "drawArc", overloads: []*callableOverload{{}, {Suffix: "Arc"}}},
				{cppName: "drawRoundRect"},
				{cppName: "drawPath"},
				{
					cppName:   "drawImage",
					overloads: []*callableOverload{{}, nil, {Suffix: "SamplingOptions"}, nil},
				},
				{
					cppName:   "drawImageRect",
					overloads: []*callableOverload{{}, {Suffix: "NoSrc"}, nil, nil},
				},
				{cppName: "drawImageNine"},
				{cppName: "drawString", overloads: []*callableOverload{{}, nil}},
				{
					cppName:   "drawGlyphs",
					overloads: []*callableOverload{{Suffix: "Clusters"}, {}, nil},
				},
				{cppName: "drawTextBlob", overloads: []*callableOverload{{}, nil}},
			},
		},

		{
			cppName: "SkColorInfo",
			ctors:   []*recordCtor{{}, {Suffix: "2"}, {Suffix: "Copy"}, nil},
			methods: []callable{
				{cppName: "colorSpace"},
				{cppName: "refColorSpace"},
				{cppName: "colorType"},
				{cppName: "alphaType"},
				{cppName: "isOpaque"},
				{cppName: "gammaCloseToSRGB"},
				{cppName: "makeAlphaType"},
				{cppName: "makeColorType"},
				{cppName: "makeColorSpace"},
				{cppName: "bytesPerPixel"},
				{cppName: "shiftPerPixel"},
			},
		},

		{
			cppName: "SkColorSpace",
			methods: []callable{
				{cppName: "MakeSRGB"},
				{cppName: "MakeSRGBLinear"},
				{cppName: "Equals"},
			},
		},

		{
			cppName: "SkData",
			methods: []callable{
				{cppName: "size"},
				{cppName: "isEmpty"},
				{cppName: "MakeWithCopy"},
				{cppName: "MakeZeroInitialized"},
				{cppName: "MakeWithCString"},
				{cppName: "MakeWithoutCopy"},
				{cppName: "MakeFromFileName"},
			},
		},

		{
			cppName: "SkFont",
			ctors: []*recordCtor{
				{},
				{Suffix: "TypefaceSize"},
				{Suffix: "Typeface"},
				{Suffix: "TypefaceSizeScaleSkew"},
			},
			enums: []enum{{cppName: "Edging"}},
			methods: []callable{
				{
					cppName:   "getMetrics",
					overloads: []*callableOverload{{Params: []param{{Out: true}}}},
				},
				{cppName: "getXPos"},
				{
					cppName: "measureText",
					overloads: []*callableOverload{
						{Params: []param{{}, {}, {}, {Out: true}}},
						{
							Suffix: "Paint",
							Params: []param{{}, {}, {}, {Out: true}, {}},
						},
					},
				},
				{cppName: "setForceAutoHinting"},
				{cppName: "getEdging"},
				{cppName: "setEdging"},
				{cppName: "setHinting"},
				{cppName: "setSubpixel"},
				{cppName: "textToGlyphs"},
				{cppName: "unicharToGlyph"},
				{cppName: "unicharsToGlyphs"},
				{cppName: "getSize"},
				{cppName: "getTypeface"},
				{cppName: "getWidths", overloads: []*callableOverload{{Suffix: "Bounds"}, nil, {}}},
			},
		},

		{
			cppName: "SkFontArguments",
			ctors:   []*recordCtor{{}},
			methods: []callable{
				{cppName: "setCollectionIndex"},
				// { CppName: "setVariationDesignPosition" }
				{cppName: "getCollectionIndex"},
				{cppName: "getVariationDesignPosition"},
				// { CppName: "setPalette" },
				{cppName: "getPalette"},
			},
			records: []record{{cppName: "VariationPosition"}, {cppName: "Palette"}},
		},

		{
			cppName:   "SkFontMetrics",
			noWrapper: true,
		},

		{
			cppName: "SkFontMgr",
			methods: []callable{
				{cppName: "matchFamily"},
				{cppName: "matchFamilyStyle"},
				{cppName: "makeFromData"},
				{cppName: "makeFromFile"},
			},
		},

		{
			cppName: "SkFontStyle",
			ctors:   []*recordCtor{{Suffix: "2"}, {}},
			enums:   []enum{{cppName: "Slant"}, {cppName: "Weight"}, {cppName: "Width"}},
			methods: []callable{
				{cppName: "Normal"},
				{cppName: "Bold"},
				{cppName: "Italic"},
				{cppName: "BoldItalic"},
			},
		},

		{
			cppName: "SkFontStyleSet",
			methods: []callable{
				{cppName: "count"},
				{cppName: "getStyle"},
				{cppName: "createTypeface"},
				{cppName: "matchStyle"},
			},
		},

		{
			cppName: "SkImage",
			enums:   []enum{{cppName: "CachingHint"}},
			methods: []callable{
				{cppName: "imageInfo"},
				{cppName: "width"},
				{cppName: "height"},
				{cppName: "dimensions"},
				{cppName: "bounds"},
				{cppName: "uniqueID"},
				{cppName: "alphaType"},
				{cppName: "colorType"},
				{cppName: "colorSpace"},
				{cppName: "isAlphaOnly"},
				{cppName: "isOpaque"},
				{cppName: "readPixels", overloads: []*callableOverload{{}, nil, nil, nil}},
				{cppName: "makeSubset", overloads: []*callableOverload{{}, nil}},
			},
		},

		{
			cppName: "SkImageInfo",
			ctors:   []*recordCtor{{}},
			methods: []callable{
				{cppName: "width"},
				{cppName: "height"},
				{cppName: "colorType"},
				{cppName: "alphaType"},
				{cppName: "colorSpace"},
				{cppName: "refColorSpace"},
				{cppName: "isEmpty"},
				{cppName: "colorInfo"},
				{cppName: "isOpaque"},
				{cppName: "dimensions"},
				{cppName: "bounds"},
				{cppName: "gammaCloseToSRGB"},
				{cppName: "makeWH"},
				{cppName: "makeDimensions"},
				{cppName: "makeAlphaType"},
				{cppName: "makeColorType"},
				{cppName: "makeColorSpace"},
				{cppName: "bytesPerPixel"},
				{cppName: "shiftPerPixel"},
				{cppName: "minRowBytes64"},
				{cppName: "minRowBytes"},
				{cppName: "computeOffset"},
				{cppName: "computeByteSize"},
				{cppName: "computeMinByteSize"},
				{cppName: "validRowBytes"},
				{cppName: "reset"},
			},
		},

		{
			cppName: "SkM44",
			ctors: []*recordCtor{
				{Suffix: "Copy"},
				{},
				{Suffix: "AB"},
				nil,
				nil,
				{Suffix: "Scalars"},
				nil,
			},
		},

		{
			cppName: "SkMemoryStream",
			as:      []string{"SkStream"},
			ctors:   []*recordCtor{nil, nil, nil, nil},
			methods: []callable{{cppName: "Make"}, {cppName: "MakeDirect"}},
		},

		{
			cppName: "SkIPoint",
		},

		{
			cppName:   "SkIRect",
			noWrapper: true,
			methods: []callable{
				{cppName: "MakeEmpty"},
				{cppName: "MakeWH"},
				{cppName: "MakeLTRB"},
				{cppName: "MakeXYWH"},
				{cppName: "MakeSize"},
				{cppName: "Intersects"},

				{cppName: "x"},
				{cppName: "y"},
				{cppName: "width"},
				{cppName: "height"},
				{cppName: "isEmpty"},
				{cppName: "setEmpty"},
				{cppName: "setLTRB"},
				{cppName: "setXYWH"},
				{cppName: "setWH"},
				{cppName: "offset", overloads: []*callableOverload{{}, {Suffix: "Point"}}},
				{cppName: "offsetTo"},
				{cppName: "inset"},
				{cppName: "outset"},
				{cppName: "adjust"},
				{
					cppName:   "contains",
					overloads: []*callableOverload{{}, {Suffix: "Rect"}, nil},
				},
				{
					cppName: "containsNoEmptyCheck",
				},
				{
					cppName:   "intersect",
					overloads: []*callableOverload{{}, nil},
				},
				{cppName: "join"},
				{cppName: "sort"},
			},
		},

		{
			cppName:   "SkISize",
			noWrapper: true,
			methods:   []callable{},
		},

		{
			cppName: "SkMatrix",
			ctors:   []*recordCtor{{}},
		},

		{
			cppName: "SkOpBuilder",
			ctors:   []*recordCtor{{}},
			methods: []callable{{cppName: "add"}, {cppName: "resolve"}},
		},

		{
			cppName: "SkPaint",
			ctors:   []*recordCtor{{}, nil, {Suffix: "Copy"}, nil},
			methods: []callable{
				{cppName: "reset"},
				{cppName: "getAlpha"},
				{cppName: "setAlpha"},
				{cppName: "setARGB"},
				{cppName: "setAntiAlias"},
				{cppName: "setBlendMode"},
				{cppName: "getColor", overloads: []*callableOverload{{}}},
				{cppName: "setColor", overloads: []*callableOverload{{}, nil}},
				{cppName: "setDither"},
				{cppName: "getStrokeCap"},
				{cppName: "setStrokeCap"},
				{cppName: "getStrokeJoin"},
				{cppName: "setStrokeJoin"},
				{cppName: "getStrokeMiter"},
				{cppName: "setStrokeMiter"},
				{cppName: "getStrokeWidth"},
				{cppName: "setStrokeWidth"},
				{cppName: "getStyle"},
				{cppName: "setStyle"},
			},
			enums: []enum{{cppName: "Cap"}, {cppName: "Join"}, {cppName: "Style"}},
		},

		{
			cppName: "SkPath",
			ctors:   []*recordCtor{{}, {Suffix: "Copy"}},
			enums: []enum{
				{cppName: "AddPathMode"},
				{cppName: "ArcSize"},
				{cppName: "SegmentMask"},
				{cppName: "Verb"},
			},
			methods: []callable{
				{cppName: "isInterpolatable"},
				{cppName: "interpolate"},
				{cppName: "getFillType"},
				{cppName: "setFillType"},
				{cppName: "isInverseFillType"},
				{cppName: "toggleInverseFillType"},
				{cppName: "isConvex"},
				{cppName: "isOval"},
				{cppName: "isRRect"},
				{cppName: "reset"},
				{cppName: "rewind"},
				{cppName: "isEmpty"},
				{cppName: "isLastContourClosed"},
				{cppName: "isFinite"},
				{cppName: "isVolatile"},
				{cppName: "setIsVolatile"},
				// { CppName: "isLine" },
				{cppName: "countPoints"},
				{cppName: "getPoint"},
				{cppName: "getPoints"},
				{cppName: "countVerbs"},
				{cppName: "getVerbs"},
				{cppName: "approximateBytesUsed"},
				{cppName: "swap"},
				{cppName: "getBounds"},
				{cppName: "updateBoundsCache"},
				{cppName: "computeTightBounds"},
				{cppName: "conservativelyContainsRect"},
				{cppName: "incReserve"},
				{cppName: "moveTo", overloads: []*callableOverload{{Suffix: "Point"}, {}}},
				{cppName: "rMoveTo"},
				{cppName: "lineTo", overloads: []*callableOverload{{}, {Suffix: "Point"}}},
				{cppName: "rLineTo"},
				{cppName: "quadTo", overloads: []*callableOverload{{}, {Suffix: "Point"}}},
				{cppName: "rQuadTo"},
				{cppName: "conicTo", overloads: []*callableOverload{{}, {Suffix: "Points"}}},
				{cppName: "rConicTo"},
				{cppName: "cubicTo", overloads: []*callableOverload{{}, {Suffix: "Points"}}},
				{cppName: "rCubicTo"},
				{
					cppName: "arcTo",
					overloads: []*callableOverload{
						{Suffix: "1"},
						{Suffix: "2"},
						{Suffix: "3"},
						{Suffix: "4"},
						{Suffix: "5"},
					},
				},
				{cppName: "rArcTo"},
				{cppName: "close"},
				{
					cppName: "isRect",
					overloads: []*callableOverload{
						{Params: []param{{Out: true}, {Out: true}, {Out: true}}},
					},
				},
				{
					cppName:   "addRect",
					overloads: []*callableOverload{{Suffix: "1"}, {Suffix: "2"}, {Suffix: "3"}},
				},
				{
					cppName:   "addOval",
					overloads: []*callableOverload{{Suffix: "1"}, {Suffix: "2"}},
				},
				{cppName: "addCircle"},
				{cppName: "addArc"},
				{
					cppName:   "addRoundRect",
					overloads: []*callableOverload{{Suffix: "1"}, {Suffix: "2"}},
				},
				{
					cppName:   "addRRect",
					overloads: []*callableOverload{{Suffix: "1"}, {Suffix: "2"}},
				},
				{
					cppName:   "addPoly",
					overloads: []*callableOverload{{}, nil},
				},
				{
					cppName:   "addPath",
					overloads: []*callableOverload{{Suffix: "1"}, {Suffix: "2"}, {Suffix: "3"}},
				},
				{cppName: "reverseAddPath"},
				{
					cppName:   "offset",
					overloads: []*callableOverload{{Suffix: "1"}, {Suffix: "2"}},
				},
				{
					cppName:   "transform",
					overloads: []*callableOverload{{Suffix: "1"}, {Suffix: "2"}},
				},
				{cppName: "makeTransform"},
				{cppName: "makeScale"},
				{cppName: "getLastPt"},
				{
					cppName:   "setLastPt",
					overloads: []*callableOverload{{Suffix: "1"}, {Suffix: "2"}},
				},
				{cppName: "getSegmentMasks"},
				{cppName: "contains"},
				{cppName: "serialize"},
				{cppName: "isValid"},
			},
		},

		{
			cppName:   "SkPoint",
			noWrapper: true,
			methods:   []callable{{cppName: "Make"}},
		},

		{
			cppName: "SkPixmap",
			ctors:   []*recordCtor{{}, {Suffix: "ImageInfo"}},
		},

		{
			cppName:   "SkRect",
			noWrapper: true,
			methods: []callable{
				{cppName: "MakeEmpty"},
				{cppName: "MakeWH"},
				{cppName: "MakeSize"},
				{cppName: "MakeLTRB"},
				{cppName: "MakeXYWH"},
				{
					cppName:   "Make",
					overloads: []*callableOverload{{Suffix: "ISize"}, {Suffix: "IRect"}},
				},
				{cppName: "Intersects"},

				{cppName: "x"},
				{cppName: "y"},
				{cppName: "width"},
				{cppName: "height"},
				{cppName: "centerX"},
				{cppName: "centerY"},
				{cppName: "isEmpty"},
				{cppName: "setEmpty"},
				{cppName: "setLTRB"},
				{cppName: "setXYWH"},
				{cppName: "setWH"},
				{cppName: "offset", overloads: []*callableOverload{{}, nil}},
				{cppName: "offsetTo"},
				{cppName: "inset"},
				{cppName: "outset"},
				{
					cppName:   "contains",
					overloads: []*callableOverload{{}, {Suffix: "Rect"}, nil},
				},
				{
					cppName:   "intersect",
					overloads: []*callableOverload{{}, nil},
				},
				{cppName: "join"},
				{cppName: "sort"},
			},
		},

		{
			cppName:   "SkRRect",
			noWrapper: true,
			ctors:     []*recordCtor{{}, {Suffix: "Copy"}},
		},

		{
			cppName: "SkRegion",
			ctors:   []*recordCtor{{}, {Suffix: "Copy"}, {Suffix: "CopyRect"}},
		},

		{
			cppName:   "SkRGBA4f",
			noWrapper: true,
		},

		{
			cppName: "SkSamplingOptions",
			ctors:   []*recordCtor{{}, {Suffix: "Copy"}, nil, nil, nil},
		},

		{
			cppName:   "SkSize",
			noWrapper: true,
		},

		{
			cppName: "SkStream",
			ctors:   []*recordCtor{nil},
		},

		{
			cppName: "SkString",
			ctors:   []*recordCtor{nil, nil, {}, nil, nil, nil, nil, nil},
			methods: []callable{{cppName: "data", overloads: []*callableOverload{{}, nil}}},
		},

		{
			cppName: "SkSurface",
			methods: []callable{
				{cppName: "getCanvas"},
				{cppName: "makeSurface", overloads: []*callableOverload{nil, {}}},
				{
					cppName:   "makeImageSnapshot",
					overloads: []*callableOverload{{}, {Suffix: "Bounds"}},
				},
			},
		},

		{
			cppName: "SkSurfaceProps",
			ctors: []*recordCtor{
				{},
				{Suffix: "PixelGeometry"},
				nil,
				{Suffix: "Copy"},
			},
			enums: []enum{{cppName: "Flags"}},
		},

		{
			cppName:    "SkSVGDOM",
			dtorCreate: true,
			methods: []callable{
				{cppName: "getRoot"},
				{cppName: "setContainerSize"},
				{cppName: "containerSize"},
				{cppName: "render"},
			},
			records: []record{
				{
					cppName: "Builder",
					ctors:   []*recordCtor{{}},
					methods: []callable{
						{cppName: "setFontManager"},
						// { CppName: "setResourceProvider" },
						{cppName: "make"},
					},
				},
			},
		},

		{
			cppName: "SkSVGSVG",
			methods: []callable{{cppName: "intrinsicSize"}},
		},

		{
			cppName: "SkSVGLengthContext",
			ctors:   []*recordCtor{{}},
			methods: []callable{{cppName: "viewPort"}, {cppName: "setViewPort"}},
		},

		{
			cppName: "SkTextBlob",
			methods: []callable{
				{cppName: "bounds"},
				{cppName: "uniqueID"},
				{cppName: "MakeFromString"},
				{cppName: "MakeFromPosTextH"},
				{cppName: "MakeFromPosText"},
			},
		},

		{
			cppName: "SkTextBlobBuilder",
			ctors:   []*recordCtor{{}},
			methods: []callable{
				{cppName: "make"},
				{
					cppName:   "allocRun",
					overloads: []*callableOverload{{Params: []param{{}, {}, {}, {}, {Out: true}}}},
				},
				{
					cppName:   "allocRunPosH",
					overloads: []*callableOverload{{Params: []param{{}, {}, {}, {Out: true}}}},
				},
				{
					cppName:   "allocRunPos",
					overloads: []*callableOverload{{Params: []param{{}, {}, {Out: true}}}},
				},
				{
					cppName: "allocRunText",
					overloads: []*callableOverload{
						{Params: []param{{}, {}, {}, {}, {}, {Out: true}}},
					},
				},
				{
					cppName:   "allocRunTextPosH",
					overloads: []*callableOverload{{Params: []param{{}, {}, {}, {}, {Out: true}}}},
				},
				{
					cppName:   "allocRunTextPos",
					overloads: []*callableOverload{{Params: []param{{}, {}, {}, {Out: true}}}},
				},
			},
			records: []record{
				{
					cppName: "RunBuffer",
					// Methods:[]callable{}
					//  [{ CppName: "points" }]
				},
			},
		},

		{
			cppName: "SkTypeface",
			methods: []callable{
				{cppName: "fontStyle"},
				{cppName: "isBold"},
				{cppName: "isItalic"},
				{cppName: "isFixedPitch"},
				{cppName: "uniqueID"},
				{cppName: "makeClone"},
				{cppName: "unicharsToGlyphs"},
				{cppName: "textToGlyphs"},
				{cppName: "unicharToGlyph"},
				{cppName: "countGlyphs"},
				{cppName: "countTables"},
				{cppName: "getUnitsPerEm"},
				{cppName: "getFamilyName"},
				{cppName: "Equal"},
				{cppName: "MakeEmpty"},
			},
		},
	}
}
