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
	assert.Equal(t, 1, int(rect.sk.Left))
	assert.Equal(t, 2, int(rect.sk.Top))
	assert.Equal(t, 3, int(rect.sk.Right))
	assert.Equal(t, 10, int(rect.sk.Bottom))
}

func TestFontMgr(t *testing.T) {
	fm := FontMgrRefDefault()
	fmt.Println(fm)
}
