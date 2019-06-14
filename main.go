package main

import (
	"C"
	"fmt"
	"os"

	"github.com/kaz/sql-mask/mask"
)

//export masked
func masked(sql *C.char) *C.char {
	result, err := mask.Mask(C.GoString(sql))
	if err != nil {
		result = "[ parse error. masked whole sql. ]"
	}
	return C.CString(result)
}

func main() {
	result, err := mask.Mask(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
