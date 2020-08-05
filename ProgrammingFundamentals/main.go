package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Programming Fundamentals\n------------------------\n")
	fmt.Printf("Types\n-----\n")
	types()
}

// testing types

var b bool

func types() {
	fmt.Printf("b is type %T, its value is %v\n", b, b)
	x := 7
	z := 8
	fmt.Println("x == z:", x == z)

	i := 42
	f := 42.45234
	fmt.Printf("i: %T\nf: %T\n", i, f)

	var testInt int8
	// testInt = -129 -- would cause an error, overflows 8-bits signed
	// testInt = 128 -- Same error
	testInt = 127
	fmt.Println(testInt)
	fmt.Println(runtime.GOARCH)

	s := "Heyo, boyo"
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	// chars are runes

	// test_string := 'test'
	// fmt.Println(test_string)

	test_c := 't'
	fmt.Println(test_c)
	fmt.Printf("%T\n", test_c)
}
