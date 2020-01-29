package format

import (
	"reflect"
	"strconv"
)

func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// 格式化一个值，且不分析它的内部结构
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8,reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8,reflect.Uint16,reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	default:
		return v.Type().String() + " value"
	}
}
