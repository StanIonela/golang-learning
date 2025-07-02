

// TODO [Implement]
// - Create a Node struct with value and next pointer
// - Write a function to append values to the list
// - Add a method to find a node by value
// - Add a method to delete a node by value

// TODO [Test]
// - Write basic unit tests for append, find, delete
// - Test edge cases: empty list, head deletion, tail deletion

// TODO [Analyze]
// - Document time complexity for all operations
// - Write short notes comparing to slices

// TODO [Extra]
// - Implement insert at head (O(1))
// - Reverse the linked list (optional)
// - Add Length() method

// TODO [Wrap-Up]
// - Clean and refactor the code
// - Push to GitHub
// - Write summary in `linked_list.md`

package main  
import (
  "fmt"
)

type Node struct {
  value int 
  next *Node
}

func append_value(head *Node, val []int) *Node {
  var tail *Node
  for _, j := range val {
    n := &Node{value: j}
    if head == nil {
      head = n
      tail = n
    } else {
      tail.next = n 
      tail = n  
    }
  }
  return head
}

func find_value(head *Node, target int) *Node {
  for curr := head; curr != nil; curr = curr.next {
    if curr.value == target{
      return curr
    }
  }
  return nil
}

func delete_node (head *Node, value int) *Node {
  if head == nil {
    return nil
  }

  if head.value == value {
    return head.next
  }

  prev := head
  curr := head.next

  for curr != nil {
    if curr.value == value {
      prev.next = curr.next
      break
    }
    prev = curr
    curr = curr.next
  }
  return head
}

func printList(head *Node) {
  for curr := head; curr != nil; curr = curr.next {
    fmt.Printf("%d", curr.value)
    if curr.next != nil {
      fmt.Printf(" -> ")
    }
  }
  fmt.Println()
}

func main()  {
  values := []int{5, 10, 15}
  head := append_value(nil, values)
  printList(head)
  f := find_value(head, 10)
  fmt.Println(f.value)
  d := delete_node(head, 10)
  printList(d)
  
}
