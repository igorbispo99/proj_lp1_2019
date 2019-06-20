package Files

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	//"fmt"
)

// Lê os dados em float
func ReadFile(fileName string) [][]float64 {
	var valOut float64
	file, err := os.Open(fileName) // Seleciona arquivo csv
	if err != nil {
		log.Fatal(err)
	} // Confere se arquivo existe
	defer file.Close()

	matrix, _ := csv.NewReader(file).ReadAll()

	inputs := make([][]float64, len(matrix))
	for i := range inputs {
		inputs[i] = make([]float64, len(matrix[0]))
	}

	for i, line := range matrix {
		for j, val := range line {
			if val == "" {
				valOut = 0
			} else {
				valOut, err = strconv.ParseFloat(val, 64)
				if err != nil {
					log.Fatal(err)
				}
			}

			inputs[i][j] = valOut
		}
	}
	return inputs
}

// Salva os arquivos CSV com as labels
