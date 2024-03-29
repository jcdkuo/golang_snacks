package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

type Rect2 struct {
	width  float64
	height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (rr Rect2) Area() float64 {
	return rr.width
}

func (rr Rect2) Perimeter() float64 {
	return 2
}

func mm(ss Shape) {
	fmt.Println("Area: ", ss.Area())
	fmt.Println("Peri: ", ss.Perimeter())
}

func main() {
	var s Shape
	s = Rect{5.0, 4.0}
	r := Rect{5.0, 4.0}
	rr := Rect2{7.0, 8.0}

	fmt.Printf("type of s is %T\n", s)
	fmt.Printf("value of s is %v\n", s)
	fmt.Println("area of rectange s", s.Area())
	fmt.Println("s == r is", s == r)

	mm(s)
	mm(r)
	mm(rr)
}
