package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){

	if(len(os.Args) < 2){
		fmt.Println(`
-p : port scan
	$> token_scanner -p port1 port2 port3
	scans only the specified ports

-r : range scan
	$> token_scanner -r begin end
	scans a range of ports between "begin" and "end" 

-t : type of scan
	$> token_scanner -t type
	scans the given target with the protocol "type" in mind, ie scans target abc on port xyz with protocol type

-sG : good script
	$> token_scanner -sG command_that_runs_script
	runs the command_that_runs_script when a successful scan is found
	the command is ran in the format "command_that_runs_script address port protocol"

-sB : bad script
	$> token_scanner -sB command_that_runs_script
	same as -sG but for unsuccessful scans

-f : file
	$> token_scanner -f file_name
	reads the targgets from file_name and scans them
	the targets must be represented in the format of address port protocol
	if you want the default values for port or protocol, put a * charcter

-s : script
	$> token_scanner -s command_that_runs_script
	same as -sG but for all scans regardless of success
		`)
		os.Exit(0)
	}

	args_and_responsibilities := GetArgsAndResponsibilities(os.Args)

	var targets []Target
	var results []string

	targets = append(targets, Target{os.Args[1],[]string{},[]string{}})

	run_good_script := false
	run_bad_script := false
	run_both := false

	var script_runner_good string
	var script_runner_bad string
	var script_runner_both string

	if(targets[0].name == "-f"){
		targets = []Target{}
	}

	for _, a_and_r := range args_and_responsibilities{
		arg := a_and_r[0]
		resp := a_and_r[1:]

		if(len(resp) == 0){
			fmt.Println("no argument to ", arg, "was given. Quitting...")
			os.Exit(1)
		}

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
			script_runner_good = strings.Join(resp, " ")
			break
		case "-sB"://script when unsuccessfull scan
			run_bad_script = true
			script_runner_bad = strings.Join(resp, " ")
			break
		case "-s"://script regardless of success
			run_both = true
			script_runner_both = strings.Join(resp, " ")
			break
		case "-f"://read from file and scan
			targets = append(targets, GetTargetsFromFile(resp[0])...)
			break
		case "-t"://type of protocol
			targets[len(targets) - 1].protocols = append(targets[len(targets) - 1].protocols, resp...)
			break
		}
	}

	targets = map_to(targets, func(t Target)Target{
		if(len(t.ports) == 0){
			t = t.get_target_w_ports(nil)
		} 
		if(len(t.protocols) == 0){
			t = t.get_target_w_protocols(nil)
		}
		return t
	})


	//should be doable with map_to
	for _, t := range targets{
		results = append(results, t.ScanAddressOnPorts()...)
	}

	if !run_good_script && !run_bad_script && !run_both{
		for _, r := range results{
			fmt.Println(r)
		}
		return
	}

	if run_good_script{
		IfSuccess(results, script_runner_good)
	}
	if run_bad_script{
		IfNotSuccess(results, script_runner_bad)
	}
	if run_both{
		EitherWay(results, script_runner_both)
	}

	fmt.Println("is go happy now?")

}
