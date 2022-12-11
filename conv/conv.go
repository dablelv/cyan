package conv

import (
	"fmt"

	"github.com/spf13/cast"
)

// ToAnyE converts one type to another type.
func ToAny[T any](i any) T {
	v, _ := ToAnyE[T](i)
	return v
}

// ToAnyE converts one type to another and returns an error if occurred.
func ToAnyE[T any](i any) (T, error) {
	var t T
	switch any(t).(type) {
	case bool:
		v, err := cast.ToBoolE(i)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case int:
		v, err := cast.ToIntE(i)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case int8:
		v, err := cast.ToInt8E(i)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case int16:
		v, err := cast.ToInt16E(i)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case int32:
		v, err := cast.ToInt32E(i)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case int64:
		v, err := cast.ToInt64E(i)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case uint:
		v, err := cast.ToUintE(i)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case uint8:
		v, err := cast.ToUint8E(i)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case uint16:
		v, err := cast.ToUint16E(i)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case uint32:
		v, err := cast.ToUint32E(i)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case uint64:
		v, err := cast.ToUint64E(i)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case float32:
		v, err := cast.ToFloat32E(i)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case float64:
		v, err := cast.ToFloat64E(i)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case string:
		v, err := cast.ToStringE(i)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	default:
		return t, fmt.Errorf("the type %T is not supported", t)
	}
	return t, nil
}
