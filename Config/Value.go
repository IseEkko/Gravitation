package Config

import (
	"errors"
	"fmt"
	"strconv"
	"sync/atomic"
)

var ConversionTypeError error = errors.New("Conversion type error")

var (
	_ Value = (*atomicValue)(nil)
)

type Value interface {
	Int() (int64, error)
	string() (string, error)
}

type atomicValue struct {
	atomic.Value
}

func (a atomicValue) Int() (int64, error) {
	switch val := a.Load().(type) {
	case int:
		return int64(val), nil
	case int8:
		return int64(val), nil
	case int16:
		return int64(val), nil
	case int32:
		return int64(val), nil
	case int64:
		return val, nil
	case uint:
		return int64(val), nil
	case uint8:
		return int64(val), nil
	case uint16:
		return int64(val), nil
	case uint32:
		return int64(val), nil
	case uint64:
		return int64(val), nil
	case float32:
		return int64(val), nil
	case float64:
		return int64(val), nil
	case string:
		return strconv.ParseInt(val, 10, 64)
	}
	return 0, ConversionTypeError //这里后续更改为错误类型
}

func (a atomicValue) string() (string, error) {
	switch val := a.Load().(type) {
	case string:
		return val, nil
	case bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return fmt.Sprint(val), nil
	case []byte:
		return string(val), nil
	case fmt.Stringer:
		return val.String(), nil
	}
	return "", ConversionTypeError
}
