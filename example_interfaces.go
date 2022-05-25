package main

import "fmt"

type geometry interface {
	area() int
	perimeter() int
}

type rectangle struct {
	length  int
	breadth int
}

type square struct {
	side int
}

func (r rectangle) area() int {
	return r.length * r.breadth
}

func (s square) area() int {
	return s.side * s.side
}

func (r rectangle) perimeter() int {
	return 2*r.length + r.breadth
}

func (s square) perimeter() int {
	return 4 * s.side
}

func getarea(geom geometry) int {
	area1 := geom.area()
	fmt.Printf("area of the first geometry is %d\n", area1)
	perimeter := geom.perimeter()
	fmt.Printf("area of the first geometry is %d\n", perimeter)
	return area1
}

func getareaofanygeometry() {
	r := rectangle{2, 2}
	s := square{2}
	getarea(r)
	getarea(s)
}
