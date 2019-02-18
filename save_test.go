package sorted_sets

import (
	"fmt"
	"reflect"
)

func ExampleSave_int64() {
	var s = reflect.ValueOf([]int64{})
	for _, v := range []int64{3, 6, 2, 8, 9, 1, 5, 2, 4, 6, 7} {
		s = Save(s, reflect.ValueOf(v))
	}
	fmt.Println(s)
	// Output: [1 2 3 4 5 6 7 8 9]
}

func ExampleSave_string() {
	var s = reflect.ValueOf([]string{})
	for _, v := range []string{"aa", "ba", "ab", "cc", "a", "ab", "b"} {
		s = Save(s, reflect.ValueOf(v))
	}
	fmt.Println(s)
	// Output: [a aa ab b ba cc]
}

type testStruct struct {
	Id   int64
	Name string
}

func ExampleSave_intField() {
	var s = reflect.ValueOf([]testStruct{})
	for _, v := range []testStruct{
		{Id: 3, Name: `name3`}, {Id: 6, Name: `name_6`}, {Id: 2, Name: `name_2`},
		{Id: 5, Name: `name5`}, {Id: 4, Name: `name4`}, {Id: 7, Name: `name7`},
		{Id: 2, Name: `name-2`}, {Id: 9, Name: `name9`}, {Id: 1, Name: `name1`},
		{Id: 2, Name: `name2`}, {Id: 8, Name: `name8`}, {Id: 6, Name: `name6`},
	} {
		s = Save(s, reflect.ValueOf(v), `Id`)
	}
	fmt.Println(s)
	// Output:
	// [{1 name1} {2 name2} {3 name3} {4 name4} {5 name5} {6 name6} {7 name7} {8 name8} {9 name9}]
}

func ExampleSave_stringField() {
	var s = reflect.ValueOf([]testStruct{})
	for _, v := range []testStruct{
		{Id: 3, Name: `name3`}, {Id: 16, Name: `name6`}, {Id: 12, Name: `name2`},
		{Id: 5, Name: `name5`}, {Id: 4, Name: `name4`}, {Id: 7, Name: `name7`},
		{Id: 22, Name: `name2`}, {Id: 9, Name: `name9`}, {Id: 1, Name: `name1`},
		{Id: 2, Name: `name2`}, {Id: 8, Name: `name8`}, {Id: 6, Name: `name6`},
	} {
		s = Save(s, reflect.ValueOf(v), `Name`)
	}
	fmt.Println(s)
	// Output:
	// [{1 name1} {2 name2} {3 name3} {4 name4} {5 name5} {6 name6} {7 name7} {8 name8} {9 name9}]
}

func ExampleSave_invalidSlice() {
	var s = reflect.Value{}
	for _, v := range []int64{3, 6, 2, 8, 9, 1, 5, 2, 4, 6, 7} {
		s = Save(s, reflect.ValueOf(v))
	}
	fmt.Println(s)
	// Output: [1 2 3 4 5 6 7 8 9]
}
