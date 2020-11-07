package main

import (
	"fmt"
	lib "random/lib"
)

func init() {
}

func main() {
	nums := lib.CreatRandomNumber(0, 7, 7)
	for i := 0; i < 5; i++ {
		nums = lib.CreatRandomNumber(0, 7, 7)
		fmt.Println(nums)
	}
}
