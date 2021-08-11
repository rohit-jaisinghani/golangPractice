package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error){
	if e!=nil{
		panic(e)
	}
}
func main(){
f,err:=os.Create("test.txt")
check(err)
defer  f.Close()
d2:=[]byte{115,111,109}
n2,err:=f.Write(d2)
check(err)
fmt.Printf("%T\n",n2)

n3,err:=f.WriteString("Hello Rohit")
check(err)
fmt.Println(n3)
f.Sync()

	data, err := ioutil.ReadFile("test.txt")
	fmt.Println(string(data))


}
