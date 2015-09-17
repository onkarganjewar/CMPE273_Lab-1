package main

import "fmt"
import i "golang-book/shapes/inter"


func main() {

var rad, le, wd float64
fmt.Println("Enter the radius for a Circle")
fmt.Scanf("%f",&rad);
fmt.Println("Enter the length for a Rectangle")
fmt.Scanf("%f",&le);
fmt.Println("Enter the width for a Rectangle")
fmt.Scanf("%f",&wd);

cc := i.Circle{rad}
rr := i.Rectangle{le,wd}

fmt.Println(i.Calc(&cc))
fmt.Println(i.Mea(&rr))

}
