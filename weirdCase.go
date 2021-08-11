package main

import (
	"fmt"
	"strings"
)

func main(){
	str:="Hello World" //HeLlO WoRlD"

	chars := []rune(str)
	for i,v :=range chars{
		if(i%2==0){
		strings.ToUpper(string(v))
		}else {
			strings.ToLower(string(v))
		}
	}
	fmt.Println(string(chars))
}
