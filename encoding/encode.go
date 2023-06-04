package encoding

import (
	"bytes"
	"io"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// GBKToUTF8 transforms GBK to UTF8.
func GBKToUTF8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	return io.ReadAll(reader)
}

// GBKToUTF8Str transforms GBK to UTF8 string.
func GBKToUTF8Str(s string) string {
	dst, _ := GBKToUTF8([]byte(s))
	return string(dst)
}

// UTF82GBK transforms utf8 to gbk.
func UTF8ToGBK(b []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(b), simplifiedchinese.GBK.NewEncoder())
	return io.ReadAll(reader)
}

// UTF82GBKStr transforms utf8 to gbk string.
func UTF8ToGBKStr(s string) string {
	dst, _ := UTF8ToGBK([]byte(s))
	return string(dst)
}
