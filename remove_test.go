package sorted_sets

import (
	"fmt"
	"reflect"
)

func ExampleRemoveInt64() {
	s := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range []int64{5, 3, 8, 0} {
		RemoveInt64(&s, v)
	}
	fmt.Println(s)

	ptr := (*[]int64)(nil)
	RemoveInt64(ptr, 3)
	fmt.Println(ptr)

	s = []int64{}
	RemoveInt64(&s, 3)
	fmt.Println(s)

	s = []int64{3}
	RemoveInt64(&s, 3)
	fmt.Println(s)
	// Output:
	// [1 2 4 6 7 9]
	// <nil>
	// []
	// []
}

func ExampleRemoveString() {
	s := []string{"a1", "a2", "a3", "a4", "a5", "a6", "a7", "a8", "a9"}
	for _, v := range []string{"a5", "a3", "a8", "a0"} {
		RemoveString(&s, v)
	}
	fmt.Println(s)

	ptr := (*[]string)(nil)
	RemoveString(ptr, "3")
	fmt.Println(ptr)

	s = []string{}
	RemoveString(&s, "3")
	fmt.Println(s)

	s = []string{"3"}
	RemoveString(&s, "3")
	fmt.Println(s)
	// Output:
	// [a1 a2 a4 a6 a7 a9]
	// <nil>
	// []
	// []
}

func ExampleRemoveValue_int64() {
	slice := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := reflect.ValueOf(&slice).Elem()
	for _, v := range []int64{5, 3, 8, 0} {
		RemoveValue(s, reflect.ValueOf(v))
	}
	fmt.Println(s)

	slice = []int64(nil)
	s = reflect.ValueOf(&slice).Elem()
	RemoveValue(s, reflect.ValueOf(3))
	fmt.Println(s)

	s = reflect.Value{}
	RemoveValue(s, reflect.ValueOf(3))
	fmt.Println(s)
	// Output:
	// [1 2 4 6 7 9]
	// []
	// <invalid reflect.Value>
}

func ExampleRemoveValue_intField() {
	slice := []testStruct{
		{1, `name1`}, {2, `name2`}, {3, `name3`}, {4, `name4`}, {5, `name5`},
		{6, `name6`}, {7, `name7`}, {8, `name8`}, {9, `name9`},
	}
	var s = reflect.ValueOf(&slice).Elem()
	for _, v := range []testStruct{
		{Id: 4}, {Id: 2}, {Id: 4}, {Id: 7}, {Id: 0},
	} {
		RemoveValue(s, reflect.ValueOf(v), `Id`)
	}
	fmt.Println(s)
	// Output:
	// [{1 name1} {3 name3} {5 name5} {6 name6} {8 name8} {9 name9}]
}

func ExampleRemoveValue_stringField() {
	slice := []testStruct{
		{1, `name1`}, {2, `name2`}, {3, `name3`}, {4, `name4`}, {5, `name5`},
		{6, `name6`}, {7, `name7`}, {8, `name8`}, {9, `name9`},
	}
	var s = reflect.ValueOf(&slice).Elem()
	for _, v := range []testStruct{
		{Name: `name4`}, {Name: `name2`}, {Name: `name4`}, {Name: `name7`}, {Name: `name0`},
	} {
		RemoveValue(s, reflect.ValueOf(v), `Name`)
	}
	fmt.Println(s)
	// Output:
	// [{1 name1} {3 name3} {5 name5} {6 name6} {8 name8} {9 name9}]
}
