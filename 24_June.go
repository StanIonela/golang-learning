package main

import (
	"fmt"
	"time"
)

/*func worker(c chan<- string) {
	time.Sleep(2 * time.Second)
	c <- "response"
}
*/

func stage1(c1 chan<- string) {
	select {
	case <-time.After(2 * time.Second):
		c1 <- "stage1: done\n"
	case <-time.After(4 * time.Second):
		c1 <- "stage1: fallback\n"
	}
}

func stage2(c2 <-chan string, c3 chan<- string) {
	select {
	case msg := <-c2:
		c3 <- msg + "stage2: done\n"
	case <-time.After(3 * time.Second):
		c3 <- "stage2: fallback\n"
	}
}
func main() {
	/*	c1 := make(chan string)
		c2 := make(chan string)

		go func() {
			time.Sleep(1 * time.Second)
			c1 <- "one"
		}()

		go func() {
			time.Sleep(2 * time.Second)
			c2 <- "two"
		}()

		for range 2 {
			select {
			case msg1 := <-c1:
				fmt.Println("received", msg1)
			case msg2 := <-c2:
				fmt.Println("received", msg2)
			case <-time.After(1 * time.Second):
				fmt.Println("timeout 2")
			default:
				fmt.Println("Something")
			}
		}*/

	/*	c1 := make(chan string)

		go worker(c1)

		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case <-time.After(1 * time.Second):
			fmt.Println("timeout 1")
		}*/

	ch1 := make(chan string)
	ch2 := make(chan string)

	go stage1(ch1)
	go stage2(ch1, ch2)

	fmt.Println("Final:\n", <-ch2)
}
