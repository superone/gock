package command

import(
	_ "path/filepath"
	"strings"
	"os"
	"fmt"
	"gock"
	"reflect"
)

func Go(c chan int){
	command := ""
	var params []string
    for i, arg := range os.Args[1:] {
		if (i==0){
			command = arg 
		}else{
			params = append( params , arg )
		}
        fmt.Println( i , "::arg =", arg, " arg Type =", reflect.TypeOf(arg))    // 可以把每个字符串参数转换成我们需要的类型
	}
	fmt.Println(params)
	if( strings.EqualFold(command,"") ||  strings.EqualFold(command,"run") ){
		fmt.Println(strings.EqualFold(command,""))
		gock.GockStart()
	}
	c <- 2
}