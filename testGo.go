package main

import (
	"fmt"
	"math"
)


func main(){

	var x float64=func (a float64) float64{
		return math.Sqrt(a)
	}(64.0)
	fmt.Println(x)

}
