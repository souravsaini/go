package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["key1"] = 10
	m["key2"] = 20
	m["key3"] = 30

	fmt.Println(m)
	fmt.Println(m["key1"])
	fmt.Println(len(m))
	a, b := m["key3"] //NOTE
	fmt.Println(a, b)

}
