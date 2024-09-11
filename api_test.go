package skia

import (
	"fmt"
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
	assert.Equal(t, 1, int(rect.Left))
	assert.Equal(t, 2, int(rect.Top))
	assert.Equal(t, 3, int(rect.Right))
	assert.Equal(t, 10, int(rect.Bottom))
}

func TestFontMgr(t *testing.T) {
	fm := FontMgrRefDefault()
	fmt.Println(fm)
}
