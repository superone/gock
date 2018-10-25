package command

import(
	_ "path/filepath"
	_"strings"
	"os"
	"fmt"
	"gock"
	"reflect"
)

func Go(){
	command := "run"
	var params []string
    for i, arg := range os.Args[1:] {
		if (i==0){
			command = arg 
		}else{
			params = append( params , arg )
		}
        //fmt.Println( i , "::arg =", arg, " arg Type =", reflect.TypeOf(arg))    // 可以把每个字符串参数转换成我们需要的类型
	}
	
	switch command {
		case "run":
			gock.GockStart(params)
		default :
			fmt.Println("None command.")
	}
	
}