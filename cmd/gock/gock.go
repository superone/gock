package main;

import(
	_"fmt"
	"command"
)


func main(){
	command.Go()
}

func done( c chan int){
	c <- 5
}