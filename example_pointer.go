package main

import "fmt"

func mypointer(value *int) {
	*value = 20
}

func structuralchange() {
	chika := 10
	fmt.Printf("value of chika is %d\n", chika)
	mypointer(&chika)
	fmt.Printf("value of chika is %d\n", chika)
}
