package main

import "fmt"

const tempF = 373.2

func main() {

	var F = tempF
	var C int = int(F - 273)
	fmt.Printf("A temperatura de ebulição em C é: %d e a temperatura de ebulição em F é: %g", C, F)
}
