package sample

import "fmt"

func RunPointerSample() {
	x := 10

	var p1 *int
	p2 := new(int)

	p1 = new(int)

	*p1 = x
	p2 = &x

	x = 20

	fmt.Println("===value===")
	fmt.Println(x)  // value of x
	fmt.Println(&x) // address of x

	fmt.Println("===pointer 1===")
	fmt.Println(*p1) // value of pointed memory
	fmt.Println(p1)  // address of pointed memory

	fmt.Println("===pointer 2===")
	fmt.Println(*p2) // value of pointed memory
	fmt.Println(p2)  // address of pointed memory
}
