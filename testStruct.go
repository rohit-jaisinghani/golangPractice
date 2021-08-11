package main

import "fmt"

type employee struct{
	name string
	age int
}
func changeAge(e *employee){
	e.age=26
}
func main(){
e:=employee{
	name:"John",
	age:25,
}
	f:=e
changeAge(&e)
fmt.Println(e)
	fmt.Println(f)
}
