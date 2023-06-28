package main

import "fmt"

type fenix int
var x fenix

func main() {
	fmt.Printf("%v\t%T\n", x, x)

	x = 42
	fmt.Println(x)
}