package main

import "fmt"
func fact(i int) int{
	if(i==0){
		return 1
	}
	return i*fact(i-1)
}
func main(){
	fmt.Println("enter the number")
	var i int
	fmt.Scanln(&i)

	n:=fact(i)
	fmt.Println(n)
}
