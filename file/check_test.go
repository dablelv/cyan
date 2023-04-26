package file

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/dablelv/go-huge-util/internal"
)

func TestIsExist(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsExist")

	path := filepath.Join("unit_test_dir", "a.txt")
	err := CreateFile(path)
	assert.IsNil(err)

	exist, err := IsExist(path)
	assert.IsNil(err)
	assert.Equal(true, exist)

	err = os.RemoveAll("unit_test_dir")
	assert.IsNil(err)
}

func TestIsDir(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsDir")

	dir := "unit_test_dir"
	err := os.Mkdir(dir, os.ModePerm)
	assert.IsNil(err)

	is := IsDir(dir)
	assert.Equal(true, is)

	// not found
	is = IsDir(filepath.Join(dir, dir))
	assert.Equal(false, is)

	err = os.RemoveAll("unit_test_dir")
	assert.IsNil(err)
}

func TestIsFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsFile")

	dir := "unit_test_dir"
	path := filepath.Join(dir, "a.txt")
	err := CreateFile(path)
	assert.IsNil(err)

	is := IsFile(path)
	assert.Equal(true, is)

	// not found
	is = IsFile(filepath.Join(dir, "b.txt"))
	assert.Equal(false, is)

	err = os.RemoveAll("unit_test_dir")
	assert.IsNil(err)
}

func TestIsShortcut(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsFile")

	path := "unit_test_dir/a.lnk"
	err := CreateFile(path)
	assert.IsNil(err)

	is := IsShortcut(path)
	assert.Equal(true, is)

	err = os.Remove(path)
	assert.IsNil(err)
}
