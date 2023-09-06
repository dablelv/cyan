package net

import (
	"errors"
	"net/url"
)

// RawURLGetParam gets the first query parameter for the specified key from raw url,
// even though key has multiple parameters.
// E.g. rawURL=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name and key="page" will get "1".
func RawURLGetParam(raw, key string) (string, error) {
	r, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	m := r.Query()
	if v, ok := m[key]; ok && len(v) > 0 {
		return v[0], nil
	}
	return "", errors.New("no param")
}

// RawURLGetParams gets all query parameter for the specified key from raw url if key has multiple parameters.
// E.g. if rawURL=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name and key=page will get [1 2].
func RawURLGetParams(raw, key string) ([]string, error) {
	r, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}

	m := r.Query()
	if v, ok := m[key]; ok {
		return v, nil
	}
	return nil, errors.New("no param")
}

// RawURLGetAllParams gets all query parameter for all keys from raw url.
// E.g. if rawURL=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name will get map[boardID:[520] page:[1 2]].
func RawURLGetAllParams(raw string) (map[string][]string, error) {
	r, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}

	m := r.Query()
	return m, nil
}

// RawURLAddParam adds query parameter to raw url.
// E.g. if rawURL=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name and key=page value=3
// will get http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2&page=3#name.
func RawURLAddParam(raw, key, value string) (string, error) {
	r, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	m := r.Query()
	m.Add(key, value)
	r.RawQuery = m.Encode()
	return r.String(), nil
}

// RawURLAddParams adds multiple query parameters to raw url.
func RawURLAddParams(raw string, params map[string]string) (string, error) {
	r, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	m := r.Query()
	for k, v := range params {
		m.Add(k, v)
	}
	r.RawQuery = m.Encode()
	return r.String(), nil
}

// RawURLDelParams deletes multiple query paramters from raw url.
func RawURLDelParams(raw string, keys ...string) (string, error) {
	r, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	m := r.Query()
	for _, v := range keys {
		m.Del(v)
	}
	r.RawQuery = m.Encode()
	return r.String(), nil
}

// RawURLSetParam sets query paramter to raw url.
// If the query key does not exist, it will be added to the url.
// E.g. if rawURL=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name and key=page value=3
// will get http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=3#name.
func RawURLSetParam(raw, key, value string) (string, error) {
	r, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	m := r.Query()
	m.Set(key, value)
	r.RawQuery = m.Encode()
	return r.String(), nil
}

// RawURLSetParams sets multiple query paramter to raw url.
func RawURLSetParams(rawURL string, params map[string]string) (string, error) {
	r, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	m := r.Query()
	for k, v := range params {
		m.Set(k, v)
	}
	r.RawQuery = m.Encode()
	return r.String(), nil
}

// RawURLGetDomain gets the domain from raw url.
// E.g. if rawURL=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name
// will get www.aspxfans.com.
func RawURLGetDomain(rawURL string) (string, error) {
	r, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	return r.Hostname(), nil
}

// RawURLGetPort gets the port from raw url.
// E.g. if rawURL=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name
// will get 8080.
func RawURLGetPort(raw string) (string, error) {
	r, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	return r.Port(), nil
}

// QueryGetParam gets the specified key parameter from query string.
// The input is encoded query values without '?' and key is parameter name.
// E.g. if query string is "a=dog&b=tiger" and key is "a" will return dog.
func QueryGetParam(query, key string) (string, error) {
	queries, err := url.ParseQuery(query)
	if err != nil {
		return "", err
	}

	if v, ok := queries[key]; ok && len(v) > 0 {
		return v[0], nil
	}
	return "", errors.New("no param")
}

// QueryGetParams gets the specified key parameters from query string.
// Rawquery is encoded query values without '?' and key is parameter name.
// E.g. if query is "a=dog&a=cat&b=tiger" and key is "a" QueryGetParams will return [dog cat]
func QueryGetParams(query, key string) ([]string, error) {
	queries, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}

	if v, ok := queries[key]; ok {
		return v, nil
	}
	return nil, errors.New("no param")
}
