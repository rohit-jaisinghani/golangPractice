package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello Go")
	var kl int = 5
	_ = kl
	var i [3]int = [3]int{1, 2}
	var a [3]int = [3]int{}
	a = i
	a[0] = 10
	var k *int = &a[2]
	*k = 20
	fmt.Println(i)
	fmt.Println(a)
	var b []int = []int{1, 2}
	var c []int = b
	var d []int = []int{} //empty slice
	_ = d                 //if we are not using d variable
	d = append(d, b[0:2]...)
	c[0] = 22  //dono slices mai change hoga(b and c)
	d[0] = 123 //sirf d slice mai change hoga
	fmt.Println(b)
	fmt.Println(c)
	var cd = map[string]int64{}
	cd["A"] = 1
	cd["B"] = 2
	fmt.Println(cd)
	var ef = make(map[string]int64)
	for k, v := range cd {
		ef[k] = v
	}
	fmt.Println(ef)
	var de = map[string]int64{}
	de = cd
	de["A"] = 11 //dono maps mai change hoga(cd and de)
	fmt.Println(de)
	fmt.Println(cd)

	var j=10
	_=j
	K:="HELLO"
	strings.ToUpper(K)


}
