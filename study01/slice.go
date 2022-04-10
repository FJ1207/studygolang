package main

import (
	"fmt"
)
func main(){
	//var slice = []int{}
	var slice01= make([]int,5,6)//创建切片
	slice01[1] = 2
	slice01=append(slice01,5,7)//020005切片翻倍
	//slice = append(slice,1,2)
	fmt.Println(slice01)
	fmt.Printf("len(x)=%d,cap(x)=%d",len(slice01),cap(slice01))

}