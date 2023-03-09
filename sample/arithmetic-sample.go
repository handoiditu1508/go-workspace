package sample

import "fmt"

func RunArithmeticSample() {
	// declare variables
	var a, b int
	// assign values
	a = 5
	b = 10
	// arithmetic operation
	// addition
	c := a + b
	fmt.Printf("%d + %d = %d \n", a, b, c) // 5 + 10 = 15
	// subtraction
	d := a - b
	fmt.Printf("%d - %d = %d \n", a, b, d) // 5 - 10 = -5
	// division
	e := float32(a) / float32(b)
	fmt.Printf("%d / %d = %.2f \n", a, b, e) // 5 / 10 = 0.50
	// multiplication
	f := a * b
	fmt.Printf("%d * %d = %d \n", a, b, f) // 5 * 10 = 50

}
