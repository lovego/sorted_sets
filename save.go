package sorted_sets

import (
	"reflect"
	"sort"
)

func Save(slice, target reflect.Value, fields ...string) reflect.Value {
	if !slice.IsValid() {
		return reflect.Append(reflect.MakeSlice(reflect.SliceOf(target.Type()), 0, 1), target)
	}
	if slice.Len() == 0 {
		return reflect.Append(slice, target)
	}

	i := sort.Search(slice.Len(), func(i int) bool {
		return Compare(slice.Index(i), target, fields...) >= 0
	})

	if i >= slice.Len() { // not found
		return reflect.Append(slice, target)
	}
	if Compare(slice.Index(i), target, fields...) == 0 {
		slice.Index(i).Set(target)
		return slice
	}
	return Insert(slice, i, target)
}

func Insert(s reflect.Value, i int, v reflect.Value) reflect.Value {
	s = reflect.Append(s, v)
	reflect.Copy(s.Slice(i+1, s.Len()), s.Slice(i, s.Len()-1))
	s.Index(i).Set(v)
	return s
}
