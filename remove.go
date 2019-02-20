package sorted_sets

import (
	"reflect"
	"sort"
)

func RemoveInt64(slice []int64, target int64, fields ...string) []int64 {
	if len(slice) == 0 {
		return slice
	}

	i := sort.Search(len(slice), func(i int) bool {
		return slice[i] >= target
	})

	if i < len(slice) && slice[i] == target {
		return append(slice[:i], slice[i+1:]...)
	}
	return slice
}

func RemoveString(slice []string, target string, fields ...string) []string {
	if len(slice) == 0 {
		return slice
	}

	i := sort.Search(len(slice), func(i int) bool {
		return slice[i] >= target
	})

	if i < len(slice) && slice[i] == target {
		return append(slice[:i], slice[i+1:]...)
	}
	return slice
}

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
