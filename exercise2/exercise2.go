// 3 4 5 4 5 6 7
// 3 4 4 5 5 6 7

package exercise2

import "sort"

func InvokeExercise2(input []int) int {
	// clone input
	inputClone := make([]int, len(input))
	copy(inputClone, input)
	sort.Ints(inputClone)

	countDifferent := 0

	for i := 0; i < len(input) && countDifferent < 3; i++ {
		if input[i] != inputClone[i] {
			countDifferent++
		}
	}

	if countDifferent > 2 {
		return 0
	} else if countDifferent == 2 {
		return 2
	}
	return len(input)
}
