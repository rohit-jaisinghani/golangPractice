package main

import "fmt"

func main(){
	//var sl=make([]int,4)
	sl:=[]int{}
	fmt.Println(sl)
	fmt.Println(len(sl))
	fmt.Println(cap(sl))
	sl=[]int{1,2,3,4}
	fmt.Println(cap(sl))
	sl=append(sl,5)
	sl=append(sl,6)
	fmt.Println(cap(sl))

	var mp=make(map[string]int)
	_=mp
	mpp:=map[string]int{}
	mpp["A"]=12
	mpp["b"]=13
	mpp["B"]=14
	mpp["C"]=15
	c:="C"
	_=c
	var mpc=make(map[string]int)

	for k,v:=range mpp{
		fmt.Println(k,v)
		mpc[k]=mpp[k]
		if(k==c){
			fmt.Println(k,mpp[k])
		}
	}
	fmt.Println(mpc)



}
