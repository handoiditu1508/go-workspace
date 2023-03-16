package sample

import "fmt"

func RunSlice2Sample() {
	// slice is a reference to a contiguous segment of an array
	// array is value type, slice is reference type
	var intArray = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var intSlice = intArray[2:5]                                                                  // index 2 to 4, size = 3
	fmt.Println("Length: ", len(intSlice), " Contents: ", intSlice, " Capacity: ", cap(intSlice)) // Length:  3  Contents:  [1 2 3]  Capacity:  8
	// slice length depend on where slice's start index and end index

	// update {intSlice} will affect {intArray}
	//intSlice[1] = 33

	intSlice2 := append(intSlice, 10)
	fmt.Println(intArray)
	fmt.Println("Length: ", len(intSlice), " Contents: ", intSlice, " Capacity: ", cap(intSlice))
	fmt.Println("Length: ", len(intSlice2), " Contents: ", intSlice2, " Capacity: ", cap(intSlice2))

	intSlice2 = append(intSlice, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	fmt.Println(intArray)
	fmt.Println("Length: ", len(intSlice), " Contents: ", intSlice, " Capacity: ", cap(intSlice))
	fmt.Println("Length: ", len(intSlice2), " Contents: ", intSlice2, " Capacity: ", cap(intSlice2))
}
