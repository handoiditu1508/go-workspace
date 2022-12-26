package sample

import (
	"fmt"
	"math/rand"
)

func RunArraySample() {
	// define array
	fmt.Println("define arrays")
	var numbers [5]int
	var cities [5]string
	var matrix [3][3]int

	// insert data
	fmt.Println(">>>>>insert array data")
	for i := 0; i < 5; i++ {
		numbers[i] = i
		cities[i] = fmt.Sprintf("City %d", i)
	}

	// insert matrix data
	fmt.Println(">>>>>insert matrix data")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = rand.Intn(100)
		}
	}

	// display data
	fmt.Println(">>>>>display array data")
	for i := 0; i < 5; i++ {
		fmt.Println(numbers[i])
		fmt.Println(cities[i])
	}

	// display matrix data
	fmt.Println(">>>>>display matrix data")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(matrix[i][j])
		}
	}
}