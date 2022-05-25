package main

import "fmt"

func TryLoop() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("current value of loop is %d\n\n", i)
		}

	}
}
