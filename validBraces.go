package main

import (
	"fmt"
	"strings"
)

func main(){
	var str []string=[]string{"}","{","[","]","(",")","{","}"}
	b:=false
	var count1,count2,count3 int
	for i:=range str{
		if str[i]=="}"{
			strings.ReplaceAll(str[i],"}","{")
			count1++
		}else if str[i]=="]"{
			strings.ReplaceAll(str[i],"]","[")
			count2++
		}else if str[i]==")" {
			strings.ReplaceAll(str[i],")","(")
			count3++
		}
	}
	fmt.Println(count3,count2,count1)
	num:=count1+count2+count3
	if len(str)%num==0{
		b=true
	}
	fmt.Println(b)
}
