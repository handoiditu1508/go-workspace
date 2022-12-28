package sample

import "fmt"

func RunSlice2Sample() {
	// slice is a reference to a contiguous segment of an array
	// array is value type, slice is reference type
	var intArray = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var intSlice = intArray[2:5]                                                                  // index 2 to 4, size = 3
	fmt.Println("Length: ", len(intSlice), " Contents: ", intSlice, " Capacity: ", cap(intSlice)) // Length:  3  Contents:  [1 2 3]  Capacity:  8
	// slice length can range from 0 to lenght of the array, depend on where slice's start index
}
