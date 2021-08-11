package main

import "fmt"

func changeI(i int) int{
	i=13
	return i
}
func change(i int){
	i=13
}
func changeP(i *int){
	*i=15
}
func main(){
	i:=12
	change(i)
	fmt.Println(i)
	i=changeI(i)
	fmt.Println(i)
	changeP(&i)
	fmt.Println(i)
}
