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

	err = os.RemoveAll(dir)
	assert.IsNil(err)
}

func TestIsDir(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsDir")

	err := os.Mkdir(dir, os.ModePerm)
	assert.IsNil(err)

	is := IsDir(dir)
	assert.Equal(true, is)

	// not found
	is = IsDir(filepath.Join(dir, dir))
	assert.Equal(false, is)

	err = os.RemoveAll(dir)
	assert.IsNil(err)
}

func TestIsFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsFile")

	path := filepath.Join(dir, "a.txt")
	err := CreateFile(path)
	assert.IsNil(err)

	is := IsFile(path)
	assert.Equal(true, is)

	// not found
	is = IsFile(filepath.Join(dir, "b.txt"))
	assert.Equal(false, is)

	err = os.RemoveAll(dir)
	assert.IsNil(err)
}

func TestIsShortcut(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsFile")

	path := filepath.Join(dir, "a.lnk")
	err := CreateFile(path)
	assert.IsNil(err)
	is := IsShortcut(path)
	assert.Equal(true, is)

	path = filepath.Join(dir, "a.exe")
	err = CreateFile(path)
	assert.IsNil(err)
	is = IsShortcut(path)
	assert.Equal(false, is)

	err = os.RemoveAll(dir)
	assert.IsNil(err)
}

func TestIsSymlink(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsSymlink")
	is := IsSymlink("file_not_exist")
	assert.Equal(false, is)
}

func TestIsSymlinkE(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsSymlinkE")
	_, err := IsSymlinkE("file_not_exist")
	assert.IsNotNil(err)
}
