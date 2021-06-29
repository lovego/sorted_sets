package sorted_sets

import (
	"fmt"
	"reflect"
)

func ExampleRemoveInt64() {
	s := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range []int64{5, 3, 8, 0} {
		s = RemoveInt64(s, v)
	}
	fmt.Println(s)
	fmt.Println(RemoveInt64(nil, 3), RemoveInt64([]int64{}, 3), RemoveInt64([]int64{3}, 3))
	// Output: [1 2 4 6 7 9]
	// [] [] []
}

func ExampleRemoveString() {
	s := []string{"a1", "a2", "a3", "a4", "a5", "a6", "a7", "a8", "a9"}
	for _, v := range []string{"a5", "a3", "a8", "a0"} {
		s = RemoveString(s, v)
	}
	fmt.Println(s)
	fmt.Println(
		RemoveString(nil, "3"), RemoveString([]string{}, "3"), RemoveString([]string{"3"}, "3"),
	)
	// Output: [a1 a2 a4 a6 a7 a9]
	// [] [] []
}

func ExampleRemoveValue_int64() {
	s := reflect.ValueOf([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	for _, v := range []int64{5, 3, 8, 0} {
		s = RemoveValue(s, reflect.ValueOf(v))
	}
	fmt.Println(s)
	fmt.Println(RemoveValue(reflect.ValueOf([]int64(nil)), reflect.ValueOf(3)))
	fmt.Println(RemoveValue(reflect.Value{}, reflect.ValueOf(3)))
	// Output:
	// [1 2 4 6 7 9]
	// []
	// <invalid reflect.Value>
}

func ExampleRemoveValue_intField() {
	var s = reflect.ValueOf([]testStruct{
		{1, `name1`}, {2, `name2`}, {3, `name3`}, {4, `name4`}, {5, `name5`},
		{6, `name6`}, {7, `name7`}, {8, `name8`}, {9, `name9`},
	})
	for _, v := range []testStruct{
		{Id: 4}, {Id: 2}, {Id: 4}, {Id: 7}, {Id: 0},
	} {
		s = RemoveValue(s, reflect.ValueOf(v), `Id`)
	}
	fmt.Println(s)
	// Output:
	// [{1 name1} {3 name3} {5 name5} {6 name6} {8 name8} {9 name9}]
}

func ExampleRemoveValue_stringField() {
	var s = reflect.ValueOf([]testStruct{
		{1, `name1`}, {2, `name2`}, {3, `name3`}, {4, `name4`}, {5, `name5`},
		{6, `name6`}, {7, `name7`}, {8, `name8`}, {9, `name9`},
	})
	for _, v := range []testStruct{
		{Name: `name4`}, {Name: `name2`}, {Name: `name4`}, {Name: `name7`}, {Name: `name0`},
	} {
		s = RemoveValue(s, reflect.ValueOf(v), `Name`)
	}
	fmt.Println(s)
	// Output:
	// [{1 name1} {3 name3} {5 name5} {6 name6} {8 name8} {9 name9}]
}
