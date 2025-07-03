package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()

	if _, ok := s.Pop(); ok {
		t.Errorf("Expected false on Pop with empty stack")
	}

	s.Push(10)
	s.Push(20)

	if val, _ := s.Peek(); val != 20 {
		t.Errorf("Expected 20 on Peek, got %d", val)
	}

	if val, _ := s.Pop(); val != 20 {
		t.Errorf("Expected 20 on Pop, got %d", val)
	}

	if val, _ := s.Pop(); val != 10 {
		t.Errorf("Expected 10 on Pop, got %d", val)
	}
}

func TestQueue(t *testing.T) {
	q := NewQueue()

	if _, ok := q.Dequeue(); ok {
		t.Errorf("Expected false on Dequeue from empty queue")
	}

	q.Enqueue(1)
	q.Enqueue(2)

	if val, _ := q.Dequeue(); val != 1 {
		t.Errorf("Expected 1 on Dequeue, got %d", val)
	}

	if val, _ := q.Dequeue(); val != 2 {
		t.Errorf("Expected 2 on Dequeue, got %d", val)
	}
}
