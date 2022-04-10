package main

import ("fmt")
//var(sto int
//	  ur1 string)一起声明
//sto :=123只能在函数体
func main(){

	sto :=123
	var ur1 ="golang"
	var url ="cd=%d,name=%s"
	var last=fmt.Sprintf(url,sto,ur1)
	fmt.Println(last)
}