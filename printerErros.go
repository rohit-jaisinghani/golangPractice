package main

import (
	"fmt"
	"strconv"
)

func main(){
	str:="aabbccddeettuuummm"
	badColors:=0
	total:=0
	for _,value:=range str {
		v:=string(value)
		if v<"a" || v>"m"{
			badColors++
		}
		total++
	}
	fmt.Println(strconv.Itoa(badColors)+"/"+strconv.Itoa(total))
}
