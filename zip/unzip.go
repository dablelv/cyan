package zip

import (
	"archive/zip"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Unzip decompresses a zip file to specified directory.
// Note that the destination directory don't need to specify the trailing path separator.
func Unzip(zipPath, dstDir string) error {
	// Open zip file.
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if err := unzipFile(file, dstDir); err != nil {
			return err
		}
	}
	return nil
}

func unzipFile(file *zip.File, dstDir string) error {
	// Prevent path traversal vulnerability.
	// Such as if the file name is "../../../path/to/file.txt" which will be cleaned to "path/to/file.txt".
	name := strings.TrimPrefix(filepath.Join(string(filepath.Separator), file.Name), string(filepath.Separator))
	filePath := path.Join(dstDir, name)

	// Create the directory of file.
	if file.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// Open the file.
	r, err := file.Open()
	if err != nil {
		return err
	}
	defer r.Close()

	// Create the file.
	w, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer w.Close()

	// Save the decompressed file content.
	_, err = io.Copy(w, r)
	return err
}
