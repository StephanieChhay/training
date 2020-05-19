package main

import (
	"fmt"

	"github.com/lambher/training/utils"
)

func main() {
	value := utils.ReadValue()
	fmt.Println(value)
	for i := 0; i < 10; i++ {
		random := utils.Random(0, 1)
		fmt.Println(random)
	}
}
