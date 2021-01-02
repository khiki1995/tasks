package main

import (
	"fmt"
	"os"
)

func main() {
	arr := []int{-6, 1, 1, 40} 
	for ind, val := range arr {
		if ind == val {
			fmt.Println(val)
			os.Exit(1)
		}
	}
	fmt.Println(false)
}