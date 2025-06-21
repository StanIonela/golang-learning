package main
import "fmt"

type User struct {
    name string 
    age int
    email string
  }

func new_user(name string) *User{
  u := User{name: name}
  u.email = "example@gmail.com"
  u.age = 30
  return &u
}

func (u User) full_name() string {
  return u.name
}

func (u *User) is_adult() bool {
  if u.age >= 18 {
    return true
  }
  return false
}

var (
  name1 = User{"Bob", 25, "example1@gmail.com"}
  name2 = User{name: "Jhon", age: 15}
  name3 = User{name: "Sussy"}
  name4 = &User{"Karen", 30, "example2@gmail.vom"}
  name5 = new_user("Batman")
)

func main() {
  name1.name = "Kiril"
  fmt.Println(name1, name2, name3, name4, name5)
  fmt.Println("Name: ", name1.full_name())
  fmt.Println("Is Adult? ", name2.is_adult())
}
