package main

// em vez de criar uma copia da variável para retorna o valor alterado
// esta recebendo um ponteiro para alterar o valor da variavel sem precisar retorna nada

func invert(num int) int {
	return num * -1
}

func invertPointer(num *int) {
	*num = *num * -1
}

func main() {
	numero1 := 10
	numInvertido := invert(numero1)
	invertPointer(&numero1)
	println(numInvertido)
	println(numero1)

}
