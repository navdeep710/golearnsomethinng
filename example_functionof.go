package main

import "strings"

func summationoftwo(a int, b int) int {
	return a + b
}

func multiplicationoftwo(a int, b int) int {
	return a * b
}

func functionoffunction(myfunc func(int, int) int, a int, b int) int {
	response := myfunc(a, b)
	return response
}

func functionreturnfirstandrest(mstr string, splitter string) (string, string) {
	splitted := strings.Split(mstr, splitter)
	return splitted[0], strings.Join(splitted[1:], splitter)

}
