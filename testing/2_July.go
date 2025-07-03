// TODO [Implement]
// - Create Stack struct with methods: Push, Pop, Peek, IsEmpty
// - Create Queue struct with methods: Enqueue, Dequeue, Peek, IsEmpty
// - Handle underflow (empty pop/dequeue)

// TODO [Test]
// - Write unit tests for stack operations
// - Write unit tests for queue operations
// - Add edge case tests (empty pop/dequeue, peek)

// TODO [Optimize]
// - Learn why slicing from the front is O(n)
// - Rewrite queue using container/list for constant-time dequeue
// - Benchmark both versions

// TODO [Review]
// - Write markdown summary with time/space complexity
// - Push code and notes to GitHub repo
package main

import (
	"fmt"
)

type Stack []int

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func NewQueue() *Queue {
	return &Queue{}
}

func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	n := len(*s)
	val := (*s)[n-1]
	*s = (*s)[:n-1]
	return val, true
}

func (s Stack) Peek() (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	return s[len(s)-1], true
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

type Queue []int

func (q *Queue) Enqueue(v int) {
	*q = append(*q, v)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(*q) == 0 {
		return 0, false
	}
	val := (*q)[0]
	*q = (*q)[1:]
	return val, true
}

func (q Queue) Front() int {
	return q[0]
}

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}

func main() {
	var s Stack
	fmt.Println(s.IsEmpty())
	s.Push(10)
	s.Push(20)
	fmt.Println(s.Peek())
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.IsEmpty())

	var q Queue
	fmt.Println(q.IsEmpty())
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Front())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.IsEmpty())
}
