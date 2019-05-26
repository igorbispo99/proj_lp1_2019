package main

import "fmt"

func main() {
  x := 10
  var y *int = &x

  fmt.Println(x)
  fmt.Println(*y)
}
