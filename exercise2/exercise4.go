package exercise2

import "strings"

var checkersBoard []string
var output int

func InvokeExercise4(input []string) int {
	// find 'O' position
	var x, y int
	for index, value := range input {
		y = strings.Index(value, "O")
		if y != -1 {
			x = index
			break
		}
	}

	checkersBoard = input
	output = 0
	deepSearch4(x, y, 0)
	return output
}

func deepSearch4(x, y, point int) {
	if x > 1 {
		if y > 1 {
			if string(checkersBoard[x-2][y-2]) == "." {
				if string(checkersBoard[x-1][y-1]) == "X" {
					deepSearch4(x-2, y-2, point+1)
				}
			}
		}

		if y < len(checkersBoard)-2 {
			if string(checkersBoard[x-2][y+2]) == "." {
				if string(checkersBoard[x-1][y+1]) == "X" {
					deepSearch4(x-2, y+2, point+1)
				}
			}
		}
	}

	if point > output {
		output = point
	}
}
