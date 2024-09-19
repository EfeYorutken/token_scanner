package main

import (
	"fmt"
	"os"
)


func string_to_int(s string) int{
	res := 0

	for i := 0; i < len(s); i++{
		res = res*10 + int(s[i] - '0')
	}

	return res
}

func main(){

	//	args := os.Args[1:]
	//
	//	target := args[0]
	//	ports := map_to(args[1:], string_to_int);
	//
	//	targ := NewTarget(target, ports, nil)
	//
	//	for _,a := range targ.ScanAddressOnPorts(){
	//		fmt.Println(a)
	//	}

	for _, a := range execute_tool(os.Args[1:]){
		fmt.Println(a)
	}

}
