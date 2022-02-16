package main

import "fmt"

func xpto(a *int) int {
	*a = 100
	return *a
}

func main() {

	b := 10
	fmt.Println(xpto(&b))
	fmt.Println(b)

	x := 10
	fmt.Println(&x)

	y := &x

	*y = 8

	fmt.Println(x)

	var z *int = &x

	fmt.Println(*z)

}
