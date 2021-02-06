package main

import "fmt"

func main() {
	fmt.Println("Variables")

	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Vijay"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(1, 2)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)
}
