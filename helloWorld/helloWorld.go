package main

import "fmt"

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

	bar()
}

func foo() {
	fmt.Println("foo")
}
