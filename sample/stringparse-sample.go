package sample

import (
	"fmt"
	"strconv"
)

func RunStringParseSample() {
	str1 := "5"
	str2 := "2.8769"

	var err error
	var int1 int64
	int1, err = strconv.ParseInt(str1, 10, 32)
	if err == nil {
		fmt.Println(int1)
	} else {
		fmt.Println(err)
	}

	var float2 float64
	float2, err = strconv.ParseFloat(str2, 64)
	if err == nil {
		fmt.Println(float2)
	} else {
		fmt.Println(err)
	}
}
