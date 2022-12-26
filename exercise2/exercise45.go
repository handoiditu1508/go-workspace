package exercise2

func InvokeExercise45(input string) int {
	// get all possible substrings
	dictionary := make(map[string]int)
	for i := 0; i < len(input); i++ {
		for j := i; j < len(input); j++ {
			subString := input[i : j+1]
			dictionary[subString] += 1

		}
	}

	smallestAppearTimes := 200
	shortestString := ""
	for subString, appearTimes := range dictionary {
		if appearTimes < smallestAppearTimes {
			smallestAppearTimes = appearTimes
			shortestString = subString
		} else if appearTimes == smallestAppearTimes && len(subString) < len(shortestString) {
			shortestString = subString
		}
	}

	return len(shortestString)
}
