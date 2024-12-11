package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("x = ")
	fmt.Scanln(&x)
	fmt.Print("y = ")
	fmt.Scanln(&y)

	if x <= 0 || y <= 0 {
		fmt.Println("tidak valid")
		return
	}
	result := 0
	for x >= y {
		x = x - y
		result++
	}

	fmt.Println(result)
}
