package main

import "fmt"


func panicdemo () {
	defer fmt.Println("What is happening? ")
	fmt.Println("This is the demo of panic function");
	fmt.Println("Program will panic in 3...");
	fmt.Println("2..");

	/*defer func () {
		str := recover();
		if str != nil {
		fmt.Println("Recovering....", str)}
	}()*/
	//panic("Abort!!!!");
	
}	

func neww () {
defer fmt.Println("It will happen if we recover from panic")
}

func random () {
fmt.Println("This will be called last")
}


func main() {
	defer random()
	panicdemo()
	neww()
defer fmt.Println("What just happened");
}