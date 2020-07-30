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

	bar()
}

func foo() {
	fmt.Println("foo")
}
