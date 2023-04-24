package file

import (
	"testing"

	"github.com/dablelv/go-huge-util/internal"
)

func TestListDir(t *testing.T) {
	assert := internal.NewAssert(t, "TestListDir")
	filenames, _ := ListDir(".")
	haveFile := len(filenames) > 0
	assert.Equal(true, haveFile)
}

func TestReadLines(t *testing.T) {
	assert := internal.NewAssert(t, "TestReadLines")

	path := "./file_unit_test_tmp.txt"
	f, _ := Create(path)

	_, err := f.WriteString("hello\nworld")
	if err != nil {
		t.Log(err)
	}

	lines, err := ReadLines(path)
	assert.IsNil(err)
	expected := []string{"hello", "world"}
	assert.Equal(expected, lines)

	f.Close()
	if err := Remove(path); err != nil {
		t.Log(err)
	}
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
	if err := Remove(path); err != nil {
		t.Log(err)
	}
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
	if err := Remove(path); err != nil {
		t.Log(err)
	}
}

func TestCreateFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestCreateFile")

	path := "./file_unit_test_tmp.txt"
	err := CreateFile(path)
	assert.IsNil(err)

	if err := Remove(path); err != nil {
		t.Log(err)
	}
}

func TestBytesToFile(t *testing.T) {
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

	if err := Remove(path); err != nil {
		t.Log(err)
	}
}
