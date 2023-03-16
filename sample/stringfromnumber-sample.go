package sample

import (
	"fmt"
	"strconv"
)

func RunStringFromNumberSample() {
	num1 := 8
	num2 := 5.872

	str1 := fmt.Sprintf("%d", num1)
	fmt.Println(str1)

	str2 := fmt.Sprintf("%.2f", num2)
	fmt.Println(str2)

	str3 := strconv.Itoa(num1)
	fmt.Println(str3)
}
