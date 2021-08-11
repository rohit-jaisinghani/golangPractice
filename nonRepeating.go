package main

import "fmt"

func main(){
	str:="stress"
	_=str
	chars:=[]rune(str)
	var s []string
	var b bool=false
	for i:=0;i<len(chars);i++ {
		for j:=i+1;j>len(chars);j++{
			if (chars[i]==chars[j]) {
				b=true
				continue
			}
		}
		if(b==false){
			s=append(s,string(chars[i]))
		}

	}
	fmt.Println(s)
}
