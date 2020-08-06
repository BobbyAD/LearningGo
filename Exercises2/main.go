package main

import (
	"fmt"
)

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
}

func exercise1() {
	x := 12
	fmt.Printf("%d\t%b\t%x", x, x, x)
}

func exercise2() {
	a := 10 == 10
	b := 10 <= 9
	c := 10 >= 9
	d := 10 != 10
	e := 10 < 10
	f := 10 > 9

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

const (
	a     = "James Bond"
	b int = 100
)

func exercise3() {
	fmt.Printf("%T\t%v\n%T%v\n", a, a, b, b)
}

func exercise4() {
	a := 100
	b := a << 1
	fmt.Printf("%b\t%v\t%x\n", a, a, a)
	fmt.Printf("%b\t%v\t%x\n", b, b, b)
}

func exercise5() {
	lit := `\n\n\n\n`
	fmt.Println(lit)
}

const (
	y1 = 2020 + iota
	y2 = 2020 + iota
	y3 = 2020 + iota
	y4 = 2020 + iota
)

func exercise6() {
	fmt.Println(y1, y2, y3, y4)
}
