package main

import "fmt"
import f "golang-book/packages_demo/fibo"

func main() {
var num int
fmt.Printf("\nEnter the number to generate fibonacci series \n");
fmt.Scanf("%d",&num);
result := f.Fib(num);
fmt.Println("result is ",result);

}