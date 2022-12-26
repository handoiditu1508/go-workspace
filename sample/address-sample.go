package sample

import "fmt"

func RunAddressSample() {
	var x int

	fmt.Println(&x) // print address of x

	x = 10

	fmt.Println(&x)

	x = 20

	fmt.Println(&x)
}
