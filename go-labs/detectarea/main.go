package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	fmt.Println("Detect area for rectangle and triangle")
	tr := triangle{base: 4, height: 4}
	sq := square{sideLength: 4}
	printArea(tr)
	printArea(sq)
}

func (tr triangle) getArea() float64 {
	return 0.5 * (tr.base) * (tr.height)
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}
