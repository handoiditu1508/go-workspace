package exercise2

import "unicode"

func InvokeExercise1(input string) string {
	numbers := ""

	for _, character := range input {
		if unicode.IsDigit(character) {
			numbers += string(character)
		}
	}

	output := ""
	for i := 0; i < len(numbers); {
		if i+3 == len(numbers)-1 {
			output += numbers[i:i+2] + "-" + numbers[i+2:i+4] + "-"
			i = len(numbers)
			break
		} else {
			nextIndex := i + min(3, len(numbers)-i)
			output += numbers[i:nextIndex] + "-"
			i = nextIndex
		}
	}

	if len(output) > 0 {
		output = output[:len(output)-1]
	}

	return output
}
