package main

import (
    "fmt"
    "math"
)

type geometricShapes interface {
	area() float64
	perimeter() float64
}

type circle struct {
    radius float64
}
type rect struct {
    width, height float64
}
/* Initial draft
func area(w int, l int) float64 {
	return w * l
}

func perimeter(l int, w int) float64 {
	return (2*l) + (2*w)
}
*/

// Solution from Go by Example https://gobyexample.com/interfaces

func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometricShapes) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
		c := circle{radius: 5}
		
		measure(r)
    measure(c)
}