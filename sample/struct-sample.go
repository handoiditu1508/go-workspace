package sample

import (
	"fmt"
	"time"
)

type Employee struct {
	id      int
	name    string
	country string
	created time.Time
}

func RunStructSample() {
	// declare variable
	var emp Employee
	newEmp := new(Employee)
	newEmp2 := Employee{1, "Employee 1", "VN", time.Now()}
	newEmp3 := emp // pass by value, emp changes won't affect newEmp3

	// set values
	emp.id = 2
	emp.name = "Employee 2"
	emp.country = "DE"
	emp.created = time.Now()

	newEmp.id = 5
	newEmp.name = "Employee 5"
	newEmp.country = "UK"
	newEmp.created = time.Now()

	// display
	fmt.Println("=====================")
	fmt.Println(emp)
	fmt.Println("=====================")
	fmt.Println(newEmp.id)
	fmt.Println(newEmp.name)
	fmt.Println(newEmp.country)
	fmt.Println(newEmp.created)
	fmt.Println("=====================")
	fmt.Println(newEmp2)
	fmt.Println(newEmp3)
}
