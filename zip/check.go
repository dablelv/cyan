package zip

import (
	"bytes"
	"io"
	"os"
)

// IsZipFile reports file whether is a zip file.
// About the zip file format can refer to https://docs.fileformat.com/compression/zip/.
func IsZipFile(filepath string) (bool, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return false, err
	}
	defer f.Close()

	buf := make([]byte, 4)
	_, err = f.Read(buf)
	if err == io.EOF {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return bytes.Equal(buf, []byte("PK\x03\x04")), nil
}
