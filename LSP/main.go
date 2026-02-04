package main

import "fmt"

// LSP - Liskov substitution principle

type Rectangle struct {
	width, height float64
}

type Rectanglar interface {
	setWidth(float64)
	setHeight(float64)
	GetArea() float64
}

func (r Rectangle) setWidth(w float64) {
	r.width = w
}

func (r Rectangle) setHeight(h float64) {
	r.height = h
}

func (r Rectangle) GetArea() float64 {
	return r.width * r.height
}

type Square struct {
	side float64
}

func (s Square) setWidth(w float64) {
	s.side = w
}

func (s Square) setHeight(h float64) {
	s.side = h
}
func (s Square) GetArea() float64 {
	return s.side * s.side
}

func main() {
	s1 := Square{}
	s1.setWidth(10)
	s1.setHeight(5)

	area1 := s1.GetArea()
	if area1 == 50 {
		fmt.Println("True logic with Rectangle")
	} else {
		fmt.Println("True logic with Square")
	}

}
