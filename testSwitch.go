package main

import (
	"fmt"
	"time"
)

func main(){
	var i=9
	switch i {
	case 9:
		fmt.Println("Number is 9")
	default:
		fmt.Println("Number is not 9")
	}

	t:=time.Now()
	switch  {
	case t.Hour()<12 :
		fmt.Println("Its before noon")
	default:
			fmt.Println("Post Noon")
	}
	varType := func( i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("type is boolean")
		case int:
			fmt.Println("type is int")
		}
	}
		varType(true)
		varType(1)


}
