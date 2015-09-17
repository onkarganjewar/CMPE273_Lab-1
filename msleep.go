package main

import(
  "fmt"
  "time"
)

func my_sleep(n int){
  select {
  case <-time.After(time.Duration(n)*time.Second):
    fmt.Println("End of custom sleep function")
  }
}

func main(){
  fmt.Println("Starting the program")
  my_sleep(3)
  fmt.Println("Ending the program")

  var sample string
  fmt.Scanf("%s",&sample)
}