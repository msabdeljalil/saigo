package main

import (
	"fmt"
	"math"
)

////////////
// Square //
////////////

type Square struct {
	side float64
}

func (s *Square) Name() string {
	return "Square"
}

func (s *Square) Perimeter() float64 {
	return 4 * s.side
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

////////////
// Circle //
////////////

type Circle struct {
	radius float64
}

func (c *Circle) Name() string {
	return "Circle"
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Triangle (Isocleles)

type Triangle struct {
	LongSide  int
	ShortSide int
	Height    float64
}

func (t *Triangle) Name() string {
	return "Triangle"
}

func (t *Triangle) Perimeter() float64 {
	return float64(t.ShortSide + t.LongSide*2)
}

func (t *Triangle) Area() float64 {
	t.Height = math.Sqrt(math.Pow(float64(t.LongSide), 2.0) - math.Pow(0.5*float64(t.ShortSide), 2.0))
	return float64(0.5) * float64(t.ShortSide) * float64(t.Height)
}

////////////////
// Efficiency //
////////////////

type Shape interface {
	Name() string
	Perimeter() float64
	Area() float64
}

func Efficiency(s Shape) {
	name := s.Name()
	area := s.Area()
	rope := s.Perimeter()

	efficiency := 100.0 * area / (rope * rope)
	fmt.Printf("Efficiency of a %s is %f\n", name, efficiency)
}

func main() {
	s := Square{side: 10.0}
	Efficiency(&s)

	c := Circle{radius: 10.0}
	Efficiency(&c)

	t := Triangle{LongSide: 4, ShortSide: 2}
	Efficiency(&t)
}
