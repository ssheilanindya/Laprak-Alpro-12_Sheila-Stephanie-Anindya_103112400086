package main

import "fmt"

func main() {
	var username, password string

	for {
		correctUsername := "Admin"
		correctPassword := "Admin"
		failedAttempts := 0

		fmt.Print("username: ")
		fmt.Scanln(&username)

		fmt.Print("password: ")
		fmt.Scanln(&password)

		if username == correctUsername && password == correctPassword {
			fmt.Printf("%d percobaan gagal login.\n", failedAttempts)
			break
		} else {
			fmt.Println("gagal login.")
			failedAttempts++
		}
	}
}
