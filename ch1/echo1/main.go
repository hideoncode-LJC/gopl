package main

import (
	"os"
	"fmt"
)




func main() {
	var s, seq string

	l := len(os.Args)

	fmt.Println("len = ", len(os.Args))

	for i := 0 ; i < l ; i++ {
		s += seq + os.Args[i]
		seq = ""
	}

	fmt.Println(s)
}
