package main

import "fmt"

var x int = 42
var nome string = "James Bond"

func main() {
	casa := "teto de vidro"
	fmt.Printf("Alá! garalera do %s", casa)
	teste(10)
}

func teste(dedo int) {
	d := fmt.Sprintf("%s %d gordo", nome, dedo)
	fmt.Printf("\nOLá %v", d)

}
