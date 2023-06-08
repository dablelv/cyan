package slice

import (
	"testing"

	"github.com/dablelv/go-huge-util/internal"
)

func TestUpdateInt(t *testing.T) {
	assert := internal.NewAssert(t, "TestUpdateInt")

	got := UpdateInt([]int{1}, 0, 2)
	assert.Equal([]int{2}, got)
}

func TestUpdateInt8(t *testing.T) {
	assert := internal.NewAssert(t, "TestUpdateInt8")

	got := UpdateInt8([]int8{1}, 0, 2)
	assert.Equal([]int8{2}, got)
}

func TestUpdateInt16(t *testing.T) {
	assert := internal.NewAssert(t, "TestUpdateInt16")

	got := UpdateInt16([]int16{1}, 0, 2)
	assert.Equal([]int16{2}, got)
}

func TestUpdateInt32(t *testing.T) {
	assert := internal.NewAssert(t, "TestUpdateInt32")

	got := UpdateInt32([]int32{1}, 0, 2)
	assert.Equal([]int32{2}, got)
}

func TestUpdateInt64(t *testing.T) {
	assert := internal.NewAssert(t, "TestUpdateInt64")

	got := UpdateInt64([]int64{1}, 0, 2)
	assert.Equal([]int64{2}, got)
}

func TestUpdateUint(t *testing.T) {
	assert := internal.NewAssert(t, "TestUpdateUint")

	got := UpdateUint([]uint{1}, 0, 2)
	assert.Equal([]uint{2}, got)
}

func TestUpdateUint8(t *testing.T) {
	assert := internal.NewAssert(t, "TestUpdateUint8")

	got := UpdateUint8([]uint8{1}, 0, 2)
	assert.Equal([]uint8{2}, got)
}

func TestUpdateUint16(t *testing.T) {
	assert := internal.NewAssert(t, "TestUpdateUint16")

	got := UpdateUint16([]uint16{1}, 0, 2)
	assert.Equal([]uint16{2}, got)
}

func TestUpdateUint32(t *testing.T) {
	assert := internal.NewAssert(t, "TestUpdateUint32")

	got := UpdateUint32([]uint32{1}, 0, 2)
	assert.Equal([]uint32{2}, got)
}

func TestUpdateUint64(t *testing.T) {
	assert := internal.NewAssert(t, "TestUpdateUint64")

	got := UpdateUint64([]uint64{1}, 0, 2)
	assert.Equal([]uint64{2}, got)
}

func TestUpdateStr(t *testing.T) {
	assert := internal.NewAssert(t, "TestUpdateStr")

	got := UpdateStr([]string{"foo"}, 0, "bar")
	assert.Equal([]string{"bar"}, got)
}

func TestUpdateE(t *testing.T) {
	assert := internal.NewAssert(t, "TestUpdateE")

	// not slice
	_, err := UpdateE(1, 0, 2)
	assert.IsNotNil(err)

	// index out of range
	_, err = UpdateE([]int{1}, 1, 2)
	assert.IsNotNil(err)

	// value type invalid
	_, err = UpdateE([]uint64{1}, 0, 2)
	assert.IsNotNil(err)

	_, err = UpdateE([]int{1}, 0, 2)
	assert.IsNil(err)
}
