package main

import "fmt"

func MakeSwitch() {
	whatami := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i am boolean")
		case int:
			fmt.Println("i am integer")
		default:
			fmt.Printf("dont know the type %T", t)
		}

	}
	whatami(1)
	whatami(true)
	whatami("hey")
}
