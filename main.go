package main

import (
	"fmt"
)

func main(){
		res := ScanFromFile("test.txt")
		for _, r := range res{
			fmt.Println(r)
		}
}
