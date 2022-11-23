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

// ListDir lists all the files in the directory.
func ListDir(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	filenames := make([]string, len(files))
	for idx, file := range files {
		filenames[idx] = file.Name()
	}
	return filenames
}

// IsPathExist checks whether a file/dir exists.
// Use os.Stat to get the info of the target file or dir to check whether exists.
// If os.Stat returns nil err, the target exists.
// If os.Stat returns a os.ErrNotExist err, the target does not exist.
// If the error returned is another type, the target is uncertain whether exists.
func IsPathExist(path string) (bool, error) {
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

// IsSymlink checks a file whether is a symbolic link.
// Note that this doesn't work for the shortcut file on windows.
func IsSymlink(path string) bool {
	if info, err := os.Lstat(path); err == nil && info.Mode()&os.ModeSymlink != 0 {
		return true
	}
	return false
}

// IsSymlinkE checks a file whether is a symbolic link.
// Note that this doesn't work for the shortcut file on windows.
func IsSymlinkE(path string) (bool, error) {
	info, err := os.Lstat(path)
	if err == nil && info.Mode()&os.ModeSymlink != 0 {
		return true, nil
	}
	return false, err
}

// RemoveFile removes the named file or empty directory.
// https://gist.github.com/novalagung/13c5c8f4d30e0c4bff27
func RemoveFile(path string) error {
	err := os.Remove(path)
	return err
}

// Create creates or truncates the target file specified by path.
// If the parent directory does not exist, it will be created with mode os.ModePerm.is cr truncated.
// If the file does not exist, it is created with mode 0666.
// If successful, methods on the returned File can be used for I/O; the associated file descriptor has mode O_RDWR.
func Create(filePath string) (*os.File, error) {
	if exist, err := IsPathExist(filePath); err != nil {
		return nil, err
	} else if exist {
		return os.Create(filePath)
	}
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return nil, err
	}
	return os.Create(filePath)
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

// BytesToFile writes data to a file. If the file does not exist it will be created with permission mode 0644.
func BytesToFile(filePath string, data []byte) error {
	exist, _ := IsPathExist(filePath)
	if !exist {
		if err := CreateFile(filePath); err != nil {
			return err
		}
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

// GetDirAllFilePaths gets dir's all file paths recusively including it self.
// If the dir or the sub file is symbolic link will follow it.
// Note that the symbolic link need to avoid loops.
func GetDirAllFilePaths(dirPath string) ([]string, error) {
	result := []string{}

	dirPath = strings.TrimSuffix(dirPath, string(os.PathSeparator))
	result = append(result, dirPath)

	// Read dir.
	fis, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return result, err
	}

	// Recursively traverse all sub files.
	for _, fi := range fis {
		fullPath := dirPath + string(os.PathSeparator) + fi.Name()
		info, err := os.Stat(fullPath)
		if err != nil {
			return nil, err
		}
		if info.IsDir() {
			tmp, err := GetDirAllFilePaths(fullPath)
			if err != nil {
				return nil, err
			}
			result = append(result, tmp...)
		} else {
			result = append(result, fullPath)
		}
	}
	return result, nil
}
