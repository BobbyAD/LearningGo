package main

import "fmt"

// This has turned in to more of a test file for trying different stuff.

func bar() {
	fmt.Println("bar")
}

func main() {
	fmt.Println("Hey, world")
	foo()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	numbers := [6]int{1, 2, 3, 5}

	for i, x := range numbers {
		fmt.Println(i, x)
	}

	fmt.Println("asdf", 1, "abc")

	var x = 3

	fmt.Println(x)

	y := 5

	fmt.Println(y)

	z := "asdf"

	fmt.Println(z)

	var a int

	fmt.Println(a)

	a = 42

	fmt.Printf("a's type: %T\n", a)

	type hotdog int

	var b hotdog

	b = 32

	fmt.Printf("b's type: %T\n", b)

	fmt.Println("Conversion, not casting:")

	a = int(b)

	fmt.Println(b)

	bar()
}

func foo() {
	fmt.Println("foo")
}
