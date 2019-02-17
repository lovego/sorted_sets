package sorted_sets

import (
	"log"
	"reflect"
)

func Compare(a, b reflect.Value, fields ...string) int {
	var aIsNil, bIsNil bool
	for a.Kind() == reflect.Ptr || a.Kind() == reflect.Interface {
		if a.IsNil() {
			aIsNil = true
			break
		}
		a = a.Elem()
	}
	for b.Kind() == reflect.Ptr || b.Kind() == reflect.Interface {
		if b.IsNil() {
			bIsNil = true
			break
		}
		b = b.Elem()
	}
	switch {
	case aIsNil && !bIsNil:
		return -1
	case !aIsNil && bIsNil:
		return 1
	case aIsNil && bIsNil:
		return 0
	}

	switch a.Kind() {
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		return int(a.Int() - b.Int())
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		return int(int64(a.Uint()) - int64(b.Uint()))
	case reflect.String:
		as, bs := a.String(), b.String()
		if as == bs {
			return 0
		}
		if as < bs {
			return -1
		}
		return +1
	case reflect.Struct:
		if len(fields) == 0 {
			log.Panic("empty fields for struct value")
		}
		for _, name := range fields {
			if r := Compare(a.FieldByName(name), b.FieldByName(name)); r != 0 {
				return r
			}
		}
		return 0
	default:
		log.Panicf("unsupported value: %v(%v)", a.Type(), a)
	}
	return 0 // should not reach here
}
