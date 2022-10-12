package main

import "fmt"

func main() {
	teste := 45
	var point *int = &teste // pointer to int
	y := *point             // y é 45 agora
	ptr := new(int)
	*ptr = 3

	fmt.Printf("teste", point, y, ptr)
}
