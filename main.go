package main

import (
	"fmt"
)

func main(){
	target_addres := "scanme.nmap.org"
	target := NewTarget(target_addres, nil, nil)
	for _, i := range target.ScanAddressOnPorts(){
		fmt.Println(i)
	}
}
