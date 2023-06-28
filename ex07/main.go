package main

import "fmt"

func main() {
	a := (1 == 10) 
	b := (1 < 10)
	c := (1 > 10)
	d := (1 <= 10)
	e := (1 >= 10)
	f := (1 != 10)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}