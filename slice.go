package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(len(s), s[1])
	s = append(s, "1", "2")
	fmt.Println(s)
	fmt.Println(s[:])
}
