package main

import (
	"fmt"
	"math/rand"
)

type node struct {
	value int
	left  *node
	right *node
}

func newnode(value int) *node {
	return &node{value, nil, nil}
}

func addtonode(noderoot *node, value int) *node {
	if noderoot == nil {
		noderoot = newnode(value)
		return noderoot
	}
	if noderoot.value > value {
		if noderoot.left == nil {
			noderoot.left = newnode(value)
			return noderoot.left
		} else {
			return addtonode(noderoot.left, value)
		}
	} else {
		if noderoot.right == nil {
			noderoot.right = newnode(value)
			return noderoot.right
		} else {
			return addtonode(noderoot.right, value)
		}
	}
}

func printinorder(root *node) {
	if root == nil {
		return
	}
	printinorder(root.left)
	printinorder(root.right)
	fmt.Println(root.value)
}

func printpreorder(root *node) {
	if root == nil {
		return
	}
	fmt.Println(root.value)
	printinorder(root.left)
	printinorder(root.right)
}

func (root *node) preorder() {
	if root == nil {
		return
	}
	fmt.Println(root.value)
	(*node).preorder(root.left)
	(*node).preorder(root.right)

}

func addandprintasampletree() {

	root := addtonode(nil, 1)
	for i := 0; i < 100; i++ {
		addtonode(root, rand.Intn(100))
	}
	//fmt.Println("printing in order")
	//printinorder(root)
	//fmt.Println("printing pre order")
	//printpreorder(root)
	fmt.Println("printing via method")
	root.preorder()
}
