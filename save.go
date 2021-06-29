package sorted_sets

import (
	"reflect"
	"sort"

	slicePkg "github.com/lovego/slice"
)

func SaveInt64(slicePointer *[]int64, target int64) {
	slice := *slicePointer
	if len(slice) == 0 {
		*slicePointer = []int64{target}
		return
	}

	i := sort.Search(len(slice), func(i int) bool {
		return slice[i] >= target
	})

	if i >= len(slice) { // not found
		*slicePointer = append(slice, target)
		return
	}
	if slice[i] == target {
		return
	}
	slicePkg.InsertInt64(slicePointer, i, target)
}

func SaveString(slicePointer *[]string, target string) {
	slice := *slicePointer
	if len(slice) == 0 {
		*slicePointer = []string{target}
		return
	}

	i := sort.Search(len(slice), func(i int) bool {
		return slice[i] >= target
	})

	if i >= len(slice) { // not found
		*slicePointer = append(slice, target)
		return
	}
	if slice[i] == target {
		return
	}
	slicePkg.InsertString(slicePointer, i, target)
}

// slice must be settable
func SaveValue(slice, target reflect.Value, fields ...string) {
	if slice.Len() == 0 {
		slice.Set(reflect.Append(slice, target))
		return
	}

	i := sort.Search(slice.Len(), func(i int) bool {
		return CompareValue(slice.Index(i), target, fields...) >= 0
	})

	if i >= slice.Len() { // not found
		slice.Set(reflect.Append(slice, target))
		return
	}
	if CompareValue(slice.Index(i), target, fields...) == 0 {
		slice.Index(i).Set(target)
		return
	}
	slicePkg.InsertValue(slice, i, target)
}
