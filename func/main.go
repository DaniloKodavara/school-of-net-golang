package main

import "fmt"

func funcName(a int) int {
	return a * a
}

func nameReturn(a string) (x string) {
	x = a

	return
}

func moreReturns(a string, b string) (string, string) {
	return a, b
}

func variadicFunc(x ...int) int {
	var res int
	for _, i := range x {
		res += i
	}
	return res
}

func funcInsideFunc() func() int {
	x := 10

	return func() int {
		return x * x
	}
}

func main() {
	fmt.Println(funcName(10))

	fmt.Println(nameReturn("Danilo"))

	x, y := moreReturns("Danilo", "Yukio")

	fmt.Printf("%s %s \n", x, y)

	fmt.Println(variadicFunc(1, 2, 5, 10))

	z := 0

	add := func() int {
		z += 2
		return z
	}

	fmt.Println(add())
	fmt.Println(add())

}
