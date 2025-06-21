package main

import "fmt"

type User struct {
	name, email string
	age         int
}

func user_data(name, email string, age int) *User {
	u := User{name: name, age: age, email: email}
	return &u
}

func (u *User) full_name2(name string) string {
	return u.name
}

func (u *User) can_vote(age int) bool {
	if u.age >= 18 {
		return true
	}
	return false
}

func (u *User) retair() string {
	if u.age >= 65 {
		return fmt.Sprintf("User is already retaired.")
	} else if u.age < 18 {
		return fmt.Sprintf("User is not working.")
	}
	ret := 65 - u.age
	return fmt.Sprintf("Years until retairment: %d", ret)
}

func (u *User) update_email(new_email string) *User {
	u.email = new_email
	return u
}

func main() {
	var name, email string
	var age int
	fmt.Println("Input User Data")
	fmt.Print("Name: ")
	fmt.Scan(&name)

	fmt.Print("Age: ")
	fmt.Scan(&age)

	fmt.Print("Email: ")
	fmt.Scan(&email)

	res := user_data(name, email, age)
	fmt.Printf("Name: %s,\nAge: %d,\nEmail: %s\n", res.name, res.age, res.email)

	fmt.Printf("Name: %s\n", res.full_name2(name))
	fmt.Println("User can vote?", res.can_vote(age))

	fmt.Println(res.retair())

	var option, update_email string
	fmt.Print("Do you eant to update email?(yes/no)")
	fmt.Scan(&option)
	switch option {
	case "yes":
		fmt.Print("Introduce new email: ")
		fmt.Scan(&update_email)
		res.update_email(update_email)
		fmt.Printf("Name: %s,\nAge: %d,\nEmail: %s\n", res.name, res.age, res.email)
	case "no":
		break
	}
}
