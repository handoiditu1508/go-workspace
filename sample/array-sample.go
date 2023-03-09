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
	// in Go, array is value type, slice is reference type
	var numbers2 = numbers
	numbers2[0] = 26
	for i := 0; i < 5; i++ {
		fmt.Println(numbers[i])
		fmt.Println(cities[i])
		fmt.Println(numbers2[i])
	}
	// 0
	// City 0
	// 26
	// 1
	// City 1
	// 1
	// 2
	// City 2
	// 2
	// 3
	// City 3
	// 3
	// 4
	// City 4
	// 4

	// display matrix data
	fmt.Println(">>>>>display matrix data")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(matrix[i][j])
		}
	}
	// 81
	// 87
	// 47
	// 59
	// 81
	// 18
	// 25
	// 40
	// 56
}
