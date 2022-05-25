package main

import "fmt"

func Arrayize() {
	myarr := [5]int{1, 2, 3, 4, 5}
	var myarr2 [5]int = [5]int{1, 2, 3, 4}
	var myarr3 [3]string
	myarr3[2] = "buleya"
	for i := 0; i < 5; i++ {
		myarr[i] = i + 5
	}
	fmt.Println(myarr)
	fmt.Println(myarr2)
	fmt.Println(myarr3)

}
