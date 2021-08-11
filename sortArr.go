package main

import (
	"fmt"
	"sort"
)

func main(){
	a := []int{3,4,5,7,6}
	sort.Slice(a, func(i, j int) bool {
		fmt.Println("Index of i",a[i])
		fmt.Println("Index of j",a[j])
		fmt.Println(a[i] < a[j])
		return a[i]<a[j]

	})
	fmt.Println(a)
	/*for _, v := range a {
		fmt.Println(v)
	}*/
}
