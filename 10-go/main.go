package main

import "fmt"

func main() {

	obj := User{"Name", "name@go.dev", true, 16}
	fmt.Println(obj)
	fmt.Printf("details are:%+v\n", obj)
	fmt.Printf("details are:%v\n", obj.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
