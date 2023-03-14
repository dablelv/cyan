package file

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// ReadLines reads all lines of the specified file.
func ReadLines(path string) ([]string, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	var lines []string
	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		lineCount++
	}

	if scanner.Err() == bufio.ErrTooLong {
		panic(scanner.Err())
	}
	return lines, lineCount, scanner.Err()
}

// ReadLinesV2 reads all lines of the specified file.
func ReadLinesV2(path string) ([]string, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	var lines []string
	lineCount := 0
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		lines = append(lines, line)
		lineCount++
		if err == io.EOF {
			return lines, lineCount, nil
		}
		if err != nil {
			return lines, lineCount, err
		}
	}
}

// ListDir lists all the file or directory names in the specified directory.
// Note that ListDir don't traverse recursively.
func ListDir(dirname string) ([]string, error) {
	infos, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}
	names := make([]string, len(infos))
	for i, info := range infos {
		names[i] = info.Name()
	}
	return names, nil
}

// IsExist checks whether a file/dir exists.
// Use os.Stat to get the info of the target file or dir to check whether exists.
// If os.Stat returns nil err, the target exists.
// If os.Stat returns a os.ErrNotExist err, the target does not exist.
// If the error returned is another type, the target is uncertain whether exists.
func IsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// IsDir checks whether a path is a directory.
// If the path is a symbolic link will follow it.
func IsDir(path string) bool {
	if info, err := os.Stat(path); err == nil && info.IsDir() {
		return true
	}
	return false
}

// IsDirE checks whether a path is a directory with error.
// If the path is a symbolic link will follow it.
func IsDirE(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil && info.IsDir() {
		return true, nil
	}
	return false, err
}

// IsFile checks whether a path is a file.
// If the path is a symbolic link will follow it.
func IsFile(path string) bool {
	if info, err := os.Stat(path); err == nil && info.Mode().IsRegular() {
		return true
	}
	return false
}

// IsFileE checks whether a path is a file with error.
// If the path is a symbolic link will follow it.
func IsFileE(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil && info.Mode().IsRegular() {
		return true, nil
	}
	return false, err
}

// IsSymlink checks a file whether is a symbolic link on Linux.
// Note that this doesn't work for the shortcut file on windows.
// If you want to check a file whether is a shortcut file on Windows please use IsShortcut function.
func IsSymlink(path string) bool {
	if info, err := os.Lstat(path); err == nil && info.Mode()&os.ModeSymlink != 0 {
		return true
	}
	return false
}

// IsSymlinkE checks a file whether is a symbolic link on Linux.
// Note that this doesn't work for the shortcut file on windows.
// If you want to check a file whether is a shortcut file on Windows please use IsShortcut function.
func IsSymlinkE(path string) (bool, error) {
	info, err := os.Lstat(path)
	if err == nil && info.Mode()&os.ModeSymlink != 0 {
		return true, nil
	}
	return false, err
}

// IsShortcut checks a file whether is a shortcut on Windows.
func IsShortcut(path string) bool {
	ext := filepath.Ext(path)
	if ext == ".lnk" {
		return true
	}
	return false
}

// RemoveFile removes the named file or empty directory.
// If there is an error, it will be of type *PathError.
func RemoveFile(path string) error {
	return os.Remove(path)
}

// Create creates or truncates the target file specified by path.
// If the parent directory does not exist, it will be created with mode os.ModePerm.
// If the file does not exist, it is created with mode 0666.
// If successful, methods on the returned file can be used for I/O; the associated file descriptor has mode O_RDWR.
func Create(path string) (*os.File, error) {
	exist, err := IsExist(path)
	if err != nil {
		return nil, err
	}
	if exist {
		return os.Create(path)
	}
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return nil, err
	}
	return os.Create(path)
}

// CreateFile creates a file specified by path.
func CreateFile(filePath string) error {
	file, err := Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

// FileToBytes serialize the file to bytes.
func FileToBytes(path string) []byte {
	byteStream, _ := ioutil.ReadFile(path)
	return byteStream
}

// BytesToFile writes data to a file.
// If the file does not exist it will be created with permission mode 0644.
func BytesToFile(path string, data []byte) error {
	exist, _ := IsExist(path)
	if !exist {
		if err := CreateFile(path); err != nil {
			return err
		}
	}
	return ioutil.WriteFile(path, data, 0644)
}

// GetDirAllEntryPaths gets all the file or dir paths in the specified directory recursively.
// If the incl is true result will include current dir.
// Note that GetDirAllEntryPaths won't follow symlink if the subdir is a symbolic link.
func GetDirAllEntryPaths(dirname string, incl bool) ([]string, error) {
	// Remove the trailing path separator if dirname has.
	dirname = strings.TrimSuffix(dirname, string(os.PathSeparator))

	infos, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	paths := make([]string, 0, len(infos))
	// Include current dir.
	if incl {
		paths = append(paths, dirname)
	}

	for _, info := range infos {
		path := dirname + string(os.PathSeparator) + info.Name()
		if info.IsDir() {
			tmp, err := GetDirAllEntryPaths(path, incl)
			if err != nil {
				return nil, err
			}
			paths = append(paths, tmp...)
			continue
		}
		paths = append(paths, path)
	}
	return paths, nil
}

// GetDirAllEntryPathsFollowSymlink gets all the file or dir paths in the specified directory recursively.
// If the incl is true result will include current dir.
func GetDirAllEntryPathsFollowSymlink(dirname string, incl bool) ([]string, error) {
	// Remove the trailing path separator if dirname has.
	dirname = strings.TrimSuffix(dirname, string(os.PathSeparator))

	infos, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	paths := make([]string, 0, len(infos))
	// Include current dir.
	if incl {
		paths = append(paths, dirname)
	}

	for _, info := range infos {
		path := dirname + string(os.PathSeparator) + info.Name()
		realInfo, err := os.Stat(path)
		if err != nil {
			return nil, err
		}
		if realInfo.IsDir() {
			tmp, err := GetDirAllEntryPathsFollowSymlink(path, incl)
			if err != nil {
				return nil, err
			}
			paths = append(paths, tmp...)
			continue
		}
		paths = append(paths, path)
	}
	return paths, nil
}

// ClearFile clears a file content.
func ClearFile(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
