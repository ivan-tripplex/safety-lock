package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a ticker that ticks every 5 seconds

	const maxRuntime = 10 * time.Second
	now := time.Now()

	yourHour := 20
	yourMinute := 1
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	closeTime := time.Date(now.Year(), now.Month(), now.Day(), yourHour, yourMinute, 0, 0, now.Location())

	for {
		<-ticker.C
		fmt.Println("Tick at", time.Now())
		if time.Now().After(closeTime) {
			fmt.Println("Max runtime reached, exiting loop.")
			break
		}
	}
}
