package main

import "fmt"

func main() {
	const i int = 21
	const j bool = true
	ya := "\u042f"
	var k float64 = 123.456

	fmt.Printf("%d \n", i)
	fmt.Printf("%T \n", i)
	fmt.Printf("%% \n")
	fmt.Printf("%t \n", j)
	fmt.Printf("\n")
	fmt.Printf("%q \n", ya)
	fmt.Printf("%d \n", i)
	fmt.Printf("%o \n", 25)
	fmt.Printf("%x \n", "f")
	fmt.Printf("%X \n", "F 13")
	fmt.Printf("%+q \n", ya)
	fmt.Printf("\n")
	fmt.Printf("%f \n", k)
	fmt.Printf("%e \n", k)
}
