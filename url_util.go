package util

import (
	"errors"
	"fmt"
	"strings"
	"net/url"
)

// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name and key="page" will get "1"
func RawURLGetParam(rawUrl, key string) (string, error) {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	m := stUrl.Query()
	if v, ok := m[key]; ok && len(v) > 0 {
		return v[0], nil
	}
	return "", errors.New("no param")
}

// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name and key=page will get [1 2]
func RawURLGetParams(rawUrl, key string) ([]string, error) {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}

	m := stUrl.Query()
	if v, ok := m[key]; ok {
		return v, nil
	}
	return nil, errors.New("no param")
}

// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name will get map[boardID:[520] page:[1 2]]
func RawURLGetAllParams(rawUrl string) (map[string][]string, error) {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}

	m := stUrl.Query()
	return m, nil
}

// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name and key=page value=3
// will get http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2&page=3#name
func RawURLAddParam(sUrl, key, value string) string {
	stUrl, err := url.Parse(sUrl)
	if err != nil{
		return sUrl
	}

	m := stUrl.Query()
	m.Add(key, value)
	stUrl.RawQuery = m.Encode()
	return stUrl.String()
}

// e.g. if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name and key=page
// will get http://www.aspxfans.com:8080/news/index.asp?boardID=520#name
func RawURLDelParam(sUrl, key string) string {
	stUrl, err := url.Parse(sUrl)
    if err != nil{
        return sUrl
	}

	m := stUrl.Query()
	m.Del(key)
	stUrl.RawQuery = m.Encode()
	return stUrl.String()
}

// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name and key=page value=3
// will get http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=3#name
func RawURLSetParam(sUrl, key, value string) string {
	stUrl, err := url.Parse(sUrl)
	if err != nil{
		return sUrl
	}

	m := stUrl.Query()
	m.Set(key, value)
	stUrl.RawQuery = m.Encode()
	return stUrl.String()
}

func RawURLAppendParam(rawurl, key, value string) (res string) {
	if strings.ContainsRune(rawurl, '?') {
		res = fmt.Sprint(rawurl, "&", key, "=", value)
	} else {
		res = fmt.Sprint(rawurl, "?", key, "=", value)
	}
	return
}

func RawURLAppendParams(rawurl, paras string) (res string) {
	if strings.ContainsRune(rawurl, '?') {
		res = fmt.Sprint(rawurl, "&", paras)
	} else {
		res = fmt.Sprint(rawurl, "?", paras)
	}
	return
}

// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name
// will get www.aspxfans.com
func RawUrlGetDomain(rawUrl string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return ""
	}
	return stUrl.Hostname()
}

// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name
// will get 8080
func RawUrlGetPort(rawUrl string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return ""
	}
	return stUrl.Port()
}