package zip

import (
	"archive/zip"
	"io"
	"os"
	"path"
	"path/filepath"
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
	// create the directory of file
	filePath := path.Join(dstDir, file.Name)
	if file.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// open the file
	r, err := file.Open()
	if err != nil {
		return err
	}
	defer r.Close()

	// create the file
	w, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer w.Close()

	// save the decompressed file content
	_, err = io.Copy(w, r)
	return err
}
