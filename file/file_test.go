package file

import (
	"os"
	"testing"

	"github.com/dablelv/go-huge-util/internal"
)

func TestReadLines(t *testing.T) {
	assert := internal.NewAssert(t, "TestReadLines")

	path := "./file_unit_test_tmp.txt"
	f, err := Create(path)
	assert.IsNil(err)

	_, err = f.WriteString("hello\nworld")
	assert.IsNil(err)

	lines, err := ReadLines(path)
	assert.IsNil(err)
	expected := []string{"hello", "world"}
	assert.Equal(expected, lines)

	f.Close()
	err = os.Remove(path)
	assert.IsNil(err)
}

func TestReadLinesV2(t *testing.T) {
	assert := internal.NewAssert(t, "TestReadLinesV2")

	path := "./file_unit_test_tmp.txt"
	f, _ := Create(path)

	_, err := f.WriteString("hello\nworld")
	if err != nil {
		t.Log(err)
	}

	lines, err := ReadLinesV2(path)
	assert.IsNil(err)
	expected := []string{"hello", "world"}
	assert.Equal(expected, lines)

	f.Close()
	err = os.Remove(path)
	assert.IsNil(err)
}

func TestReadLinesV3(t *testing.T) {
	assert := internal.NewAssert(t, "TestReadLinesV3")

	path := "./file_unit_test_tmp.txt"
	f, _ := Create(path)

	_, err := f.WriteString("hello\nworld")
	if err != nil {
		t.Log(err)
	}

	lines, err := ReadLinesV3(path)
	assert.IsNil(err)
	expected := []string{"hello", "world"}
	assert.Equal(expected, lines)

	f.Close()
	err = os.Remove(path)
	assert.IsNil(err)
}

func TestCreateFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestCreateFile")

	path := "file_unit_test_tmp.txt"
	err := CreateFile(path)
	assert.IsNil(err)

	err = os.Remove(path)
	assert.IsNil(err)
}

func TestBytesToFile_FileToBytes_ClearFile(t *testing.T) {
	assert := internal.NewAssert(t, "BytesToFile and FileToBytes and ClearFile")

	path := "./file_unit_test_tmp.txt"
	err := BytesToFile(path, []byte("hello world"))
	assert.IsNil(err)

	bytes := FileToBytes(path)
	assert.Equal([]byte("hello world"), bytes)

	err = ClearFile(path)
	assert.IsNil(err)
	bytes = FileToBytes(path)
	assert.Equal([]byte(""), bytes)

	err = os.Remove(path)
	assert.IsNil(err)
}
