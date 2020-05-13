package util

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// GbkToUtf8 transform gbk to utf8
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// GbkToUtf8String transform gbk to utf8 string
func GbkToUtf8String(s string) string {
	if dst, err := GbkToUtf8([]byte(s)); err == nil {
		return string(dst)
	}
	return ""
}

// Utf8ToGbk transform utf8 to gbk
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// Utf8ToGbkString transform utf8 to gbk string
func Utf8ToGbkString(s string) string {
	if dst, err := Utf8ToGbk([]byte(s)); err == nil {
		return string(dst)
	}
	return ""
}