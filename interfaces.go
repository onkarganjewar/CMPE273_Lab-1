package main

import (
	"fmt"
	"math"
	)

type Shape interface {
perimeter() float64
}

type Circle struct {
r float64
}


type Rectangle struct {
l,w float64
}

func (r Rectangle) perimeter() float64 {
return 2*(r.l+r.w)
}

func (c Circle) perimeter() float64 {
return 2*math.Pi*c.r
}

func calc (s Shape) float64 {
return (s.perimeter());
}

func mea (ss Shape) float64 {
	return (ss.perimeter());
}

func main() {
var rad, le, wd float64
fmt.Println("Enter the radius for a Circle")
fmt.Scanf("%f",&rad);
fmt.Println("Enter the length for a Rectangle")
fmt.Scanf("%f",&le);
fmt.Println("Enter the width for a Rectangle")
fmt.Scanf("%f",&wd);

cc := Circle{rad}
rr := Rectangle{le,wd}

fmt.Println(calc(cc))
fmt.Println(mea(rr))

}
