package main

import "fmt"

func main() {
	var user1 user
	user1.name = "www"
	name := user.notify(user1).name
	fmt.Println(name)
}

type user struct {
	name  string
}

// notify implements a method with a value receiver.
func (u user) notify()(user) {
	u.name = "1212"
	return u
}

// changeEmail implements a method with a pointer receiver.
func (u *user) changeEmail()(*user) {
	u.name = "1212"
	return u
}