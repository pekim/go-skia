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

	assert.False(t, paint.IsNil())
	paint.Delete()
	assert.True(t, paint.IsNil())

	assert.False(t, paint2.IsNil())
	paint2.Delete()
	assert.True(t, paint2.IsNil())
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
	fontMgr := FontMgrDefault()
	typeface := fontMgr.MatchFamilyStyle("monospace", FontStyleNormal())
	assert.NotNil(t, typeface.sk)
}

func TestPointerArgument(t *testing.T) {
	typeface := FontMgrDefault().MatchFamilyStyle("sans-serif", FontStyleNormal())
	font := NewFontTypefaceSize(typeface, 22)
	var metrics FontMetrics
	lineSpacing := font.GetMetrics(&metrics)
	assert.NotZero(t, lineSpacing)
	assert.NotZero(t, metrics.fAscent)

}

func TestFontGlyphIDs(t *testing.T) {
	typeface := FontMgrDefault().MatchFamilyStyle("sans-serif", FontStyleNormal())
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

func TestOutParams(t *testing.T) {
	typeface := FontMgrDefault().MatchFamilyStyle("sans-serif", FontStyleNormal())
	font := NewFontTypefaceSize(typeface, 12)
	var bounds Rect
	text := "Qwerty"

	advance := font.MeasureText([]byte(text), uint32(len(text)), TextEncodingUTF8, &bounds)
	assert.NotZero(t, advance)
	assert.NotZero(t, bounds.Width())
	assert.NotZero(t, bounds.Height())
}

func TestAllStructsSizesMatch(_ *testing.T) {
	assertAllStructsSizesMatch()
}
