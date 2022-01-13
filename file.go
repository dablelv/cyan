package util

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
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

// ListDir lists all the files in the directory
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

// IsPathExist determines whether a file/dir path exists.
// User os.Stat to get the info of target file or dir to check whether exists.
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

// IsDir determines whether a path is a directory
func IsDir(path string) bool {
	if info, err := os.Stat(path); err == nil && info.IsDir() {
		return true
	}
	return false
}

// IsFile determines whether a path is a file.
func IsFile(path string) bool {
	if info, err := os.Stat(path); err == nil && info.Mode().IsRegular() {
		return true
	}
	return false
}

// RemoveFile removes the named file or empty directory.
// https://gist.github.com/novalagung/13c5c8f4d30e0c4bff27
func RemoveFile(path string) error {
	err := os.Remove(path)
	return err
}

// CreateFile creates a new file with the specified file path.
func CreateFile(filePath string) error {
	if exist, _ := IsPathExist(filePath); exist {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}
	file, err := os.Create(filePath)
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
