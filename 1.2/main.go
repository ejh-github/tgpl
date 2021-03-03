package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Executing: ", os.Args[0])
	fmt.Println("Echoing: ", strings.Join(os.Args[1:], " "))
	for i := 0; i < len(os.Args); i++ {
		fmt.Println("Index: ", i, " = ", os.Args[i])
	}
}
