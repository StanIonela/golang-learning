package main

import (
	"fmt"
	"time"
)

/*func m(message string, delay time.Duration) {
* time.Sleep(delay)
	fmt.Println(message)
}*/

func sleepAndPrint(message string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println(message)
}
func main() {
	go sleepAndPrint("Something", 1*time.Second)
	go sleepAndPrint("Feels", 2*time.Second)
	go sleepAndPrint("Different", 3*time.Second)
	go sleepAndPrint("I", 4*time.Second)
	go sleepAndPrint("Think", 5*time.Second)

	time.Sleep(6 * time.Second)

	fmt.Println("Done")

	// TODO [Mini Project]
	// - Create a function: sleepAndPrint(message, delay)
	// - Launch 5 goroutines with different delays
	// - Expected result: each prints after its delay

	// TODO [Organize Code]
	// - Split code into main.go and worker.go (optional)
	// - Add comments to explain goroutine behavior

	// TODO [Git]
	// - Stage and commit your work
	// - Commit message: "add goroutine examples and mini project"
}
