package main

import "fmt"

func ifelse() {
	nums := []int{1, 2, 3, 45, 5}
	for i, j := range nums {
		if i%3 == 0 {
			fmt.Printf("i am divisible by 3, because i am %d", j)
		} else if i%2 == 0 {
			fmt.Printf("i am divisible by 2, because i am %d", j)
		} else {
			fmt.Printf("i am %d", j)
		}
	}

}
