package exercise2

func InvokeExercise44(n, k int) int {
	kTemp := k
	count := 0

	for i := n; i > 0; {
		if i < kTemp {
			kTemp -= i
			i--
		} else {
			kTemp = 0
			i = 0
		}
		count++
	}

	if kTemp != 0 {
		return -1
	}
	return count
}

// 1 2 3 4 5 6 7 8 9 10
// 44 - 10 = 34
// 34 - 9 = 25
// 25 - 8 = 17
// 17 - 7 = 10
// 10 - 6 = 4
// 4 - 4
