package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	num, _ := strconv.Atoi(arg)

	for i := 1; i <= num; i++ {
		fmt.Println(i)
	}	 
}