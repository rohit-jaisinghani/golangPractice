package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	str:="1.02.3.4"
	var s []string
	var b bool=true
	s=strings.Split(str,".")
	fmt.Println(s)
	for i:=0;i<len(s);i++{
		d,err:=strconv.Atoi(s[i])
		_=err
		if len(s)<4{
			b=false
		}
		if strings.HasPrefix(s[i],"0") && !strings.HasSuffix(s[i],"0"){
			b=false
		}
			if !(d>0 && d<255){
				b=false
			}
		}

	fmt.Println(b)
	}


