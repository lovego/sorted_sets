package sorted_sets

import (
	"reflect"
	"sort"
)

func RemoveValue(slice, target reflect.Value, fields ...string) reflect.Value {
	if !slice.IsValid() || slice.Len() == 0 {
		return slice
	}

	i := sort.Search(slice.Len(), func(i int) bool {
		return CompareValue(slice.Index(i), target, fields...) >= 0
	})

	if i < slice.Len() && CompareValue(slice.Index(i), target, fields...) == 0 {
		return reflect.AppendSlice(slice.Slice(0, i), slice.Slice(i+1, slice.Len()))
	}
	return slice
}
