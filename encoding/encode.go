package encoding

import (
	"bytes"
	"io"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// GBK2UTF8 transforms gbk to utf8.
func GBK2UTF8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	return io.ReadAll(reader)
}

// GBK2UTF8Str transforms gbk to utf8 string.
func GBK2UTF8Str(s string) string {
	dst, _ := GBK2UTF8([]byte(s))
	return string(dst)
}

// UTF82GBK transforms utf8 to gbk.
func UTF82GBK(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	return io.ReadAll(reader)
}

// UTF82GBKStr transforms utf8 to gbk string.
func UTF82GBKStr(s string) string {
	dst, _ := UTF82GBK([]byte(s))
	return string(dst)
}
