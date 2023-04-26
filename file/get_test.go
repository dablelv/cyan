package file

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/dablelv/go-huge-util/internal"
)

func TestListDir(t *testing.T) {
	assert := internal.NewAssert(t, "TestListDir")

	path := "./file_unit_test/tmp.txt"
	err := CreateFile(path)
	assert.IsNil(err)

	names, err := ListDir("./file_unit_test/")
	assert.IsNil(err)
	assert.Equal([]string{"tmp.txt"}, names)

	err = os.RemoveAll("./file_unit_test")
	assert.IsNil(err)
}

func TestListFilenames(t *testing.T) {
	assert := internal.NewAssert(t, "TestListFilenames")

	err := CreateFile("./file_unit_test/a.txt")
	assert.IsNil(err)
	err = CreateFile("./file_unit_test/b.txt")
	assert.IsNil(err)

	names, err := ListFilenames("./file_unit_test/")
	assert.IsNil(err)
	assert.Equal([]string{"a.txt", "b.txt"}, names)

	err = os.RemoveAll("./file_unit_test")
	assert.IsNil(err)
}

func TestFileMD5(t *testing.T) {
	assert := internal.NewAssert(t, "TestFileMD5")

	path := "./file_unit_test_tmp.txt"
	err := CreateFile(path)
	assert.IsNil(err)

	md5, err := FileMD5(path)
	assert.IsNil(err)
	assert.Equal("d41d8cd98f00b204e9800998ecf8427e", md5)

	err = os.Remove(path)
	assert.IsNil(err)
}

func TestFileMD5Reader(t *testing.T) {
	assert := internal.NewAssert(t, "TestFileMD5Reader")

	path := "./file_unit_test_tmp.txt"
	f, err := Create(path)
	assert.IsNil(err)

	md5, err := FileMD5Reader(f)
	assert.IsNil(err)
	assert.Equal("d41d8cd98f00b204e9800998ecf8427e", md5)

	f.Close()
	err = os.Remove(path)
	assert.IsNil(err)
}

func TestFileSize(t *testing.T) {
	assert := internal.NewAssert(t, "TestFileSize")

	path := "./file_unit_test_tmp.txt"
	f, err := Create(path)
	assert.IsNil(err)

	_, err = f.WriteString("hello world")
	assert.IsNil(err)

	size, err := FileSize(path)
	assert.IsNil(err)
	assert.Equal(int64(11), size)

	f.Close()
	err = os.Remove(path)
	assert.IsNil(err)
}

func TestListDirEntryPaths(t *testing.T) {
	assert := internal.NewAssert(t, "TestListDirEntryPaths")

	path1 := filepath.Join("file_unit_test_dir", "a.txt")
	err := CreateFile(path1)
	assert.IsNil(err)
	path2 := filepath.Join("file_unit_test_dir", "b.txt")
	err = CreateFile(path2)
	assert.IsNil(err)

	dir := "file_unit_test_dir"
	paths, err := ListDirEntryPaths(dir, false)
	assert.IsNil(err)
	assert.Equal([]string{path1, path2}, paths)

	err = os.RemoveAll(dir)
	assert.IsNil(err)
}

func TestListDirEntryPathsSymlink(t *testing.T) {
	assert := internal.NewAssert(t, "TestListDirEntryPathsSymlink")

	path1 := filepath.Join("file_unit_test_dir", "a.txt")
	err := CreateFile(path1)
	assert.IsNil(err)
	path2 := filepath.Join("file_unit_test_dir", "b.txt")
	err = CreateFile(path2)
	assert.IsNil(err)

	dir := "file_unit_test_dir"
	paths, err := ListDirEntryPathsSymlink(dir, true)
	assert.IsNil(err)
	assert.Equal([]string{dir, path1, path2}, paths)

	err = os.RemoveAll(dir)
	assert.IsNil(err)
}
