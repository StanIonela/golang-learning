package main
import (
  "fmt"
  "bufio"
  "os"
  "sync"
//  "time"
)

/*func worker (name string, ch chan <- string) {
  for i := 1; i <= 3; i++ {
    ch <- fmt.Sprintf("%s: %d", name, i)
    time.Sleep(500 * time.Millisecond)
  }
}

func fan_in(ch1, ch2 <- chan string) <- chan string {
  out := make(chan string)
  go func() {
    for v := range ch1 {
      out <- v
    } 
  }()

  go func() {
    for v := range ch2 {
      out <- v
    }
  }()

  return out
}
*/

func count_lines (filename string, ch chan <- string, wg *sync.WaitGroup) {
  defer wg.Done()

  file, err := os.Open(filename)
  if err != nil {
    fmt.Println("Couldnt open file: ", err)
    return
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)
  lines := 0
  for scanner.Scan() {
    lines++
  }

  ch <- fmt.Sprintf("%s: %d lines", filename, lines)
}

func main() {
/*  msg := make (chan string)

  go func() {msg <- "ping"}()
  message := <- msg
  fmt.Println(message)

  go func() {msg <- "something"}()
  message1 := <- msg
  fmt.Println(message1)

  msg1 := make(chan string, 3)
  msg1 <- "buffered"
  msg1 <- "unbuffered"
  msg1 <- "channels"

  fmt.Println(<-msg1)
  fmt.Println(<-msg1)
  fmt.Println(<-msg1)

  ch := make(chan int)
  go func() {
    for i := 1; i <= 5; i++ {
       ch <- i
    }
    close(ch)
  }()

  for v := range ch {
    fmt.Println("Got: ", v)
  }*/

 /* a := make(chan string)
  b := make(chan string)

  go worker("A", a)
  go worker("B", b)

  merge := fan_in(a, b)

  go func() {
    time.Sleep(2 * time.Second)
    close(a)
    close(b)
  }()

  for i := 0; i < 6; i++ {
    fmt.Println(<-merge)
  } */

  files := []string{"test1.txt", "test2.txt", "test3.txt"}

  ch := make(chan string)
  var wg sync.WaitGroup

  for _, f := range files{
    wg.Add(1)
    go count_lines(f, ch, &wg)
  }

  go func() {
    wg.Wait()
    close(ch)
  }()

  for msg := range ch {
    fmt.Println(msg)
  }
}

