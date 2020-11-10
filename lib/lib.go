package lib

import (
	"fmt"
	"math/rand"
	"time"

	viper "github.com/spf13/viper"
)

//myConfigStruct 我自定義的config結構
type myConfigStruct struct {
	Seat  []string
	Eat   []string
	Drink []string
}

//MyConfig 自定義的config變數
var MyConfig myConfigStruct

//InitMyConfig 我的設定檔
func InitMyConfig() {
	viper.SetConfigName("data")                    // 指定文件的名稱
	viper.AddConfigPath("/Projects/random/config") // 配置文件和執行檔目錄
	err := viper.ReadInConfig()                    // 根據以上定讀取文件
	if err != nil {
		fmt.Println("Fatal error config file" + err.Error())
	}
	MyConfig.Seat = append(MyConfig.Seat, viper.GetString("TABLE.name1"))
	MyConfig.Seat = append(MyConfig.Seat, viper.GetString("TABLE.name2"))
	MyConfig.Seat = append(MyConfig.Seat, viper.GetString("TABLE.name3"))
	MyConfig.Seat = append(MyConfig.Seat, viper.GetString("TABLE.name4"))
	MyConfig.Seat = append(MyConfig.Seat, viper.GetString("TABLE.name5"))
	MyConfig.Seat = append(MyConfig.Seat, viper.GetString("TABLE.name6"))
	MyConfig.Seat = append(MyConfig.Seat, viper.GetString("TABLE.name7"))
	MyConfig.Seat = append(MyConfig.Seat, viper.GetString("TABLE.name8"))
	MyConfig.Seat = append(MyConfig.Seat, viper.GetString("TABLE.name9"))
	MyConfig.Seat = append(MyConfig.Seat, viper.GetString("TABLE.name10"))

	

}

//CreatRandomNumber 生成不重複亂數的array 參數說明 (起始數字 結束數字 array大小)
func CreatRandomNumber(start int, end int, count int) []int {
	//防呆檢查
	//第2個條件: 起始1 結尾4 大小要5 一定會重複
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)

	for len(nums) < count {
		/*
			隨機的偽隨機數
			我們已經知道了預設隨機種子是從 1 開始，那麼我們只要在每次生成隨機數之前先設定一個不一樣的種子，那麼其結果自然也就不一樣了。
			我們要儘可能保證每次偽隨機數生成器工作時使用的是不同的種子，通常的做法是採用當前時間作為種子。
		*/
		rand.Seed(time.Now().UnixNano())

		//產生亂數
		num := rand.Intn((end - start)) + start

		//檢查重複
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}
