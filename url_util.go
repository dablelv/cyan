package util

import (
	"errors"
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
func RawURLAddParam(rawUrl, key, value string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl
	}

	m := stUrl.Query()
	m.Add(key, value)
	stUrl.RawQuery = m.Encode()
	return stUrl.String()
}

func RawURLAddParams(rawUrl string, params map[string]string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl
	}

	m := stUrl.Query()
	for k, v := range params {
		m.Add(k, v)
	}
	stUrl.RawQuery = m.Encode()
	return stUrl.String()
}

// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name and key=page
// will get http://www.aspxfans.com:8080/news/index.asp?boardID=520#name
func RawURLDelParam(rawUrl, key string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl
	}

	m := stUrl.Query()
	m.Del(key)
	stUrl.RawQuery = m.Encode()
	return stUrl.String()
}

func RawURLDelParams(rawUrl string, keys []string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl
	}

	m := stUrl.Query()
	for _, v := range keys {
		m.Del(v)
	}
	stUrl.RawQuery = m.Encode()
	return stUrl.String()
}

// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name and key=page value=3
// will get http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=3#name
func RawURLSetParam(rawUrl, key, value string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl
	}

	m := stUrl.Query()
	m.Set(key, value)
	stUrl.RawQuery = m.Encode()
	return stUrl.String()
}

func RawURLSetParams(rawUrl string, params map[string]string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl
	}

	m := stUrl.Query()
	for k, v := range params {
		m.Set(k, v)
	}
	stUrl.RawQuery = m.Encode()
	return stUrl.String()
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

// RawQueryGetParam get the specified key parameter from query string
// rawquery is encoded query values without '?'. key is parameter name
// e.g. if query is "a=dog&b=tiger" and key is "a" will return dog
func RawQueryGetParam(rawquery, key string) (string, error) {
	queries, err := url.ParseQuery(rawquery)
	if err != nil {
		return "", err
	}

	if v, ok := queries[key]; ok && len(v) > 0 {
		return v[0], nil
	}
	return "", errors.New("no param")
}

// RawQueryGetParams get the specified key parameters from query string
// rawquery is encoded query values without '?'. key is parameter name
// e.g. if query is "a=dog&a=cat&b=tiger" and key is "a" will return [dog cat]
func RawQueryGetParams(rawquery, key string) ([]string, error) {
	queries, err := url.ParseQuery(rawquery)
	if err != nil {
		return nil, err
	}

	if v, ok := queries[key]; ok && len(v) > 0 {
		return v, nil
	}
	return nil, errors.New("no param")
}
