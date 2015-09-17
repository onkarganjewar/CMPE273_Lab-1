package inter

import (
//	"fmt"
	"math"
	)

type Shape interface {
Perimeter() float64
}

type Circle struct {
Radius float64
}

type Rectangle struct {
Lenn,wdr float64
}

func (r Rectangle) Perimeter() float64 {
return 2*(r.Lenn+r.wdr)
}

func (c Circle) Perimeter() float64 {
return 2*math.Pi*c.Radius
}

func Calc (s Shape) float64 {
return (s.Perimeter());
}

func Mea (ss Shape) float64 {
	return (ss.Perimeter());
}
