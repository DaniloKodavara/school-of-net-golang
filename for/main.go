package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}

	for {
		x++

		if x == 100 {
			break
		}
	}

}
