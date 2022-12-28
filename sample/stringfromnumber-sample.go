package sample

import "fmt"

func RunStringFromNumberSample() {
	num1 := 8
	num2 := 5.872

	str1 := fmt.Sprintf("%d", num1)
	fmt.Println(str1)

	str2 := fmt.Sprintf("%f", num2)
	fmt.Println(str2)
}
