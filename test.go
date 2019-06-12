package main


import(
		"fmt"
		"./Files"
		//"./Forest" 
)

func main(){
	//inputs := make([][]interface{}, 0)
	inputs := Files.ReadFile("./data/test.csv")
	//fmt.Printf("%d", len(inputs))
 	for _,line := range inputs{
		for _,val := range line{
			fmt.Printf("%f ",val)
		}
		fmt.Println()
	} 
}