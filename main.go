package main

import "fmt"

func main() {
var Age int
var count int

count = 0

fmt.Println("Hello! Please enter any number.")

fmt.Scanln(&Age)

for Age != count{
  fmt.Println("Are we there yet!?")
  count ++
}
}