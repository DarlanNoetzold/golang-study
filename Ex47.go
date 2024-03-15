package main

import "fmt"

func alem() func() int {
	x := 0
	y := 0
	fmt.Println(x, y)
	return func() int {
		x++
		y++
		fmt.Println(x, y)
		return x + y
	}
}

func main() {
	a := alem()
	b := alem()
	c := alem()

	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

}