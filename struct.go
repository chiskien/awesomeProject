package main

import "fmt"

type example struct {
	pi      float32
	counter int16
	flag    bool
}

func main() {
	var ex1 example
	ex1.flag = true
	ex1.counter = 200
	ex1.pi = 3.14159
	fmt.Println(ex1)
}

