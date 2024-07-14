package main

import "fmt"

func main(){
  fmt.Println("What is your name?")

  var name string

  fmt.Scanln(&name)

  var pName *string = &name

  greeting := fmt.Sprintf("Hello %v! Wait a second, what?", pName)

  fmt.Println(greeting)

  fixedGreeting := fmt.Sprintf("Is it...hello %v?", *pName)

  fmt.Println(fixedGreeting)
}
