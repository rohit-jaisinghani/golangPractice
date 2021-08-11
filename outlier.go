package main

import "fmt"

func main(){
	var arr []int=[]int{3,5,4,9}
	_=arr
	var counte []int=[]int{}
	var counto []int=[]int{}
	var even,odd int
	for i:=0;i<len(arr);i++ {
		if(arr[i]%2==0){
			counte=append(counte,i)
			even++
		}else {
			counto=append(counto,i)
			odd++
		}
	}
	if (even>odd){
		fmt.Println(arr[counto[0]])
	}else{
		fmt.Println(arr[counte[0]])
	}
}
