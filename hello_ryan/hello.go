package main

import "fmt"

func main(){
  fmt.Println("What's your name?")

  var name string

  fmt.Scanln(&name)

  greeting := fmt.Sprintf("Hello %v!", name)

  fmt.Println(greeting)
}
