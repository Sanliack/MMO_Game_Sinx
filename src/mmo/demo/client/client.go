package main

import "fmt"

func main() {

	gg := map[string]int{}
	gg["试试"] = 2
	gg["萨格"] = 23
	gg["埃是法国"] = 25
	gg["爱国"] = 263
	gg["活动方式"] = 82
	gg["aa"] = 92
	gg["11"] = 102
	//aa := []int{2, 4, 3, 6, 7}
	for i, v := range gg {
		fmt.Println(i, v)
	}
}
