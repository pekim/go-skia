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
