package sample

import "fmt"

func RunAddressSample() {
	var x int

	// print address of x
	fmt.Println(&x) // 0xc00000a098

	x = 10

	fmt.Println(&x) // 0xc00000a098

	x = 20

	fmt.Println(&x) // 0xc00000a098
}
