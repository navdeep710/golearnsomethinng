package main

import "fmt"

type student struct {
	name string
	age  int
}

type class struct {
	students []student
	name     string
}

func makeclass() {
	sixthClass := class{
		name:     "5 class",
		students: []student{{"hamar", 3}, {"nam", 4}},
	}
	fmt.Println(sixthClass)
}
