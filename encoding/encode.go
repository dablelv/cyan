package encoding

import (
	"bytes"
	"io"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// GbkToUtf8 transforms GBK to UTF8.
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	return io.ReadAll(reader)
}

// GbkToUtf8Str transforms GBK to UTF8 string.
func GbkToUtf8Str(s string) string {
	dst, _ := GbkToUtf8([]byte(s))
	return string(dst)
}

// Utf8ToGbk transforms utf8 to gbk.
func Utf8ToGbk(b []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(b), simplifiedchinese.GBK.NewEncoder())
	return io.ReadAll(reader)
}

// Utf8ToGbkStr transforms utf8 to gbk string.
func Utf8ToGbkStr(s string) string {
	dst, _ := Utf8ToGbk([]byte(s))
	return string(dst)
}
