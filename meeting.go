package main

import (
	"fmt"
	"sort"
	"strings"
)
type Person struct {
	LastName string
	FirstName string
}
func main() {
	persons := []Person{}
	str := "Ramesh:J;Yogesh:M;Suresh:G;Bhavesh:P;Durgesh:K"
	names := strings.Split(str, ";")
	fmt.Println(names)
	for _, name := range names {
		full_name := strings.Split(name,":")
		p := Person{strings.ToUpper(full_name[1]),strings.ToUpper(full_name[0])}
		persons = append(persons, p)
	}
	sort.Slice(persons, func(i, j int) bool{
	if strings.EqualFold(persons[i].LastName,persons[j].LastName){
		if strings.Compare(persons[i].FirstName,persons[j].FirstName) >=1{
			return false
		}else {
			return true
		}
}
	if strings.Compare(persons[i].LastName,persons[j].LastName)>=1{
		return false
	} else{
		return true
	}
	})
	result:=""
	for _,person:=range persons{
		result+="("+person.LastName+","+person.FirstName+")"
	}
	fmt.Println(result)
}
