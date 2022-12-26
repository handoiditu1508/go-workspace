package sample

import "fmt"

func RunForIterationSample() {
	var i int
	for i = 0; i < 5; i++ {
		fmt.Println(i)
	}

	for j := 5; j < 11; j++ {
		fmt.Println(j)
	}

	// foreach
	numbers := []int{2, 4, 6, 8, 10}
	for _, number := range numbers {
		fmt.Println(number)
	}
}
