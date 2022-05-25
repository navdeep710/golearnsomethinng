package main

import "fmt"

func mapmygenes() {
	var mmap = make(map[int]string)
	mmap[1] = "one"
	fmt.Println(mmap)
	delete(mmap, 1)
	_, is_present := mmap[1]
	fmt.Printf("is the key 1 present in the map %t", is_present)
}
