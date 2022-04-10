package main

import "fmt"

func main(){
	_,a,b:=abc()//函数体内声明必须得用，外不必:=直接声明
	fmt.Println(a,b)
}
func abc()(int,int,string){
	return 1,1,"str"


}