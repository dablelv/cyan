package file

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

const dir = "unit_test_dir"

func TestCreateFile(t *testing.T) {
	assert := utest.NewAssert(t, "TestCreateFile")

	path := "file_unit_test_tmp.txt"
	err := CreateFile(path)
	assert.IsNil(err)

	// already exist.
	err = CreateFile(path)
	assert.IsNil(err)

	err = os.Remove(path)
	assert.IsNil(err)

	// The specified path is invalid.
	err = CreateFile("//file_unit_test_tmp.txt")
	assert.IsNotNil(err)
}

func TestBytesToFile_FileToBytes_ClearFile(t *testing.T) {
	assert := utest.NewAssert(t, "BytesToFile and FileToBytes and ClearFile")

	path := filepath.Join(dir, "foo.txt")
	err := BytesToFile(path, []byte("hello world"))
	assert.IsNil(err)

	bytes := FileToBytes(path)
	assert.Equal([]byte("hello world"), bytes)

	err = ClearFile(path)
	assert.IsNil(err)
	bytes = FileToBytes(path)
	assert.Equal([]byte(""), bytes)

	err = ClearFile("file_not_exist.txt")
	assert.IsNotNil(err)

	err = os.Remove(path)
	assert.IsNil(err)
}
