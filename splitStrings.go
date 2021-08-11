package main

import "fmt"

func main(){
	str:="abcde"
	result:=[]string{}
	counter:=0
	tempStr:=""
	lastChar:=' '
	for _,v:=range str{
		counter++
		tempStr+=string(v)
		lastChar=v

		if counter==2{
			result=append(result,tempStr)
			counter=0
			tempStr=""
		}
	}
	if counter==1{
		result=append(result,string(lastChar)+"_")
	}
	fmt.Println(result)
	}



