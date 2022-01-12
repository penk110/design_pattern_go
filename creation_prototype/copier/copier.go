package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/copier"
)

type User struct {
	Name string
	Age  int
}

func main() {
	var (
		u1  *User
		u2  *User
		err error
	)

	u1 = &User{
		Name: "u1",
		Age:  18,
	}
	u2 = &User{}

	fmt.Printf("u1 ptr: %p\n", u1)

	// target source
	if err = copier.Copy(u2, u1); err != nil {
		log.Fatalln(err)
		return
	}
	// u1 ptr: 0xc00011c018
	// u2 ptr: 0xc00011c030 &main.User{Name:"u1", Age:18}
	fmt.Printf("u2 ptr: %p %#v\n", u2, u2)
}
