package main

import "fmt"

func sliceit() {
	var slice1 = make([]int, 5)
	fmt.Println(slice1)
	for i := 0; i < 5; i++ {
		slice1[i] = i + 9
	}
	fmt.Println(slice1[:3])
	fmt.Println(slice1[3:])
	fmt.Println(len(slice1))
	c := make([]int, 3)
	fmt.Println(copy(c, slice1))
}
