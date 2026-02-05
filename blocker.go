package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type User struct {
	Name       string
	OpenAfter  string
	BlockAfter string
}

func main() {

	now := time.Now()

	user := createUser()

	timeToOpen := strings.Split(user.OpenAfter, ":")
	timeToClose := strings.Split(user.BlockAfter, ":")

	yourHourOpen, _ := strconv.Atoi(timeToOpen[0])
	yourMinuteOpen, _ := strconv.Atoi(timeToOpen[1])

	yourHourClose, _ := strconv.Atoi(timeToClose[0])
	yourMinuteClose, _ := strconv.Atoi(timeToClose[1])

	openTime := time.Date(now.Year(), now.Month(), now.Day(), yourHourOpen, yourMinuteOpen, 0, 0, now.Location())
	closeTime := time.Date(now.Year(), now.Month(), now.Day(), yourHourClose, yourMinuteClose, 0, 0, now.Location())

	// Create a ticker that ticks every 5 seconds
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		fmt.Println("Tick at", time.Now())
		if time.Now().Before(openTime) || time.Now().After(closeTime) {
			fmt.Printf("User %s its time to sleep! Blocking access to the internet...\n", user.Name)
			break
		}
	}
}

// Create a new user
func createUser() User {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Enter open time (HH:MM): ")
	scanner.Scan()
	openTime := scanner.Text()

	fmt.Print("Enter block time (HH:MM): ")
	scanner.Scan()
	blockTime := scanner.Text()

	return User{
		Name:       name,
		OpenAfter:  openTime,
		BlockAfter: blockTime,
	}
}
