package encoding

import (
	"bytes"
	"io"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// GBK2UTF8 transform gbk to utf8.
func GBK2UTF8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// GBK2UTF8Str transform gbk to utf8 string.
func GBK2UTF8Str(s string) string {
	if dst, err := GBK2UTF8([]byte(s)); err == nil {
		return string(dst)
	}
	return ""
}

// UTF82GBK transform utf8 to gbk.
func UTF82GBK(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// UTF82GBKStr transform utf8 to gbk string.
func UTF82GBKStr(s string) string {
	if dst, err := UTF82GBK([]byte(s)); err == nil {
		return string(dst)
	}
	return ""
}
