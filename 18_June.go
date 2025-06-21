package main
import "fmt"

//func test() {
//    var p *int
//    fmt.Println("pointer: ", &p)

//    i := 2
//    p2 := &i
//    fmt.Println("pointer2: ", *p2)
//    *p2 = 3
//    fmt.Println("pointer3: ", i)
//  }

//type Counter struct {
//  score int
//}

//func user_score(score int) *Counter{
//  s := Counter{score: score}
//  return &s;
//}

//func (c *Counter) add(nr int) {
//  c.score += nr
//}

//func (c *Counter) remove(nr int) {
//  c.score -= nr
//}

func SafeDivide(a, b int) (result int) {
  defer func() {
    if r := recover(); r != nil {
      fmt.Println("Recovered from panic: \n", r)
      result = 0
    }
  }()

  if b == 0 {
    panic("Division by zero can't be")
  }

  return a/b
}

func main() {
//  var opt string
//  var nr int
//  res := user_score(0)
//  for {
//    fmt.Printf("Your score: %d\n", res.score)
//    fmt.Println("You want to add or remove point from user_score or you finised? (+/-/done)")
//    fmt.Scan(&opt)
//    switch opt {
//      case "+":
//        fmt.Println("How much point to add?")
//        fmt.Scan(&nr)
//       res.add(nr)
//      case "-":
//        fmt.Println("How much point to remove?")
//        fmt.Scan(&nr)
//        res.remove(nr)
//      case "done":
//        fmt.Printf("Your final score: %d\n", res.score)
//        return
//      default:
//        fmt.Println("Invalid option. Use +, - or done")
//    }
//  }
//  defer func() {
//   if r := recover(); r != nil {
//      fmt.Println("Recovered.Error: \n", r)
//    }
//  }()
//  panic("A problem!")
//  fmt.Println("After the panic")
  a := 25
  b := 25
  fmt.Println(SafeDivide(a, b))
}
