package time

import (
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

func TestInCloseIntvl(t *testing.T) {
	// 2008-08-08 20:00:00 UTC+8
	ms := int64(1218196800000)

	{
		// target < left
		assert := utest.NewAssert(t, "target<left")
		target := FromMs(ms, LocAsiaShanghai)
		left := FromMs(ms+1, LocAsiaShanghai)
		right := FromMs(ms+2, LocAsiaShanghai)
		b := InCloseIntvl(target, left, right)
		assert.False(b)
	}

	{
		// target > right
		assert := utest.NewAssert(t, "target>left")
		target := FromMs(ms, LocAsiaShanghai)
		left := FromMs(ms-2, LocAsiaShanghai)
		right := FromMs(ms-1, LocAsiaShanghai)
		b := InCloseIntvl(target, left, right)
		assert.False(b)
	}

	{
		// left = target < right
		assert := utest.NewAssert(t, "left=target<right")
		target := FromMs(ms, LocAsiaShanghai)
		left := FromMs(ms, LocAsiaShanghai)
		right := FromMs(ms+1, LocAsiaShanghai)
		b := InCloseIntvl(target, left, right)
		assert.True(b)
	}

	{
		// left < target = right
		assert := utest.NewAssert(t, "left<target=right")
		target := FromMs(ms+1, LocAsiaShanghai)
		left := FromMs(ms, LocAsiaShanghai)
		right := FromMs(ms+1, LocAsiaShanghai)
		b := InCloseIntvl(target, left, right)
		assert.True(b)
	}
}

func TestInLeftCloseIntvl(t *testing.T) {
	// 2008-08-08 20:00:00 UTC+8
	ms := int64(1218196800000)

	{
		// target < left
		assert := utest.NewAssert(t, "target<left")
		target := FromMs(ms, LocAsiaShanghai)
		left := FromMs(ms+1, LocAsiaShanghai)
		right := FromMs(ms+2, LocAsiaShanghai)
		b := InLeftCloseIntvl(target, left, right)
		assert.False(b)
	}

	{
		// target > right
		assert := utest.NewAssert(t, "target>left")
		target := FromMs(ms, LocAsiaShanghai)
		left := FromMs(ms-2, LocAsiaShanghai)
		right := FromMs(ms-1, LocAsiaShanghai)
		b := InLeftCloseIntvl(target, left, right)
		assert.False(b)
	}

	{
		// left = target < right
		assert := utest.NewAssert(t, "left=target<right")
		target := FromMs(ms, LocAsiaShanghai)
		left := FromMs(ms, LocAsiaShanghai)
		right := FromMs(ms+1, LocAsiaShanghai)
		b := InLeftCloseIntvl(target, left, right)
		assert.True(b)
	}

	{
		// left < target = right
		assert := utest.NewAssert(t, "left<target=right")
		target := FromMs(ms+1, LocAsiaShanghai)
		left := FromMs(ms, LocAsiaShanghai)
		right := FromMs(ms+1, LocAsiaShanghai)
		b := InLeftCloseIntvl(target, left, right)
		assert.False(b)
	}
}

func TestInRightCloseIntvl(t *testing.T) {
	// 2008-08-08 20:00:00 UTC+8
	ms := int64(1218196800000)

	{
		// target < left
		assert := utest.NewAssert(t, "target<left")
		target := FromMs(ms, LocAsiaShanghai)
		left := FromMs(ms+1, LocAsiaShanghai)
		right := FromMs(ms+2, LocAsiaShanghai)
		b := InRightCloseIntvl(target, left, right)
		assert.False(b)
	}

	{
		// target > right
		assert := utest.NewAssert(t, "target>left")
		target := FromMs(ms, LocAsiaShanghai)
		left := FromMs(ms-2, LocAsiaShanghai)
		right := FromMs(ms-1, LocAsiaShanghai)
		b := InRightCloseIntvl(target, left, right)
		assert.False(b)
	}

	{
		// left = target < right
		assert := utest.NewAssert(t, "left=target<right")
		target := FromMs(ms, LocAsiaShanghai)
		left := FromMs(ms, LocAsiaShanghai)
		right := FromMs(ms+1, LocAsiaShanghai)
		b := InRightCloseIntvl(target, left, right)
		assert.False(b)
	}

	{
		// left < target = right
		assert := utest.NewAssert(t, "left<target=right")
		target := FromMs(ms+1, LocAsiaShanghai)
		left := FromMs(ms, LocAsiaShanghai)
		right := FromMs(ms+1, LocAsiaShanghai)
		b := InRightCloseIntvl(target, left, right)
		assert.True(b)
	}
}

func TestInOpenIntvl(t *testing.T) {
	// 2008-08-08 20:00:00 UTC+8
	ms := int64(1218196800000)

	{
		// target < left
		assert := utest.NewAssert(t, "target<left")
		target := FromMs(ms, LocAsiaShanghai)
		left := FromMs(ms+1, LocAsiaShanghai)
		right := FromMs(ms+2, LocAsiaShanghai)
		b := InOpenIntvl(target, left, right)
		assert.False(b)
	}

	{
		// target > right
		assert := utest.NewAssert(t, "target>left")
		target := FromMs(ms, LocAsiaShanghai)
		left := FromMs(ms-2, LocAsiaShanghai)
		right := FromMs(ms-1, LocAsiaShanghai)
		b := InOpenIntvl(target, left, right)
		assert.False(b)
	}

	{
		// left = target < right
		assert := utest.NewAssert(t, "left=target<right")
		target := FromMs(ms, LocAsiaShanghai)
		left := FromMs(ms, LocAsiaShanghai)
		right := FromMs(ms+1, LocAsiaShanghai)
		b := InOpenIntvl(target, left, right)
		assert.False(b)
	}

	{
		// left < target = right
		assert := utest.NewAssert(t, "left<target=right")
		target := FromMs(ms+1, LocAsiaShanghai)
		left := FromMs(ms, LocAsiaShanghai)
		right := FromMs(ms+1, LocAsiaShanghai)
		b := InOpenIntvl(target, left, right)
		assert.False(b)
	}

	{
		// left < target < right
		assert := utest.NewAssert(t, "left<target<right")
		target := FromMs(ms, LocAsiaShanghai)
		left := FromMs(ms-1, LocAsiaShanghai)
		right := FromMs(ms+1, LocAsiaShanghai)
		b := InOpenIntvl(target, left, right)
		assert.True(b)
	}
}
