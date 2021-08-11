package main

import "fmt"

func main(){
	var a []int
	//var count int
	a=[]int{1,1,2,1,1}
	for i:=0;i<len(a);i++ {
			for j:=i+1;j<len(a);j++ {
				if a[i] != a[j] {
					fmt.Println(a[j])
					break
				}
			}
			break
		}

	}

