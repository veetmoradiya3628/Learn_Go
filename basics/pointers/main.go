package main

import "fmt"

type User struct {
	name string
	age  int
}

func (u *User) updateName(newName string) {
	u.name = newName
}

func main() {
	userId := 42
	fmt.Println(&userId)

	user := User{name: "Alice", age: 30}
	fmt.Println("Before update:", user.name)
	user.updateName("Bob")
	fmt.Println("After update:", user.name)

}
