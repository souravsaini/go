package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	s := person{name: "Sourav", age: 20}
	fmt.Println(s.name, s.age)
}
