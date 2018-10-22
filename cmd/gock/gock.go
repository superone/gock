package main;

import(
	"fmt"
	"command"
)


func main(){
	ch := make(chan int)
	go command.Go( ch )
	//go done( ch )
	fmt.Println( <-ch )
}

func done( c chan int){
	c <- 5
}