package main

import (
	"fmt"
	"strings"
)

func main(){
	str:="stresster"
	_=str
	chars:=[]rune(str)
	var s []string
	var s1 []string
	var str3 string
	for i:=0;i<len(chars);i++ {
		for j:=i+1;j<len(chars);j++{
			if(chars[i]==chars[j]){
				if !(strings.Contains(str3,string(chars[i]))){
					s = append(s, string(chars[i]))
					str3 = strings.Join(s, "")
				}
			}else{
				s1=append(s1,string(chars[i]))
			}
		}
	}
	var s4 []string
	for i:=0;i<len(s1);i++{
		if !strings.Contains(str3,s1[i]){
				s4=append(s4,s1[i])
		}
	}
fmt.Println("Repeatative chars:",str3)
	//fmt.Println(s1)
	fmt.Println("Non repeatative chars",s4)

}
