package main

import (
	"fmt"
)
type cb func(int)int

func main(){
    call(1,test)
	call(2,func(x int)int{
		fmt.Println(x)
		return 0
	})

}
func call(x int, f cb){
	f(x)
}
func test(a int) int {
	fmt.Println(a)
	return 0

}