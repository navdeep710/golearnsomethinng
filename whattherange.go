package main

import "fmt"

func checkitsrange() {
	marr := []int{1, 2, 3, 4}
	for index, value := range marr {
		fmt.Println(index, value)
	}

	mmap := make(map[int]string)
	mmap[1] = "1"
	mmap[2] = "2"
	mmap[3] = "3"

	for key, value := range mmap {
		fmt.Println(key, value)
	}
}
