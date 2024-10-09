package main

import (
	"fmt"
)

func main(){
	target_addres := "scanme.nmap.org"
	target := NewTarget(target_addres, []string{ "1", "2", "3", "80" }, nil)
	fmt.Println(target)
	for _, i := range target.ScanAddressOnPorts(){
		fmt.Println(i)
	}
}
