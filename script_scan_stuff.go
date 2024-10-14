package main

import (
	"fmt"
	"os/exec"
	"strings"
)


type Result struct {
	address string
	port string
	protocol string
	successfull bool
}

func result_from_text(txt string) Result{
	//hacky way to get the address, ip and protocol
	arr := strings.Split(txt, " ON ")
	last_elem := arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	last_elem_splitted := strings.Split(last_elem, " with ")
	arr = append(arr, last_elem_splitted[0])
	arr = append(arr, last_elem_splitted[1])
	successfull := arr[0] == "SUCCESSFUL SCAN"

	return Result{strings.Split(arr[1], ":")[0], strings.Split(arr[1], ":")[1], arr[2], successfull}
}

//will call the script as "script.sh address port protocol"
func IfSuccess(results []string, command_that_runs_the_script string){

	for _, result := range results{


		current_result := result_from_text(result)
		//should probably determine the shell to be used dynamically
		//alternatively different versions of script_scan_stuff for different OSs can exist but seems ugly

		if current_result.successfull{
			arr := strings.Split(command_that_runs_the_script," ")
			cmd := exec.Command(arr[0],arr[1], current_result.address, current_result.port, current_result.protocol)

			out,err := cmd.Output()

			if err != nil{
				fmt.Println(err)
			}



			fmt.Println(string(out))


		}

	}

}

func IfNotSuccess(results []string, command_that_runs_the_script string){

	for _, result := range results{


		current_result := result_from_text(result)
		//should probably determine the shell to be used dynamically
		//alternatively different versions of script_scan_stuff for different OSs can exist but seems ugly

		if !current_result.successfull{
			arr := strings.Split(command_that_runs_the_script," ")
			cmd := exec.Command(arr[0],arr[1], current_result.address, current_result.port, current_result.protocol)

			out,err := cmd.Output()

			if err != nil{
				fmt.Println(err)
			}



			fmt.Println(string(out))


		}

	}

}


func EitherWay(results []string, command_that_runs_the_script string){

	for _, result := range results{


		current_result := result_from_text(result)
		//should probably determine the shell to be used dynamically
		//alternatively different versions of script_scan_stuff for different OSs can exist but seems ugly

		arr := strings.Split(command_that_runs_the_script," ")
		cmd := exec.Command(arr[0],arr[1], current_result.address, current_result.port, current_result.protocol)

		out,err := cmd.Output()

		if err != nil{
			fmt.Println(err)
		}

		fmt.Println(string(out))

	}

}
