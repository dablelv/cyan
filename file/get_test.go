package file

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/dablelv/cyan/internal"
)

func TestReadLines(t *testing.T) {
	assert := internal.NewAssert(t, "TestReadLines")

	path := filepath.Join(dir, "foo.txt")
	f, err := Create(path)
	assert.IsNil(err)

	_, err = f.WriteString("hello\nworld")
	assert.IsNil(err)

	lines, err := ReadLines(path)
	assert.IsNil(err)
	expected := []string{"hello", "world"}
	assert.Equal(expected, lines)

	f.Close()
	err = os.RemoveAll(dir)
	assert.IsNil(err)

	_, err = ReadLines("file_not_exist.txt")
	assert.IsNotNil(err)
}

func TestReadLinesV2(t *testing.T) {
	assert := internal.NewAssert(t, "TestReadLinesV2")

	path := filepath.Join(dir, "foo.txt")
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
	err = os.RemoveAll(dir)
	assert.IsNil(err)

	_, err = ReadLinesV2("file_not_exist.txt")
	assert.IsNotNil(err)
}

func TestReadLinesV3(t *testing.T) {
	assert := internal.NewAssert(t, "TestReadLinesV3")

	path := filepath.Join(dir, "foo.txt")
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
	err = os.RemoveAll(dir)
	assert.IsNil(err)

	_, err = ReadLinesV3("file_not_exist")
	assert.IsNotNil(err)
}

func TestListDir(t *testing.T) {
	assert := internal.NewAssert(t, "TestListDir")

	path := filepath.Join(dir, "foo.txt")
	err := CreateFile(path)
	assert.IsNil(err)

	names, err := ListDir(dir)
	assert.IsNil(err)
	assert.Equal([]string{"foo.txt"}, names)

	err = os.RemoveAll(dir)
	assert.IsNil(err)

	// Dir not exist.
	_, err = ListDir("not_exist_dir")
	assert.IsNotNil(err)
}

func TestListFilenames(t *testing.T) {
	assert := internal.NewAssert(t, "TestListFilenames")

	err := CreateFile(filepath.Join(dir, "a.txt"))
	assert.IsNil(err)
	err = CreateFile(filepath.Join(dir, "b.txt"))
	assert.IsNil(err)

	names, err := ListFilenames(dir)
	assert.IsNil(err)
	assert.Equal([]string{"a.txt", "b.txt"}, names)

	err = os.RemoveAll(dir)
	assert.IsNil(err)

	// Dir not exist.
	_, err = ListFilenames("dir_not_exist")
	assert.IsNotNil(err)
}

func TestFileMD5(t *testing.T) {
	assert := internal.NewAssert(t, "TestFileMD5")

	path := filepath.Join(dir, "foo.txt")
	err := CreateFile(path)
	assert.IsNil(err)

	md5, err := FileMD5(path)
	assert.IsNil(err)
	assert.Equal("d41d8cd98f00b204e9800998ecf8427e", md5)

	err = os.Remove(path)
	assert.IsNil(err)

	// File not exist.
	_, err = FileMD5("not_exist_file")
	assert.IsNotNil(err)
}

func TestFileMD5Reader(t *testing.T) {
	assert := internal.NewAssert(t, "TestFileMD5Reader")

	path := filepath.Join(dir, "foo.txt")
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

	path := filepath.Join(dir, "foo.txt")
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

	// file not exist.
	_, err = FileSize("file_not_exist")
	assert.IsNotNil(err)
}

func TestListDirEntryPaths(t *testing.T) {
	assert := internal.NewAssert(t, "TestListDirEntryPaths")

	path1 := filepath.Join(dir, "a.txt")
	err := CreateFile(path1)
	assert.IsNil(err)
	path2 := filepath.Join(dir, "b.txt")
	err = CreateFile(path2)
	assert.IsNil(err)
	path3 := filepath.Join(dir, "subdir")
	err = os.MkdirAll(path3, os.ModePerm)
	assert.IsNil(err)

	paths, err := ListDirEntryPaths(dir, true)
	assert.IsNil(err)
	assert.Equal([]string{dir, path1, path2, path3}, paths)

	err = os.RemoveAll(dir)
	assert.IsNil(err)

	_, err = ListDirEntryPaths("dir_not_exist", true)
	assert.IsNotNil(err)
}

func TestListDirEntryPathsSymlink(t *testing.T) {
	assert := internal.NewAssert(t, "TestListDirEntryPathsSymlink")

	path1 := filepath.Join(dir, "a.txt")
	err := CreateFile(path1)
	assert.IsNil(err)
	path2 := filepath.Join(dir, "b.txt")
	err = CreateFile(path2)
	assert.IsNil(err)
	path3 := filepath.Join(dir, "subdir")
	err = os.MkdirAll(path3, os.ModePerm)
	assert.IsNil(err)

	paths, err := ListDirEntryPathsSymlink(dir, true)
	assert.IsNil(err)
	assert.Equal([]string{dir, path1, path2, path3}, paths)

	err = os.RemoveAll(dir)
	assert.IsNil(err)

	_, err = ListDirEntryPathsSymlink("dir_not_exist", true)
	assert.IsNotNil(err)
}
