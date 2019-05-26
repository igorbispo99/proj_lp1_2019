package main

import "fmt"

func main() {

  // Declaracoes de array
  // Forma 1
  // var A[5] int

  // Forma 2
  B := [5]int{0,1,2,3,4}

  // Percorrer array
  for i, value := range B {
    fmt.Println(value, i)
  }

  // Slices - arranjos dinamicos e referencias a arrays

  // Forma 1
  
  Slice := []int{0,1,2,3,4}

  sub := Slice[2:4]
  fmt.Println(sub)

  // Forma 2(tipo, tamanho, capacidade)
  sub2 := make([]int, 5, 10)

  copy(sub2, Slice)

  fmt.Println(sub2)

  sub3 := append(Slice,2,-3,0)
  fmt.Println(sub3)
}
