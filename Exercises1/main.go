package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
}

func exercise1() {

	// assign these values to variables x y z
	// 42, "James Bond", true

	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)

}

// package level variables
var a int
var b string
var c bool

func exercise2() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func exercise3() {

	// assigning them
	a = 42
	b = "James Bond"
	c = true

	comboStr := fmt.Sprintf("%v\t %v\t %v", a, b, c)
	fmt.Println(comboStr)
}

// creating own type
type myType int

var typeVar myType

func exercise4() {
	fmt.Printf("\n\nTYPE\n----\n")
	fmt.Println(typeVar)

	fmt.Printf("%T\n", typeVar)

	typeVar = 42

	fmt.Println(typeVar)
}

// exercise 5

var i int

func exercise5() {
	i = int(typeVar)
	fmt.Printf("i's value and type: (%v, %t)", i, i)
}
