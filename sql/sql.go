package util

import (
	"bytes"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cast"
)

type DataRow map[string]interface{}

func (row DataRow) IsExists(str string) bool {
	_, ok := row[str]
	return ok
}

func (row DataRow) GetString(str string) string {
	v, ok := row[str]
	if ok {
		return cast.ToString(v)
	}

	return ""
}

func (row DataRow) GetStringWithDef(str string, def string) string {
	v, ok := row[str]
	if ok {
		return cast.ToString(v)
	}

	return def
}

func (row DataRow) GetInt(str string) int {
	v, ok := row[str]
	if ok {
		return cast.ToInt(v)
	}

	return 0
}

func (row DataRow) GetIntWithDef(str string, def int) int {
	v, ok := row[str]
	if ok {
		return cast.ToInt(v)
	}

	return def
}

func (row DataRow) GetUint(str string) uint {
	v, ok := row[str]
	if ok {
		return cast.ToUint(v)
	}

	return 0
}

func (row DataRow) GetUintWithDef(str string, def uint) uint {
	v, ok := row[str]
	if ok {
		return cast.ToUint(v)
	}

	return def
}

func (row DataRow) GetInt32(str string) int32 {
	v, ok := row[str]
	if ok {
		return cast.ToInt32(v)
	}

	return 0
}

func (row DataRow) GetInt32WithDef(str string, def int32) int32 {
	v, ok := row[str]
	if ok {
		return cast.ToInt32(v)
	}

	return def
}

func (row DataRow) GetUint32(str string) uint32 {
	v, ok := row[str]
	if ok {
		return cast.ToUint32(v)
	}

	return 0
}

func (row DataRow) GetUint32WithDef(str string, def uint32) uint32 {
	v, ok := row[str]
	if ok {
		return cast.ToUint32(v)
	}

	return def
}

func (row DataRow) GetInt64(str string) int64 {
	v, ok := row[str]
	if ok {
		return cast.ToInt64(v)
	}

	return 0
}

func (row DataRow) GetInt64WithDef(str string, def int64) int64 {
	v, ok := row[str]
	if ok {
		return cast.ToInt64(v)
	}

	return def
}

func (row DataRow) GetUint64(str string) uint64 {
	v, ok := row[str]
	if ok {
		return cast.ToUint64(v)
	}

	return 0
}

func (row DataRow) GetUint64WithDef(str string, def uint64) uint64 {
	v, ok := row[str]
	if ok {
		return cast.ToUint64(v)
	}

	return def
}

func (row DataRow) GetFloat32(str string) float32 {
	v, ok := row[str]
	if ok {
		return cast.ToFloat32(v)
	}

	return 0
}

func (row DataRow) GetFloat32WithDef(str string, def float32) float32 {
	v, ok := row[str]
	if ok {
		return cast.ToFloat32(v)
	}

	return def
}

func (row DataRow) GetFloat64(str string) float64 {
	v, ok := row[str]
	if ok {
		return cast.ToFloat64(v)
	}

	return 0
}

func (row DataRow) GetFloat64WithDef(str string, def float64) float64 {
	v, ok := row[str]
	if ok {
		return cast.ToFloat64(v)
	}

	return def
}

// EscapeString escapes string with backslashes (\)
// This escapes the contents of a string  by adding backslashes before special
// characters, and turning others into specific escape sequences, such as
// turning newlines into \n and null bytes into \0.
// https://github.com/mysql/mysql-server/blob/mysql-5.7.5/mysys/charset.c#L823-L932
func EscapeString(src string) string {
	if len(src) == 0 {
		return ""
	}

	var pos int
	dst := make([]byte, len(src)*2)
	for i := 0; i < len(src); i++ {
		switch src[i] {
		case '\x00':
			dst[pos] = '\\'
			dst[pos+1] = '0'
			pos += 2
		case '\n':
			dst[pos] = '\\'
			dst[pos+1] = 'n'
			pos += 2
		case '\r':
			dst[pos] = '\\'
			dst[pos+1] = 'r'
			pos += 2
		case '\x1a':
			dst[pos] = '\\'
			dst[pos+1] = 'Z'
			pos += 2
		case '\'':
			dst[pos] = '\\'
			dst[pos+1] = '\''
			pos += 2
		case '"':
			dst[pos] = '\\'
			dst[pos+1] = '"'
			pos += 2
		case '\\':
			dst[pos] = '\\'
			dst[pos+1] = '\\'
			pos += 2
		default:
			dst[pos] = src[i]
			pos++
		}
	}

	return string(dst[0:pos])
}

// FormatSql replace the placeholder in sql with the specified string
// src: original sql string with placeholder
// dic: placeholder map with type map[string]string or map[string]interface{}
// escape: whether to escape the string corresponding to the placeholder
// e.g. if src="select * from t where field1={name}" and dic=map[name:dog], the result
// sql statement will be "select * from t where field1=dog"
func FormatSql(src string, dic interface{}, escape bool) string {
	var idx = 0
	var lastIdx = 0
	var end = len(src)
	var state = 0
	var buf bytes.Buffer
	m := cast.ToStringMapString(dic)

	for {
		if state == 0 {
			if idx == end {
				buf.WriteString(src[lastIdx:idx])
				break
			} else if src[idx] == '{' {
				state = 1
				buf.WriteString(src[lastIdx:idx])
				lastIdx = idx
				idx++
			} else {
				idx++
			}
		} else if state == 1 {
			if idx == end {
				buf.WriteString(src[lastIdx:idx])
				break
			} else if src[idx] == '{' {
				buf.WriteString(src[lastIdx:idx])
				lastIdx = idx
				idx++
			} else if src[idx] == '}' {
				state = 0
				if v, ok := m[src[lastIdx+1:idx]]; ok {
					if escape {
						buf.WriteString(EscapeString(v))
					} else {
						buf.WriteString(v)
					}
				} else {
					buf.WriteString(src[lastIdx : idx+1])
				}
				idx++
				lastIdx = idx
			} else {
				idx++
			}
		}
	}
	return buf.String()
}

func FetchRowBySql(conn *sql.DB, sSql string, dic map[string]interface{}) (DataRow, error) {
	if dic != nil {
		sSql = FormatSql(sSql, dic, true)
	}

	rows, err := conn.Query(sSql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(DataRow)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		for i, v := range values {
			if v != nil {
				record[columns[i]] = string(v.([]byte))
			} else {
				record[columns[i]] = v
			}
		}
		break
	}

	return record, nil
}

func FetchRowsBySql(conn *sql.DB, sSql string, dic map[string]interface{}) ([]DataRow, error) {
	if dic != nil {
		sSql = FormatSql(sSql, dic, true)
	}

	rows, err := conn.Query(sSql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	records := make([]DataRow, 0)
	for rows.Next() {
		record := make(DataRow)
		err = rows.Scan(scanArgs...)
		for i, v := range values {
			if v != nil {
				record[columns[i]] = string(v.([]byte))
			} else {
				record[columns[i]] = v
			}
		}

		records = append(records, record)
	}
	return records, nil
}
