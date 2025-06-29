// TODO [Mini Project]
// - Build a `Counter` struct that is safe for concurrent use
// - Add methods to increment and read the value
// - Launch multiple goroutines using this struct
// - Use race detector to ensure clean output

// TODO [Reflection]
// - Write down why race conditions are dangerous
// - Summarize when to use `sync.Mutex` or `sync.RWMutex`
// - Push your progress to Git with message: "Add safe counter with mutex"
package main

import (
	"fmt"
	"sync"
	//"time"
)

/*var (
	counter int
	mu      sync.Mutex
)

func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Counter: ", counter)
}*/

type Counters struct {
	mu    sync.Mutex
	value int
}

func (c *Counters) increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counters) read_value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := Counters{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.increment()
		}()
	}

	wg.Wait()
	fmt.Println("Count: ", counter.read_value())
}
