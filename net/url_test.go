package net

import (
	"reflect"
	"testing"
)

func TestRawURLGetParam(t *testing.T) {
	type args struct {
		raw string
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "single param",
			args: args{
				raw: "http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name",
				key: "boardID",
			},
			want:    "520",
			wantErr: false,
		},
		{
			name: "multiple params",
			args: args{
				raw: "http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name",
				key: "page",
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "no param",
			args: args{
				raw: "http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name",
				key: "foo",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "url invalid",
			args: args{
				raw: "\r",
				key: "foo",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RawURLGetParam(tt.args.raw, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("RawURLGetParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RawURLGetParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRawURLGetParams(t *testing.T) {
	type args struct {
		raw string
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "single param",
			args: args{
				raw: "http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name",
				key: "boardID",
			},
			want:    []string{"520"},
			wantErr: false,
		},
		{
			name: "multiple params",
			args: args{
				raw: "http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name",
				key: "page",
			},
			want:    []string{"1", "2"},
			wantErr: false,
		},
		{
			name: "no param",
			args: args{
				raw: "http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name",
				key: "foo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "url invalid",
			args: args{
				raw: "http://www.aspxfans.com:8080/news/index.asp?boardID\r=520&page=1&page=2#name",
				key: "foo",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RawURLGetParams(tt.args.raw, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("RawURLGetParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RawURLGetParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRawURLGetAllParams(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string][]string
		wantErr bool
	}{
		{
			name: "single param",
			args: args{
				raw: "http://www.aspxfans.com:8080/news/index.asp?boardID=520#name",
			},
			want:    map[string][]string{"boardID": {"520"}},
			wantErr: false,
		},
		{
			name: "multiple params",
			args: args{
				raw: "http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name",
			},
			want:    map[string][]string{"boardID": {"520"}, "page": {"1", "2"}},
			wantErr: false,
		},
		{
			name: "no param",
			args: args{
				raw: "http://www.aspxfans.com:8080/news/index.asp#name",
			},
			want:    map[string][]string{},
			wantErr: false,
		},
		{
			name: "url invalid",
			args: args{
				raw: "http://www.aspxfans.com:8080/news/index.asp\r#name",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RawURLGetAllParams(tt.args.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("RawURLGetAllParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RawURLGetAllParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRawURLAddParam(t *testing.T) {
	type args struct {
		raw   string
		key   string
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "single param",
			args: args{
				raw:   "http://www.aspxfans.com:8080/news/index.asp?boardID=520#name",
				key:   "page",
				value: "1",
			},
			want:    "http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1#name",
			wantErr: false,
		},
		{
			name: "multiple params",
			args: args{
				raw:   "http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1#name",
				key:   "page",
				value: "2",
			},
			want:    "http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name",
			wantErr: false,
		},
		{
			name: "only host url",
			args: args{
				raw:   "foo",
				key:   "page",
				value: "1",
			},
			want:    "foo?page=1",
			wantErr: false,
		},
		{
			name: "url invalid",
			args: args{
				raw:   "foo\r",
				key:   "page",
				value: "1",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RawURLAddParam(tt.args.raw, tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("RawURLAddParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RawURLAddParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRawURLAddParams(t *testing.T) {
	type args struct {
		raw    string
		params map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "multiple params",
			args: args{
				raw:    "http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1#name",
				params: map[string]string{"page": "2"},
			},
			want:    "http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name",
			wantErr: false,
		},
		{
			name: "only host url",
			args: args{
				raw:    "foo",
				params: map[string]string{"board_id": "520", "page": "1"},
			},
			want:    "foo?board_id=520&page=1",
			wantErr: false,
		},
		{
			name: "url invalid",
			args: args{
				raw:    "foo\r",
				params: map[string]string{"board_id": "520", "page": "1"},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RawURLAddParams(tt.args.raw, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("RawURLAddParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RawURLAddParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRawURLDelParams(t *testing.T) {
	type args struct {
		raw  string
		keys []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "multiple params",
			args: args{
				raw:  "http://www.aspxfans.com:8080/news/index.asp?boardID=520&&page=1&page=2#name",
				keys: []string{"boardID", "page"},
			},
			want:    "http://www.aspxfans.com:8080/news/index.asp#name",
			wantErr: false,
		},
		{
			name: "only host url",
			args: args{
				raw:  "foo?",
				keys: []string{"page"},
			},
			want:    "foo?",
			wantErr: false,
		},
		{
			name: "no keys",
			args: args{
				raw:  "foo?",
				keys: nil,
			},
			want:    "foo?",
			wantErr: false,
		},
		{
			name: "url invalid",
			args: args{
				raw:  "foo?\r",
				keys: []string{"page"},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RawURLDelParams(tt.args.raw, tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("RawURLDelParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RawURLDelParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRawURLSetParam(t *testing.T) {
	type args struct {
		raw   string
		key   string
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "multiple params",
			args: args{
				raw:   "http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name",
				key:   "page",
				value: "3",
			},
			want:    "http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=3#name",
			wantErr: false,
		},
		{
			name: "only host url",
			args: args{
				raw:   "foo",
				key:   "page",
				value: "1",
			},
			want:    "foo?page=1",
			wantErr: false,
		},
		{
			name: "url invalid",
			args: args{
				raw:   "foo\r",
				key:   "page",
				value: "1",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RawURLSetParam(tt.args.raw, tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("RawURLSetParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RawURLSetParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRawURLSetParams(t *testing.T) {
	type args struct {
		raw    string
		params map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "multiple params",
			args: args{
				raw:    "http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name",
				params: map[string]string{"boardID": "1024", "page": "3"},
			},
			want:    "http://www.aspxfans.com:8080/news/index.asp?boardID=1024&page=3#name",
			wantErr: false,
		},
		{
			name: "only host url",
			args: args{
				raw:    "foo?",
				params: map[string]string{"boardID": "1024", "page": "3"},
			},
			want:    "foo?boardID=1024&page=3",
			wantErr: false,
		},
		{
			name: "url invalid",
			args: args{
				raw:    "foo\r?",
				params: map[string]string{"board_id": "1024", "page": "3"},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RawURLSetParams(tt.args.raw, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("RawURLSetParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RawURLSetParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRawURLGetDomain(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "multiple params",
			args: args{
				raw: "http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name",
			},
			want:    "www.aspxfans.com",
			wantErr: false,
		},
		{
			name: "no host url",
			args: args{
				raw: "http://:8080/news/index.asp?boardID=520&page=1&page=2#name",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "url invalid",
			args: args{
				raw: "http://foo\r:8080/news/index.asp?boardID=520&page=1&page=2#name",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RawURLGetDomain(tt.args.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("RawURLGetDomain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RawURLGetDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRawURLGetPort(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "multiple params",
			args: args{
				raw: "http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name",
			},
			want:    "8080",
			wantErr: false,
		},
		{
			name: "no port url",
			args: args{
				raw: "http://www.aspxfans.com/news/index.asp?boardID=520&page=1&page=2#name",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "url invalid",
			args: args{
				raw: "http://foo\r:8080/news/index.asp?boardID=520&page=1&page=2#name",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RawURLGetPort(tt.args.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("RawURLGetPort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RawURLGetPort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRawQueryGetParam(t *testing.T) {
	type args struct {
		rawquery string
		key      string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "single value",
			args:    args{"a=dog&b=tiger", "a"},
			want:    "dog",
			wantErr: false,
		},
		{
			name:    "multiple values",
			args:    args{"a=dog&a=cat&b=tiger", "a"},
			want:    "dog",
			wantErr: false,
		},
		{
			name:    "no value",
			args:    args{"a=dog&a=cat&b=tiger&c=", "c"},
			want:    "",
			wantErr: false,
		},
		{
			name:    "no key",
			args:    args{"a=dog&a=cat&b=tiger&c=", "d"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "query invalid",
			args:    args{"a=dog;&a=cat&b=tiger&c=", "a"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryGetParam(tt.args.rawquery, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryGetParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("QueryGetParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryGetParams(t *testing.T) {
	type args struct {
		query string
		key   string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "single value",
			args:    args{"a=dog&b=tiger", "a"},
			want:    []string{"dog"},
			wantErr: false,
		},
		{
			name:    "multiple values",
			args:    args{"a=dog&a=cat&b=tiger", "a"},
			want:    []string{"dog", "cat"},
			wantErr: false,
		},
		{
			name:    "no value",
			args:    args{"a=dog&a=cat&b=tiger&c=", "c"},
			want:    []string{""},
			wantErr: false,
		},
		{
			name:    "query invalid",
			args:    args{"a=dog;", "d"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "no key",
			args:    args{"a=dog&a=cat&b=tiger&c=", "d"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryGetParams(tt.args.query, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryGetParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryGetParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
