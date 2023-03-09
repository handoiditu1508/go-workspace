package exercise2

import (
	"strconv"
	"strings"
)

func InvokeFullGameCount(A string, B string) int {
	time1 := strings.Split(A, ":")
	startH, _ := strconv.Atoi(time1[0])
	startM, _ := strconv.Atoi(time1[1])

	time2 := strings.Split(B, ":")
	endH, _ := strconv.Atoi(time2[0])
	endM, _ := strconv.Atoi(time2[1])
	endHCap := endH
	if endM > 0 {
		endHCap++
	}

	if A == B {
		return 0
	}

	result := 0
	if startH <= endH {
		result = (endHCap - startH) * 4
		result -= (startM / 15)
		if startM%15 != 0 {
			result--
		}
		result -= (4 - (endM / 15))
	} else if startH > endH {
		result = ((24 - startH) * 4) + (endHCap * 4)
		result -= (startM / 15)
		if startM%15 != 0 {
			result--
		}
		if endM != 0 {
			result -= (4 - (endM / 15))
		}
	}
	return result
}
