package sorted_sets

import (
	"reflect"
	"sort"
)

func RemoveInt64(slicePointer *[]int64, target int64) {
	if slicePointer == nil {
		return
	}
	slice := *slicePointer
	if len(slice) == 0 {
		return
	}

	i := sort.Search(len(slice), func(i int) bool {
		return slice[i] >= target
	})

	if i < len(slice) && slice[i] == target {
		*slicePointer = append(slice[:i], slice[i+1:]...)
	}
}

func RemoveString(slicePointer *[]string, target string) {
	if slicePointer == nil {
		return
	}
	slice := *slicePointer
	if len(slice) == 0 {
		return
	}

	i := sort.Search(len(slice), func(i int) bool {
		return slice[i] >= target
	})

	if i < len(slice) && slice[i] == target {
		*slicePointer = append(slice[:i], slice[i+1:]...)
	}
}

// slice must be settlable or invalid
func RemoveValue(slice, target reflect.Value, fields ...string) {
	if !slice.IsValid() || slice.Len() == 0 {
		return
	}

	i := sort.Search(slice.Len(), func(i int) bool {
		return CompareValue(slice.Index(i), target, fields...) >= 0
	})

	if i < slice.Len() && CompareValue(slice.Index(i), target, fields...) == 0 {
		slice.Set(reflect.AppendSlice(slice.Slice(0, i), slice.Slice(i+1, slice.Len())))
	}
	return
}
