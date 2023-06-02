package conv

import (
	"encoding/json"
	"errors"
	"html/template"
	"reflect"
	"testing"
	"time"
)

func TestToBoolE(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"string true to bool", args{"true"}, true, false},
		{"string false to bool", args{"false"}, false, false},
		{"string True to bool", args{"True"}, true, false},
		{"string False to bool", args{"False"}, false, false},
		{"nil to bool", args{i: nil}, false, false},
		{"int 0 to bool", args{0}, false, false},
		{"int 1 to bool", args{1}, true, false},
		{"int64 0 to bool", args{int64(0)}, false, false},
		{"int64 1 to bool", args{int64(1)}, true, false},
		{"int32 0 to bool", args{int32(0)}, false, false},
		{"int32 1 to bool", args{int32(1)}, true, false},
		{"int16 0 to bool", args{int16(0)}, false, false},
		{"int16 1 to bool", args{int16(1)}, true, false},
		{"int8 0 to bool", args{int8(0)}, false, false},
		{"int8 1 to bool", args{int8(1)}, true, false},
		{"uint 0 to bool", args{uint(0)}, false, false},
		{"uint 1 to bool", args{uint(1)}, true, false},
		{"uint64 0 to bool", args{uint64(0)}, false, false},
		{"uint64 1 to bool", args{uint64(1)}, true, false},
		{"uint32 0 to bool", args{uint32(0)}, false, false},
		{"uint32 1 to bool", args{uint32(1)}, true, false},
		{"uint16 0 to bool", args{uint16(0)}, false, false},
		{"uint16 1 to bool", args{uint16(1)}, true, false},
		{"uint8 0 to bool", args{uint8(0)}, false, false},
		{"uint8 1 to bool", args{uint8(1)}, true, false},
		{"time.Duration 0 to bool", args{time.Duration(0)}, false, false},
		{"time.Duration 1 to bool", args{time.Duration(1)}, true, false},
		{"json.Number 0 to bool", args{json.Number("0")}, false, false},
		{"json.Number 1 to bool", args{json.Number("1")}, true, false},
		{"json.Number foo to bool failed", args{json.Number("foo")}, false, true},
		{"array int to bool failed", args{[]int{}}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToBoolE(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToBoolE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToBoolE() = %v, want %v", got, tt.want)
			}
		})
	}
}

type TestCase struct {
	name    string
	input   any
	want    any
	wantErr bool
}

var (
	one    int  = 1
	ptrOne *int = &one
)

func genIntegerTests() []TestCase {
	tests := []TestCase{
		{"bool true", true, 1, false},
		{"bool false", false, 0, false},
		{"int 1", 1, 1, false},
		{"int8 1 ", int8(1), 1, false},
		{"int16 1", int16(1), 1, false},
		{"int32 1", int32(1), 1, false},
		{"int64 1", int64(1), 1, false},
		{"uint 1", uint(1), 1, false},
		{"uint8 1", uint8(1), 1, false},
		{"uint16 1", uint16(1), 1, false},
		{"uint32 1", uint32(1), 1, false},
		{"uint64 1", uint64(1), 1, false},
		{"float32 1.1", float32(1.1), 1, false},
		{"float64 1.1", 1.1, 1, false},
		{"string 1.00", "1.00", 1, false},
		{"string 1.01 failed", "1.01", 0, true},
		{"json.Number 1", json.Number("1"), 1, false},
		{"nil", nil, 0, false},
		{"time.Month 1", time.Month(1), 1, false},
		{"time.Weekday 1", time.Weekday(1), 1, false},
		{"pointer to int 1", ptrOne, 1, false},
		{"complex64 3-5i failed", complex(3, -5), 0, true},
	}
	return tests
}

func TestToIntE(t *testing.T) {
	for _, tt := range genIntegerTests() {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToIntE(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToIntE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToIntE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt8E(t *testing.T) {
	for _, tt := range genIntegerTests() {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt8E(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt8E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != int8(tt.want.(int)) {
				t.Errorf("ToInt8E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt16E(t *testing.T) {
	for _, tt := range genIntegerTests() {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt16E(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt16E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != int16(tt.want.(int)) {
				t.Errorf("ToInt16E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt32E(t *testing.T) {
	for _, tt := range genIntegerTests() {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt32E(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt32E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != int32(tt.want.(int)) {
				t.Errorf("ToInt32E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt64E(t *testing.T) {
	for _, tt := range genIntegerTests() {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt64E(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt64E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != int64(tt.want.(int)) {
				t.Errorf("ToInt64E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func genUnsignedInteger() []TestCase {
	tests := genIntegerTests()
	unsignedTests := []TestCase{
		{"int -1", -1, 0, true},
		{"int8 -1 ", int8(-1), 0, true},
		{"int16 -1", int16(-1), 0, true},
		{"int32 -1", int32(-1), 0, true},
		{"int64 -1", int64(-1), 0, true},
		{"float32 -1.1", float32(-1.1), 0, true},
		{"float64 -1.1", -1.1, 0, true},
		{"string -1.00", "-1.00", 0, true},
		{"json.Number -1", json.Number("-1"), 0, true},
	}
	return append(tests, unsignedTests...)
}

func TestToUintE(t *testing.T) {
	for _, tt := range genUnsignedInteger() {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUintE(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUintE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != uint(tt.want.(int)) {
				t.Errorf("ToUintE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint8E(t *testing.T) {
	for _, tt := range genUnsignedInteger() {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint8E(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint8E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != uint8(tt.want.(int)) {
				t.Errorf("ToUint8E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint16E(t *testing.T) {
	for _, tt := range genUnsignedInteger() {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint16E(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint16E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != uint16(tt.want.(int)) {
				t.Errorf("ToUint16E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint32E(t *testing.T) {
	for _, tt := range genUnsignedInteger() {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint32E(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint32E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != uint32(tt.want.(int)) {
				t.Errorf("ToUint32E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint64E(t *testing.T) {
	for _, tt := range genUnsignedInteger() {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint64E(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint64E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != uint64(tt.want.(int)) {
				t.Errorf("ToUint64E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func genFloatTests() []TestCase {
	tests := []TestCase{
		{"true", true, 1.0, false},
		{"false", false, 0.0, false},
		{"int 1", 1, 1.0, false},
		{"int64 1", int64(1), 1.0, false},
		{"int32 1", int32(1), 1.0, false},
		{"int16 1", int16(1), 1.0, false},
		{"int8 1", int8(1), 1.0, false},
		{"uint 1", uint(1), 1.0, false},
		{"uint64 1", uint64(1), 1.0, false},
		{"uint32 1", uint32(1), 1.0, false},
		{"uint16 1", uint16(1), 1.0, false},
		{"uint8 1", uint8(1), 1.0, false},
		{"float64 1.1", 1.1, 1.1, false},
		{"float32 1.1", float32(1.5), 1.5, false},
		{"float64 1.5", float64(1.1), 1.1, false},
		{"json.Number 1.0", json.Number("1.0"), 1.0, false},
		{"nil", nil, 0.0, false},
		{"string 1.00", "1.00", 1.0, false},
		{"string foo failed", "foo", 0.0, true},
		{"complex64 3-5i failed", complex(3, -5), 0.0, true},
	}
	return tests
}

func TestToFloat32E(t *testing.T) {
	for _, tt := range genFloatTests() {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToFloat32E(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFloat32E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != float32(tt.want.(float64)) {
				t.Errorf("ToFloat32E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat64E(t *testing.T) {
	for _, tt := range genFloatTests() {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToFloat64E(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFloat64E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want.(float64) {
				t.Errorf("ToFloat64E() = %v, want %v", got, tt.want)
			}
		})
	}
}

type foo struct {
	val string
}

func (f foo) String() string {
	return f.val
}

func TestToStringE(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"string foo", args{"foo"}, "foo", false},
		{"true", args{true}, "true", false},
		{"false", args{false}, "false", false},
		{"int 8", args{8}, "8", false},
		{"int64 8", args{int64(8)}, "8", false},
		{"int32 8", args{int32(8)}, "8", false},
		{"int16 8", args{int16(8)}, "8", false},
		{"int8 8", args{int8(8)}, "8", false},
		{"uint 8", args{uint(8)}, "8", false},
		{"uint64 8", args{uint64(8)}, "8", false},
		{"uint32 8", args{uint32(8)}, "8", false},
		{"uint16 8", args{uint16(8)}, "8", false},
		{"uint8 8", args{uint8(8)}, "8", false},
		{"float64 8.31", args{8.31}, "8.31", false},
		{"float32 8.31", args{float32(8.31)}, "8.31", false},
		{"json.Number 1.0", args{json.Number("1.0")}, "1.0", false},
		{"byte slice foo", args{[]byte("foo")}, "foo", false},
		{"nil", args{nil}, "", false},
		{"HTML", args{template.HTML("one time")}, "one time", false},
		{"HTMLAttr", args{template.HTMLAttr("a")}, "a", false},
		{"URL", args{template.URL("https://host.com")}, "https://host.com", false},
		{"JS", args{template.JS("(1+2)")}, "(1+2)", false},
		{"JSStr", args{template.JSStr("foo\\nbar")}, "foo\\nbar", false},
		{"CSS", args{template.CSS("a")}, "a", false},
		{"template.Srcset", args{template.Srcset("secret")}, "secret", false},
		{"fmt.Stringer", args{foo{"foo"}}, "foo", false},
		{"error to string", args{errors.New("err")}, "err", false},
		{"int slice failed", args{[]int{}}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToStringE(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToStringE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToStringE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToDurationE(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    time.Duration
		wantErr bool
	}{
		{"duration", args{time.Second}, time.Second, false},
		{"int 8", args{8}, time.Duration(8), false},
		{"int64 8", args{int64(8)}, time.Duration(8), false},
		{"int32 8", args{int32(8)}, time.Duration(8), false},
		{"int16 8", args{int16(8)}, time.Duration(8), false},
		{"int8 8", args{int8(8)}, time.Duration(8), false},
		{"uint 8", args{uint(8)}, time.Duration(8), false},
		{"uint64 8", args{uint64(8)}, time.Duration(8), false},
		{"uint32 8", args{uint32(8)}, time.Duration(8), false},
		{"uint16 8", args{uint16(8)}, time.Duration(8), false},
		{"uint8 8", args{uint8(8)}, time.Duration(8), false},
		{"float64 8.31", args{8.31}, time.Duration(8), false},
		{"float32 8.31", args{float32(8.31)}, time.Duration(8), false},
		{"string 8ns", args{"8ns"}, time.Duration(8), false},
		{"string foo failed", args{"foo"}, time.Duration(0), true},
		{"json.Number", args{json.Number("8")}, time.Duration(8), false},
		{"bool failed", args{true}, time.Duration(0), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotD, err := ToDurationE(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToDurationE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotD, tt.want) {
				t.Errorf("ToDurationE() = %v, want %v", gotD, tt.want)
			}
		})
	}
}
