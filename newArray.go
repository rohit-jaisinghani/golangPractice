package main

import "fmt"

func main(){
	var arr [3]string=[3]string{"Hello","World"}
	_=arr

	b:=[...]bool{}
	_=b
	myArray := [4]float64{1.2, 5.6}

	myArray[2] = 6

	a := 10.0

	myArray[0] =a

	myArray[3] = 10.10

	fmt.Println(myArray)

	var s []string=[]string{"a","b","c"}
	var i int
	for i=0;i<len(s);i++ {
		fmt.Println(s[i])
	}
	nums := []float64{1.1, 2.2, 3.3}
	nums = append(nums, 10.1)
	nums = append(nums, 4.1, 5.5, 6.6)
	fmt.Println(nums)
	n := []float64{10.10, 20.20}
	nums = append(nums, n...)
	fmt.Println(nums)

	mySlice := []float64{1.2, 5.6}
	mySlice[0] = 6
	a1 := 10.0
	mySlice[0] = a1
	mySlice[1] = 10.10
	mySlice = append(mySlice, a1)
	fmt.Println(mySlice)

	friends := []string{"Marry", "John", "Paul", "Diana"}
	people:=make([] string,len(friends))
	nn:=copy(people,friends)
	fmt.Println(nn)

	nums_s := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	sum := 0
	for _, v := range nums_s[2 : len(nums_s)-2] {
		fmt.Println(v)
		sum += v
	}
	fmt.Println("Sum:", sum)

}
