package main

import (
	"fmt"
	"math"
)

const k="HELLO"
func main(){
	const k="hello"
	const j=20
	var i=math.Sin(5000)
	fmt.Println(i)
	//i int=math.Sin(4000)

	for {
		fmt.Println("FOR LOOP")
		//return return ke baad kuch nhi chalega next for loop nhi
		break
	}
	for i:=0;i<5;i++ {
		if(i==0) {
			fmt.Println("NUmber is zero")
		}else if j%i==0 {
			fmt.Println("Number is divisble")
		}else {
			fmt.Println("number is not divisble")
		}
	}

}