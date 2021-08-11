package main

import "fmt"

func main(){
	var counter int
	num:=1234
	for num>0{
		counter+=num%2
		num/=2
		fmt.Println(counter)
	}
	fmt.Println(counter)
}
