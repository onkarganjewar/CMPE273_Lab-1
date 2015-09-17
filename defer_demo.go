package main

import "fmt"


func start () {
	defer fmt.Println("Last statement in start")
	fmt.Println("Starting the demo")
}	


func close () {
fmt.Println("This function will be called last")
}


func main() {
	fmt.Println("This is the demo of defer function")
	defer close()
	start()
	fmt.Println("Inside main function")
}