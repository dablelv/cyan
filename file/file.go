package file

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ReadLines reads all lines of the file.
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

// ReadLinesV2 reads all lines of the file.
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

// ClearFile clears a file content.
func ClearFile(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
