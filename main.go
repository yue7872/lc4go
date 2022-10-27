package main

import (
	"fmt"

	"lc4go/lc_167"
)

func main() {
	input_1 := []int{2, 7, 11, 15}
	input_2 := 9
	fmt.Println("input_1:", input_1)
	fmt.Println("input_2:", input_2)
	fmt.Println("answer:", lc_167.Answer(input_1, input_2))
}
