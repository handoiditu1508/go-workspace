package sample

import "fmt"

func RunWhileIterationSample() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	j := 5
	for ; j < 10; j++ {
		fmt.Println(j)
	}
}
