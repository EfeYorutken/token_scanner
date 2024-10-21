package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	args_and_responsibilities := GetArgsAndResponsibilities(os.Args)

	var targets []Target
	var results []string

	targets = append(targets, Target{os.Args[1],nil,nil})

	run_good_script := false
	run_bad_script := false

	var script_path string


	for _, a_and_r := range args_and_responsibilities{
		arg := a_and_r[0]
		resp := a_and_r[1:]

		switch(arg){
		case "-p"://ports
			targets[len(targets) - 1].ports = append(targets[len(targets) - 1].ports, resp...)
			break
		case "-r"://range of ports
			begin,_ := strconv.Atoi(resp[0])
			end,_ := strconv.Atoi(resp[1])
			for i := begin; i < end; i++{
				targets[len(targets) - 1] = targets[len(targets) - 1].get_target_w_ports([]string{strconv.Itoa(i)})
			}
			break
		case "-sG"://script when successfull scan
			run_good_script = true
			script_path = resp[1]
			break
		case "-sB"://script when unsuccessfull scan
			run_bad_script = true
			script_path = resp[1]
			break
		case "-s"://script regardless of success
			run_bad_script = true
			run_good_script = true
			script_path = resp[1]
			break
		case "-f"://read from file and scan
			targets = append(targets, GetTargetsFromFile(resp[0])...)
			break
		case "-t"://type of protocol
			targets[len(targets) - 1].protocols = append(targets[len(targets) - 1].protocols, resp...)
			break
		}

	}

	//should be doable with map_to
	for i, t := range targets{
		results = append(results, t.ScanAddressOnPorts()...)
		fmt.Printf("scanning target %v which is %v\n", i, t)
	}

	if run_bad_script && !run_good_script{
		IfNotSuccess(results, script_path)
	}else if !run_bad_script &&  run_good_script{
		IfSuccess(results, script_path)
	}else if run_bad_script &&  run_good_script{
		EitherWay(results, script_path)
	}else{
		for _, a := range results{
			fmt.Println(a)
		}
	}

}
