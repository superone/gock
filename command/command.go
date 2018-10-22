package command

import(
	"path/filepath"
	"os"
	"fmt"
	"gock"
)

func Go(){
	fmt.Println(os.Args)
	fmt.Println(os.Args[:0])
	fmt.Println(os.Args[:1])
	fmt.Println(os.Args[:2])
	//fmt.Println(os.Args[:3])

	
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)

	gock.GockStart()
}