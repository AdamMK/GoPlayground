package main

import "fmt"

type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape)  {
	fmt.Println("The area of the figure is ", s.getArea())
}

func main()  {
	sq := square{
		sideLength: 10,
	}
	tr := triangle{
		height: 10,
		base: 10,
	}

	printArea(sq)
	printArea(tr)
}
