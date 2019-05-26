package main

import "fmt"

func main() {
	x,y := 5,3
	
	fmt.Println("x + y = ", x+y)
	fmt.Println("x - y = ", x-y)
	fmt.Println("x * y = ", x*y)
	fmt.Println("x / y = ", x/y)
	fmt.Println("x mod y = ", x%y)

	varbool := true
	otherbool := false


  // O printf

	// Equivale ao printf do C
	
	fmt.Printf("Printf: \n")
  fmt.Printf("x + y = %d \n", x+y)

	// Equivalente da declaração e definição anterior
	//		var varbool bool = true
	//		var otherbool bool = false
	

  	/*
  	  Aliás, os comentarios são iguais ao C.
  	*/

	fmt.Println(varbool && otherbool)
	fmt.Println(varbool || otherbool)
	fmt.Println(!(varbool) == otherbool)
}
