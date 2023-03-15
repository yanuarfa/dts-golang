package main

import "fmt"

func main() {
	const i int = 21
	const j bool = true
	ya := 'Я'
	var k float64 = 123.456

	fmt.Printf("%d \n", i)
	fmt.Printf("%T \n", i)
	fmt.Printf("%% \n")
	fmt.Printf("%t \n", j)
	fmt.Printf("\n")
	fmt.Printf("%c \n", ya)
	fmt.Printf("%d \n", 21)
	fmt.Printf("%o \n", 21)
	fmt.Printf("%x \n", 15)
	fmt.Printf("%X \n", 3859)
	fmt.Printf("%U \n", ya)
	fmt.Printf("\n")
	fmt.Printf("%f \n", k)
	fmt.Printf("%e \n", k)
	fmt.Scanln()
}
