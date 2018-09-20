package main

import "fmt"

func main() {
	var user1 user
	user1.name = "asfa"
	i := 56
	user1.update(i)
	fmt.Println(user1.name)
}

type user struct {
	name string
}
func (u user)update(i int)user{
	if i == 3{
		u.name = "wwww"
	}else{
		u.name = "gggg"
	}
	return u
}