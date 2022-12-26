package sample

import (
	"fmt"
	"math/rand"
)

func RunMapSample() {
	// define map
	fmt.Println("define map")
	products := make(map[string]int)
	customers := make(map[string]int)

	// insert data
	fmt.Println(">>>>>insert map data")
	products["product1"] = rand.Intn(100)
	products["product2"] = rand.Intn(100)

	customers["customer1"] = rand.Intn(100)
	customers["customer2"] = rand.Intn(100)

	// display data
	fmt.Println(">>>>>display map data")
	fmt.Println(products["product1"])
	fmt.Println(products["product2"])
	fmt.Println(customers["customer1"])
	fmt.Println(customers["customer2"])
}
