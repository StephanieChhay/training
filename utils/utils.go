package utils

import (
	"fmt"
	"math/rand"
)

func ReadValue() int {
	var value int
	fmt.Print("enter a integer value: ")
	_, err := fmt.Scanf("%d", &value)
	if err != nil {
		fmt.Println(err)
	}
	return value
}

func Random(min, max int) int {
	return rand.Intn(max+1-min) + min
}
