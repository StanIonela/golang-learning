package main
import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "strings"
)

func safe_divide (a, b int) (result int) {
    defer func() {
     if r := recover(); r != nil {
      fmt.Println("Recovered after panic: \n", r)
        result = 0 
    }
  }()
  if b == 0 {
    panic("Divide by 0 don't work")
  }
  return a / b
}

func main() {
  reader := bufio.NewReader(os.Stdin)

  for {
    fmt.Println("Introduce two number for division or 'q' to quit: ")
    line, _ := reader.ReadString('\n')
    line = strings.TrimSpace(line)

    if line == "q" {
      break
    }

    parts := strings.Fields(line)
    if len(parts) != 2 {
      fmt.Println("Please enter two numbers")
      continue
    }
  

  a, err1 := strconv.Atoi(parts[0])
  b, err2 := strconv.Atoi(parts[1])
  if err1 != nil || err2 != nil {
    fmt.Println("Invalid input. Enter valid integers.")
    continue
  }

  result := safe_divide(a, b)
  fmt.Printf("Result: %d / %d = %d\n", a, b, result)
}
}

