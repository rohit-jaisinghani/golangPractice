package main

import (
	"fmt"
)

func main(){
	str:="Hello World"
		chars:=[]rune(str)
		firstChar:=""
		secondChar:=""
		lastChar:=""
		result:=""


	for i:=0;i<len(chars);i++{
		if i==0{
		firstChar=fmt.Sprintf("%d",chars[i])
		}else if i==1{
			secondChar=string(chars[i])
		}else if i==len(chars)-1{
			lastChar=string(chars[i])
		}else {
			result+=string(chars[i])
		}
	}

result=lastChar+result+secondChar
result=firstChar+result
fmt.Println(result)
}
