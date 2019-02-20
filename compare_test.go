package sorted_sets

import (
	"fmt"
	"reflect"
)

func ExampleCompareValue_int() {
	for _, t := range []struct {
		a, b int
	}{
		{7, 8}, {7, 7}, {8, 7},
	} {
		fmt.Println(CompareValue(reflect.ValueOf(t.a), reflect.ValueOf(t.b)))
	}
	// Output:
	// -1
	// 0
	// 1
}

func ExampleCompareValue_string() {
	for _, t := range []struct {
		a, b string
	}{
		{"a", "b"}, {"a", "ab"}, {"", "a"},
		{"a", "a"}, {"ab", "ab"}, {"", ""},
		{"b", "a"}, {"ab", "a"}, {"a", ""},
	} {
		fmt.Println(CompareValue(reflect.ValueOf(t.a), reflect.ValueOf(t.b)))
	}
	// Output:
	// -1
	// -1
	// -1
	// 0
	// 0
	// 0
	// 1
	// 1
	// 1
}

type T struct {
	s string
	i int64
}

var a1, a2 = T{"a", 1}, T{"a", 2}
var b1, b2 = T{"b", 1}, T{"b", 2}

func ExampleCompareValue_case1() {
	for _, t := range []struct {
		a, b T
	}{
		{a1, a2}, {a1, b1}, {a1, b2}, {a2, b1},
		{a1, a1}, {b2, b2},
		{a2, a1}, {b1, a1}, {b1, a2}, {b2, b1},
	} {
		fmt.Println(CompareValue(reflect.ValueOf(t.a), reflect.ValueOf(t.b), "s", "i"))
	}

	// Output:
	// -1
	// -1
	// -1
	// -1
	// 0
	// 0
	// 1
	// 1
	// 1
	// 1
}

func ExampleCompareValue_case2() {
	for _, t := range []struct {
		a, b T
	}{
		{a1, a2}, {a1, b1}, {a1, b2}, {b1, a2},
		{a1, a1}, {b2, b2},
		{a2, a1}, {b1, a1}, {a2, b1}, {b2, a1},
	} {
		fmt.Println(CompareValue(reflect.ValueOf(t.a), reflect.ValueOf(t.b), "i", "s"))
	}

	// Output:
	// -1
	// -1
	// -1
	// -1
	// 0
	// 0
	// 1
	// 1
	// 1
	// 1
}

func ExampleCompareValue_intField() {
	for _, t := range []struct {
		a, b T
	}{
		{a1, a2}, {a1, b2}, {b1, a2}, {b1, b2},
		{a1, b1}, {a2, b2},
		{a2, a1}, {a2, b1}, {b2, a1}, {b2, b1},
	} {
		fmt.Println(CompareValue(reflect.ValueOf(t.a), reflect.ValueOf(t.b), "i"))
	}

	// Output:
	// -1
	// -1
	// -1
	// -1
	// 0
	// 0
	// 1
	// 1
	// 1
	// 1
}

func ExampleCompareValue_stringField() {
	for _, t := range []struct {
		a, b T
	}{
		{a1, b1}, {a1, b2}, {a2, b1}, {a2, b2},
		{a1, a2}, {b1, b2},
		{b1, a1}, {b1, a2}, {b2, a1}, {b2, a2},
	} {
		fmt.Println(CompareValue(reflect.ValueOf(t.a), reflect.ValueOf(t.b), "s"))
	}

	// Output:
	// -1
	// -1
	// -1
	// -1
	// 0
	// 0
	// 1
	// 1
	// 1
	// 1
}

func ExampleCompareValue_nil() {
	var p1, p2 *int
	fmt.Println(CompareValue(reflect.ValueOf(p1), reflect.ValueOf(p2)))
	fmt.Println(CompareValue(reflect.ValueOf(p1), reflect.ValueOf(10)))
	fmt.Println(CompareValue(reflect.ValueOf(10), reflect.ValueOf(p2)))
	var i = 3
	p1 = &i
	fmt.Println(CompareValue(reflect.ValueOf(p1), reflect.ValueOf(p2)))
	fmt.Println(CompareValue(reflect.ValueOf(p2), reflect.ValueOf(p1)))
	fmt.Println(CompareValue(reflect.ValueOf(p1), reflect.ValueOf(2)))
	fmt.Println(CompareValue(reflect.ValueOf(p1), reflect.ValueOf(3)))
	fmt.Println(CompareValue(reflect.ValueOf(p1), reflect.ValueOf(4)))

	// Output:
	// 0
	// -1
	// 1
	// 1
	// -1
	// 1
	// 0
	// -1

}
