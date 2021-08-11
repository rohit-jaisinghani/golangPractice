package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println("enter the size")
	var i,k int
	fmt.Scanln(&i)
	var arr=make([]int,i)
	fmt.Println("enter the numbers ")
	for j:=0;j<i;j++ {
		fmt.Scanln(&k)
		arr[j]=k
	}
	for j:=0;j<i;j++ {
		sort.Ints(arr)
	}
var largest [2]int=[2]int {arr[len(arr)-1],arr[len(arr)-2]}
fmt.Println(largest)

}
