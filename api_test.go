package skia

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnumValue(t *testing.T) {
	assert.Equal(t, PaintJoin(2), PaintJoinBevel)
}

func TestCtorDtor(t *testing.T) {
	paint := NewPaint()
	paint2 := NewPaintCopy(paint)
	paint.Delete()
	paint2.Delete()
}

func TestClassStaticMethod(t *testing.T) {
	bm := NewBitmap()
	assert.False(t, BitmapComputeIsOpaque(bm))
}

func TestMethodReturningNonPointerRecord(t *testing.T) {
	rect := IRectMakeLTRB(1, 2, 3, 10)
	assert.Equal(t, 1, int(rect.Left()))
	assert.Equal(t, 2, int(rect.Top()))
	assert.Equal(t, 3, int(rect.Right()))
	assert.Equal(t, 10, int(rect.Bottom()))
}

func TestFontMgr(t *testing.T) {
	fontMgr := FontMgrRefDefault()
	typeface := fontMgr.MatchFamilyStyle("monospace", FontStyleNormal())
	assert.NotNil(t, typeface.sk)
}

func TestPointerArgument(t *testing.T) {
	typeface := FontMgrRefDefault().MatchFamilyStyle("sans-serif", FontStyleNormal())
	font := NewFontTypefaceSize(typeface, 22)
	var metrics FontMetrics
	lineSpacing := font.GetMetrics(&metrics)
	assert.NotZero(t, lineSpacing)
	assert.NotZero(t, metrics.fAscent)

}

func TestFontGlyphIDs(t *testing.T) {
	typeface := FontMgrRefDefault().MatchFamilyStyle("sans-serif", FontStyleNormal())
	font := NewFontTypefaceSize(typeface, 22)

	chars := []int32{'A', 'z', '?'}
	glyphs := make([]uint16, len(chars))
	font.UnicharsToGlyphs(chars, int32(len(chars)), glyphs)
	for i, char := range chars {
		assert.NotZero(t, glyphs[i])
		assert.Equal(t, font.UnicharToGlyph(char), glyphs[i])
	}
}

func TestString(t *testing.T) {
	s := "qwerty"
	sk := NewString(s)
	assert.Equal(t, s, sk.Data())
}
