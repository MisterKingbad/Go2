package main

import "fmt"

type fenix int

var x fenix

var y fenix

func main() {

	fmt.Printf("%v\t%T\n", x, x)

	x = 42
	fmt.Println(x)

	x =fenix(x)

	y = x
	fmt.Println(y)
	fmt.Printf("%T", y)
}