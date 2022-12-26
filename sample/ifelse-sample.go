package sample

import "fmt"

func RunIfElseSample() {
	var (
		a = 5
		b = 8
	)

	if a > b || a-b < a {
		fmt.Println("conditional-->a>b || a-b<a")
	} else {
		fmt.Println("..another")
	}
}
