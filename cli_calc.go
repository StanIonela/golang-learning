package main 

import (
  "fmt"
)

func main() {
    var a, b int
    fmt.Print("Enter two numbers: ")
    fmt.Scan(&a, &b)
    fmt.Printf("Sum = %d\n", a+b)
  }
