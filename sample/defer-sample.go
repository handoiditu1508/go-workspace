package sample

import "fmt"

func RunDeferSample() {
	// Defer is commonly used to simplify functions that perform various clean-up actions
	// defer statement pushes function call onto a stack
	// the stack of saved calls is executed after the surrounding function returns

	fmt.Println("===test1===")
	test1()

	fmt.Println("===test2===")
	test2()

	fmt.Println("\n===test3===")
	fmt.Println(test3())
}

// A deferred function’s arguments are evaluated when the defer statement is evaluated
// The deferred call will print “0” after the function returns
func test1() {
	i := 0
	defer fmt.Println(i)
	i++
}

// Deferred function calls are executed in Last In First Out order after the surrounding function returns
// This function prints "3210"
func test2() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

// Deferred functions may read and assign to the returning function’s named return values
// this function returns 2
func test3() (i int) {
	defer func() { i++ }()
	return 1
}
