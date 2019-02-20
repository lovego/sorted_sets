package sorted_sets

import (
	"fmt"
	"reflect"
)

func ExampleRemove_int64() {
	s := reflect.ValueOf([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	for _, v := range []int64{5, 3, 8, 0} {
		s = RemoveValue(s, reflect.ValueOf(v))
	}
	fmt.Println(s)
	// Output: [1 2 4 6 7 9]
}

func ExampleRemove_intField() {
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

func ExampleRemove_stringField() {
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
