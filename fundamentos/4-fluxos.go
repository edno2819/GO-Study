package main

import (
	"fmt"
	"time"
)

func loop_for() {
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}

	g := "string de teste"
	for x := 0; x < len(g); x++ {
		fmt.Printf("Letra: %c   |  Posição: %v \n", g[x], x)
	}

	//for without a condition will loop repeatedly until you
	//break out of the loop or return from the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	for pos, char := range "outra string" {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}

	//You can also continue to the next iteration of the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func conditions_if_else() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	//You can have an if statement without an else.

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
}

func conditions_swtch() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// You can use commas to separate multiple expressions in
	// the same case statement. We use the optional
	// default case in this example as well.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// switch without an expression is an alternate way to express
	// if/else logic. Here we also show how the case expressions
	// can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// A type switch compares types instead of values. You can use this to
	// discover the type of an interface value. In this example, the variable
	// t will have the type corresponding to its clause.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

func main() {
	// conditions_if_else()
	conditions_swtch()
	// loop_for()

}
