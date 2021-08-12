package main

import "fmt"

func push(s []int,i int) []int{
	s=append(s,i)
	return s
}
func pop(s []int,i int) []int{
s=s[:i]
return s
}
func main(){
	
	var s []int=[]int{6,7,8}
	num:=len(s)-1

	s=pop(s,num)
	fmt.Println(s)
	j:=9
	kl:=push(s,j)

	fmt.Println(kl)


}
