package main

import (
	"fmt"
)

func main(){
	test := "hello there"
	for _, c := range test{
		fmt.Println(c)
	}
}
