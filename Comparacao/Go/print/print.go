package main

import "fmt"

func main() {

	// No C isso não é possível
	var name string = "Some text"
	name += " more text"
	win := true
	const pi float64 = 3.141592

  // O printf

  // Equivale ao printf do C
  // Formatação um pouco mais flexivel
	fmt.Printf("%d \n", 3)

	fmt.Printf("%s \n", name)

	// C não exibe booleanos
	fmt.Printf("%t \n", win)

	// C não é capaz de printar binarios
	// com seus recursos primarios
	fmt.Printf("%b \n", 12)

	// Iguais ao C
	fmt.Printf("%c \n", 34)
	fmt.Printf("%X \n", 12)
	fmt.Printf("%e \n", pi)
}