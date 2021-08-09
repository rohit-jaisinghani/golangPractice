package main

import (
	"fmt"
	"time"
)
func sum(a ... int) int{
	s:=0
	for _,v:=range a{
		s=s+v
	}
	return s
}

func searchItem(s []string,st string) bool{
	for _,k:=range s{
		if k==st {
			return true
		}
	}
	return false
}
func main() {
	sum:=sum(1,2,3)
	var animals []string=[]string{"lion","pig","fox"}
	var d bool=searchItem(animals,"lion")
	fmt.Println(d)
	var e bool=searchItem(animals,"dear")
	fmt.Println(e)
	fmt.Println("SUM IS",sum)
	 t := time.Date(2015, 02, 21,23,35,0,0,time.UTC)
	fmt.Println(t)
	fmt.Println("Hello Calc")
	var i int = 12
	for i = 0; i < 24; i++ {
		if i%3 == 0 {
			fmt.Println("number is divisible by 3", i)
		} else if i%5 == 0 {
			fmt.Println("number is divisible by 5", i)
		} else {
			fmt.Println("number is not divisible by 3 and 5", i)
		}
	}
	var k int
	for k=0;k<5;k++{
		if(k==2){
			continue

		}
		fmt.Println("no is not 2")
	}
	j := 12
	switch j {
	case 12:
		fmt.Println("number is", j)
	case 11:
		fmt.Println("number is 11")
	default:
		fmt.Println("default number")

	}
}
