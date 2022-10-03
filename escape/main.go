package main

import (
	"fmt"
)

type User struct {
	ID    int
	Login string
}

func (u *User) GetID() int {
	return u.ID
}

func newUser(login string) *User {
	return &User{123, login}
}

var g *int

func setToZero(in *int) {
	// g = in
	*in = 0
}

func main() {

	u := newUser("test1")
	u.ID = 1

	data := make([]string, 20)
	data = append(data, "test2")

	i := 1
	setToZero(&i)

	// _ = fmt.Sprint(data)
	// _ = fmt.Sprint(u)
	fmt.Println("test3")

	foo(u)

	foo2(u)

}
