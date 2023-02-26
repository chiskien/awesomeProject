package main

import "fmt"

func main() {
	var array [10]int
	array = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 19}
	var slices []int
	fmt.Println(slices, len(slices), cap(slices))
	slices = append(slices, 10)
	fmt.Println(slices, len(slices), cap(slices))
	fmt.Println(array)

}
