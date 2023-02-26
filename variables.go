package main

import "fmt"

func main() {
	a := "String"
	b := 122
	c := 3.14159
	d := true
	fmt.Printf("%T \n", a)
	fmt.Printf("a := \"String\" \t %T [%v]\n", a, a)
	fmt.Printf("b := 0 \t %T [%v]\n", b, b)
	fmt.Printf("c := 3.14158 \t %T [%v]\n", c, c)
	fmt.Printf("d := true \t %T [%v]\n", d, d)

	var (
		firstName string
		lastName  string
	)
	firstName = "Chi"
	lastName = "Kien"
	fmt.Printf(firstName + " " + lastName)

	bb := int64(b) //conversion
	const constant = 10
	fmt.Printf("Cast to int64 %v\n", bb)
	fmt.Println(constant)
}
