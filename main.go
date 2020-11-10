package main

import (
	"fmt"
	lib "random/lib"
)

func init() {
	lib.InitMyConfig()
}

func main() {
	nums := lib.CreatRandomNumber(0, 10, 10)
	for x := 0; x < len(nums); x++ {
		fmt.Printf("index=%d: [%d號座位, 是%s]\n", x, nums[x], lib.MyConfig.Seat[nums[x]])
	}
	//fmt.Println(len(lib.MyConfig.Seat))
	// for i := 0; i < len(lib.MyConfig.Seat); i++ {
	// 	fmt.Println(lib.MyConfig.Seat[nums[i]])
	// }
}
