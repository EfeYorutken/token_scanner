package main

import (
	"fmt"
)

func main(){
	results := Target{"scanme.nmap.org", []string{"10","20","80"}, []string{"tcp"}}.ScanAddressOnPorts()

	fmt.Println("doing the good shit")
	IfSuccess(results, "python /home/efey/home/projects/token_scanner/good.py")


	fmt.Println("doing some bad shit")
	IfNotSuccess(results, "python /home/efey/home/projects/token_scanner/bad.py")

	fmt.Println("doing some nutural shit")
	EitherWay(results, "python /home/efey/home/projects/token_scanner/good.py")

}
