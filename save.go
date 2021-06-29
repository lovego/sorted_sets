package sorted_sets

import (
	"reflect"
	"sort"

	slicePkg "github.com/lovego/slice"
)

func SaveInt64(slice []int64, target int64) []int64 {
	if len(slice) == 0 {
		return []int64{target}
	}

	i := sort.Search(len(slice), func(i int) bool {
		return slice[i] >= target
	})

	if i >= len(slice) { // not found
		return append(slice, target)
	}
	if slice[i] == target {
		return slice
	}
	return slicePkg.InsertInt64(slice, i, target)
}

func SaveString(slice []string, target string) []string {
	if len(slice) == 0 {
		return []string{target}
	}

	i := sort.Search(len(slice), func(i int) bool {
		return slice[i] >= target
	})

	if i >= len(slice) { // not found
		return append(slice, target)
	}
	if slice[i] == target {
		return slice
	}
	return slicePkg.InsertString(slice, i, target)
}

func SaveValue(slice, target reflect.Value, fields ...string) reflect.Value {
	if !slice.IsValid() {
		return reflect.Append(reflect.MakeSlice(reflect.SliceOf(target.Type()), 0, 1), target)
	}
	if slice.Len() == 0 {
		return reflect.Append(slice, target)
	}

	i := sort.Search(slice.Len(), func(i int) bool {
		return CompareValue(slice.Index(i), target, fields...) >= 0
	})

	if i >= slice.Len() { // not found
		return reflect.Append(slice, target)
	}
	if CompareValue(slice.Index(i), target, fields...) == 0 {
		slice.Index(i).Set(target)
		return slice
	}
	return slicePkg.InsertValue(slice, i, target)
}
