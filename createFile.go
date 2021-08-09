package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var newFile *os.File
	_ = newFile
	var err error
	_ = err

	/*newFile,err=os.Create("a.txt") //this will create a file
	if err!= nil{
		log.Fatal(err)
	}*/
	//file,err=os.Open("a.txt")
	bs := []byte("go program")
	err = ioutil.WriteFile("a.txt", bs, 0644) //write to file
	data, err := ioutil.ReadFile("a.txt")     //read a file
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)
	fmt.Println()
	s:="ma hello"
	x:=s[0]
	s1:="mama"
	s2:=s+s1
	fmt.Println(s2)
	fmt.Println(string(x))

	x,y:=10,2
	ptrx,ptry:=&x,&y
	fmt.Println(*ptrx,*ptry)
	z:=(int)(*ptrx)/(*ptry)
	fmt.Println(z)



}
