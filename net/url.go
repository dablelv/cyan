package util

import (
	"errors"
	"net/url"
)

// RawURLGetParam gets the first query parameter for the specified key from raw url,
// even though key has multiple parameters.
// E.g. rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name and key="page" will get "1".
func RawURLGetParam(raw, key string) (string, error) {
	stUrl, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	m := stUrl.Query()
	if v, ok := m[key]; ok && len(v) > 0 {
		return v[0], nil
	}
	return "", errors.New("no param")
}

// RawURLGetParams gets all query parameter for the specified key from raw url if key has multiple parameters.
// E.g. if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name and key=page will get [1 2].
func RawURLGetParams(raw, key string) ([]string, error) {
	stUrl, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}

	m := stUrl.Query()
	if v, ok := m[key]; ok {
		return v, nil
	}
	return nil, errors.New("no param")
}

// RawURLGetAllParams gets all query parameter for all keys from raw url.
// E.g. if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name will get map[boardID:[520] page:[1 2]].
func RawURLGetAllParams(raw string) (map[string][]string, error) {
	stUrl, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}

	m := stUrl.Query()
	return m, nil
}

// RawURLAddParam adds query parameter to raw url.
// E.g. if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name and key=page value=3
// will get http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2&page=3#name.
func RawURLAddParam(raw, key, value string) (string, error) {
	stUrl, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	m := stUrl.Query()
	m.Add(key, value)
	stUrl.RawQuery = m.Encode()
	return stUrl.String(), nil
}

// RawURLAddParams adds multiple query parameters to raw url.
func RawURLAddParams(raw string, params map[string]string) (string, error) {
	stUrl, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	m := stUrl.Query()
	for k, v := range params {
		m.Add(k, v)
	}
	stUrl.RawQuery = m.Encode()
	return stUrl.String(), nil
}

// RawURLDelParam deletes query paramter from the raw url.
// E.g. if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name and key=page
// will get http://www.aspxfans.com:8080/news/index.asp?boardID=520#name.
func RawURLDelParam(raw, key string) (string, error) {
	stUrl, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	m := stUrl.Query()
	m.Del(key)
	stUrl.RawQuery = m.Encode()
	return stUrl.String(), nil
}

// RawURLDelParams deletes multiple query paramters from raw url.
func RawURLDelParams(raw string, keys []string) (string, error) {
	stUrl, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	m := stUrl.Query()
	for _, v := range keys {
		m.Del(v)
	}
	stUrl.RawQuery = m.Encode()
	return stUrl.String(), nil
}

// RawURLSetParam sets query paramter to raw url.
// E.g. if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name and key=page value=3
// will get http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=3#name.
func RawURLSetParam(raw, key, value string) (string, error) {
	stUrl, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	m := stUrl.Query()
	m.Set(key, value)
	stUrl.RawQuery = m.Encode()
	return stUrl.String(), nil
}

// RawURLSetParams sets multiple query paramter to raw url.
func RawURLSetParams(raw string, params map[string]string) (string, error) {
	stUrl, err := url.Parse(raw)
	if err != nil {
		return "", nil
	}

	m := stUrl.Query()
	for k, v := range params {
		m.Set(k, v)
	}
	stUrl.RawQuery = m.Encode()
	return stUrl.String(), nil
}

// RawUrlGetDomain gets the domain from raw url.
// E.g. if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name
// will get www.aspxfans.com.
func RawUrlGetDomain(raw string) (string, error) {
	stUrl, err := url.Parse(raw)
	if err != nil {
		return "", nil
	}
	return stUrl.Hostname(), nil
}

// RawUrlGetPort gets the port from raw url.
// E.g. if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name
// will get 8080.
func RawUrlGetPort(raw string) (string, error) {
	stUrl, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	return stUrl.Port(), nil
}

// RawQueryGetParam gets the specified key parameter from query string.
// Rawquery is encoded query values without '?' and key is parameter name.
// E.g. if query string is "a=dog&b=tiger" and key is "a" will return dog.
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

// RawQueryGetParams gets the specified key parameters from query string.
// Rawquery is encoded query values without '?' and key is parameter name
// E.g. if query is "a=dog&a=cat&b=tiger" and key is "a" will return [dog cat]
func RawQueryGetParams(query, key string) ([]string, error) {
	queries, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}

	if v, ok := queries[key]; ok {
		return v, nil
	}
	return nil, errors.New("no param")
}
