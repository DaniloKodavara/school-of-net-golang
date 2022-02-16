package scope

import "fmt"

var Y int = 20

func main() {
	x := 10
	fmt.Println(x)
	fmt.Println(z)
}

func PrintY() {
	fmt.Println(Y)
}
