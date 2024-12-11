package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)

	if bilangan <= 0 {
		fmt.Println("Tidak valid")
		return
	}

	fmt.Println()

	for bilangan > 0 {
		digit := bilangan % 10
		fmt.Println(digit)
		bilangan = bilangan / 10
	}
}
