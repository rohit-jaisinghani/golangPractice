package main

import "fmt"

func main(){
	var arr []int

	//var count []int
	var counter int
	arr=[]int{1,2,5,5,6,6,7,2,2}
	for i:=0;i<len(arr)-1;i++ {
		for j:=i+1;j<len(arr);j++ {
			if arr[i]==arr[j]  {
				fmt.Println("Duplicate element",arr[i])
				counter=arr[i]
			}
			if counter==arr[j]{
				fmt.Println("index of other duplicates",j)
			}
		}

	}


}
