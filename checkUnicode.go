package main

import (
	"fmt"
	"unicode"
)

func main(){
	str:="ac1%%$23c"
	var b bool=true
	for _,c:=range str{
		if unicode.IsDigit(c) || unicode.IsLetter(c){
			continue
		}else {
			b=false
		}
	}

	fmt.Println(b)
}
