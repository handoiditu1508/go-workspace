package sample

import (
	"fmt"
	"math/rand"
)

func RunSliceSample() {
	// define slice
	fmt.Println("define slices")
	var numbers []int
	numbers = make([]int, 5)
	matrix := make([][]int, 3*3)

	// insert data
	fmt.Println(">>>>>insert slice data")
	for i := 0; i < 5; i++ {
		numbers[i] = rand.Intn(100)
	}

	// insert matrix data
	fmt.Println(">>>>>insert slice matrix data")
	for i := 0; i < 3; i++ {
		matrix[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			matrix[i][j] = rand.Intn(100)
		}
	}

	// display data
	fmt.Println(">>>>>display sclice data")
	for i := 0; i < 5; i++ {
		fmt.Println(numbers[i])
	}

	// display matrix data
	fmt.Println(">>>>>display sclice 2D data")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(matrix[i][j])
		}
	}
}
