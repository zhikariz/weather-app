package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time
	currentTime := time.Now()

	// Format the time according to the desired layout
	formattedTime := currentTime.Format("2006-01-02T15:04:05-07:00")

	fmt.Println("Formatted Time:", formattedTime)
}
