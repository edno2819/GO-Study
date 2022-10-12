package main

import "fmt"

var meu_bol bool = "casa" != "sdfsddfs"
var nume float32 = 100.45
var sd = `Teste de 
			Formataao 
				Muito 
					LOUCA`

func main() {
	tes := 5 < 4
	tes2 := (5 != 4)

	// var tes = 5 < 4

	fmt.Printf("%t %t\n", tes, tes2)
	fmt.Println(meu_bol)
	fmt.Println(nume)
	fmt.Println(sd)

}
