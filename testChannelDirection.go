package main

import (
	"fmt"
)

func send(s chan<- string,str string){
s<-str
}
func rec(s <-chan string,r chan<- string){
	mesg:=<-s
	r<-mesg


}
func main(){
	s:=make(chan string,1)
	r:=make(chan string,1)
	send(s,"Hii")
	rec(s,r)
	fmt.Println(<-r)

}
