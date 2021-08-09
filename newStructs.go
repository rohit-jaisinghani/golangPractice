package main

import "fmt"

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}
func main() {
	type Parent struct {
		motherName, fatherName string
	}
	type Employee struct {
		name, email string
		age         int
		parentInfo  Parent
	}
	rohit := Employee{
		name:  "Rohit",
		email: "xyz@gmail.com",
		age:   24,
		parentInfo: Parent{
			motherName: "Bhavna",
			fatherName: "Mahesh",
		},
	}
	sneha := rohit
	fmt.Printf("%+v\n", rohit)
	fmt.Printf("%+v\n", sneha)
	sneha = Employee{
		name: "Sneha",
	}
	fmt.Printf("%+v\n", rohit)
	fmt.Printf("%+v\n", sneha)

	for index,name:=range rohit.name{
		fmt.Printf("%d -> %q\n", index, name)
	}
	func(msg string) {
		fmt.Println(msg)
	}("I'm an anonymous function!") // calling the anonymous function

	a := increment(10)
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	type grade struct{
		grade,course string
	}
	type student struct{
		name string
		grades grade
	}
	you:=student{
		name:"rohit",
		grades:grade{
			grade:"90",
			course:"GO",
		},
	}
	fmt.Printf("%+v\n",you)


}
