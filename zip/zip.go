package zip

import (
	"archive/zip"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/dablelv/go-huge-util/file"
)

//
// How to implement zip and unzip please refer to https://dablelv.blog.csdn.net/article/details/122441250.
//

// Zip compresses the specified files or dirs to zip archive.
// If a path is a dir don't need to specify the trailing path separator.
// For example calling Zip("archive.zip", "dir", "csv/baz.csv") will get archive.zip and the content of which is
// baz.csv
// dir
// ├── bar.txt
// └── foo.txt
// Note that if a file is a symbolic link on Linux it will be skipped.
// If you want to follow a symbolic link please use the function ZipFollowSymlink.
func Zip(zipPath string, paths ...string) error {
	// Create zip file and it's parent dir.
	if err := os.MkdirAll(filepath.Dir(zipPath), os.ModePerm); err != nil {
		return err
	}
	archive, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer archive.Close()

	// New zip writer.
	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	// Traverse the file or directory.
	for _, rootPath := range paths {
		// Remove the trailing path separator if path is a directory.
		rootPath = strings.TrimSuffix(rootPath, string(os.PathSeparator))

		// Visit all the files or directories in the tree.
		err = filepath.Walk(rootPath, walkFunc(rootPath, zipWriter))
		if err != nil {
			return err
		}
	}
	return nil
}

func walkFunc(rootPath string, zipWriter *zip.Writer) filepath.WalkFunc {
	return func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// If a file is a symbolic link it will be skipped.
		if info.Mode()&os.ModeSymlink != 0 {
			return nil
		}

		// Create a local file header.
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Set compression method.
		header.Method = zip.Deflate

		// Set relative path of a file as the header name.
		header.Name, err = filepath.Rel(filepath.Dir(rootPath), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += string(os.PathSeparator)
		}

		// Create writer for the file header and save content of the file.
		headerWriter, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(headerWriter, f)
		return err
	}
}

// ZipFollowSymlink compresses the specified files or dirs to zip archive.
// If the specified files or dirs is a symbolic link ZipFollowSymlink will follow it.
// Note that the symbolic link need to avoid loops.
func ZipFollowSymlink(zipPath string, paths ...string) error {
	// Create zip file and it's parent dir.
	if err := os.MkdirAll(filepath.Dir(zipPath), os.ModePerm); err != nil {
		return err
	}
	archive, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer archive.Close()

	// New zip writer.
	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	// Get all the file or directory paths.
	var allFilePaths []string
	pathToRoot := make(map[string]string)
	for _, path := range paths {
		// If the path is a dir or symlink to dir, get all files in it.
		info, err := os.Stat(path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			// Remove the trailing path separator if path is a directory.
			path = strings.TrimSuffix(path, string(os.PathSeparator))
			filePaths, err := file.GetDirAllEntryPathsFollowSymlink(path, true)
			if err != nil {
				return err
			}
			allFilePaths = append(allFilePaths, filePaths...)
			for _, p := range filePaths {
				pathToRoot[p] = path
			}
			continue
		}
		allFilePaths = append(allFilePaths, path)
		pathToRoot[path] = path
	}

	// Traverse all the file or directory.
	for _, path := range allFilePaths {
		info, err := os.Stat(path)
		if err != nil {
			return err
		}

		// Create a local file header.
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Set compression method.
		header.Method = zip.Deflate

		// Set relative path of a file as the header name.
		header.Name, err = filepath.Rel(filepath.Dir(pathToRoot[path]), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += string(os.PathSeparator)
		}

		// Create writer for the file header and save content of the file.
		headerWriter, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// If file is a directory.
		if info.IsDir() {
			continue
		}

		// If file is a file or symlink to file.
		realPath, err := filepath.EvalSymlinks(path)
		if err != nil {
			return err
		}
		f, err := os.Open(realPath)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(headerWriter, f)
		if err != nil {
			return err
		}
	}
	return nil
}
