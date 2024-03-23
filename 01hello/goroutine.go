package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Second) // Simulate some work
	}
}

func main() {
	// Start a new goroutine to execute printNumbers concurrently
	go printNumbers()

	// Start another goroutine to print numbers from 5 to 9 concurrently
	go func() {
		for i := 5; i < 10; i++ {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}()

	// Wait for a keystroke before exiting to see the goroutine output
	var input string
	fmt.Scanln(&input)
}
