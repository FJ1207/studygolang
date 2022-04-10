package main

import "fmt"

func main(){
	const a int = 1
	const b = 3
	//var area=a*b
	fmt.Println(a)
	const (c =iota
		   d =iota)
	fmt.Println(c,d)//iota++

}