package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/ivan-tripplex/safety-lock/utils"
)

func main() {
	// Initialize the user store and load users from file
	store := utils.NewUserStore("users.json")
	store.LoadFromFile()

	var runTime, closeTime time.Time

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your Login: ")
	scanner.Scan()
	name := scanner.Text()

	// Check if the user exists
	user, err := store.UserByUsername(name)
	if err == nil {
		fmt.Print("Enter your Password: ")
		scanner.Scan()
		password := scanner.Text()
		for utils.CheckPassword(user.PasswordHash, password) != nil {
			fmt.Println("Incorrect password. Exiting...")
			fmt.Print("Enter your Password: ")
			scanner.Scan()
			password = scanner.Text()
		}

		fmt.Println("Login successful!")
		runTime, closeTime = store.AddToTime(user.OpenAfter, user.BlockAfter)
	} else {
		fmt.Print("Enter your Password: ")
		scanner.Scan()
		password := scanner.Text()

		fmt.Print("Enter open time (HH:MM): ")
		scanner.Scan()
		openTime := scanner.Text()

		fmt.Print("Enter block time (HH:MM): ")
		scanner.Scan()
		blockTime := scanner.Text()

		store.CreateUser(name, password, openTime, blockTime)
		fmt.Println("User created successfully!")
		runTime, closeTime = store.AddToTime(openTime, blockTime)
	}

	// Create a ticker that ticks every 5 seconds
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		fmt.Println("Tick at", time.Now())
		if time.Now().Before(runTime) || time.Now().After(closeTime) {
			fmt.Printf("Its time to sleep! Blocking access to the internet...\n")
			utils.BlockWebsites()
			break
		} else {
			utils.UnblockWebsites()
			break
		}
	}
}
