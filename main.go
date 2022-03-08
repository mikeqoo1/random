package main

import (
	"fmt"
	lib "random/lib"
)

func init() {
	lib.InitMyConfig()
}

func main() {
	nums := lib.CreatRandomNumber(0, 10, 1)
	for x := 0; x < len(nums); x++ {
		//fmt.Printf("index=%d: [%d號座位, 是%s]\n", x, nums[x], lib.MyConfig.Seat[nums[x]])
		fmt.Printf("今天, 是吃%s\n", lib.MyConfig.Eat[nums[x]])
		fmt.Printf("今天, 是喝%s\n", lib.MyConfig.Drink[nums[x]])
	}
}
