package main

import (
	"fmt"
	"os"
	"strings"
)

func get_file_content(file_name string)string{
	file_content, err := os.ReadFile(file_name)
	if err != nil{
		fmt.Printf("file %s not found", file_name)
		os.Exit(1)
	}
	return string(file_content)
}

//file syntax should be target_name port protocol
//for now
func lines_to_targets(lines []string)[]Target{
	var res []Target
	temp := map[string]Target{}
	for i := 0; i < len(lines); i++{
		if len(lines[i]) != 0{
			line := lines[i]
			split_line := strings.Split(line, " ")
			target_name := split_line[0]
			port :=split_line[1]
			protocol := string(split_line[2])

			
			//if the current target address is in temp
			//ie the address was listed in the file more than once
			//perhaps with different ports and protocols
			if _, in_map := temp[target_name];in_map{
				old_target := temp[target_name]

				//if the port exists in the old target, dont change the ports
				new_ports := ternary(
					contains(old_target.ports, port), old_target.ports,
					push(old_target.ports, port))//otherwise push the new port to the old target.ports

					//if the protocol exists in the old target, dont change the protocols
				new_protocols := ternary(
					contains(old_target.protocols, protocol), old_target.protocols,
					push(old_target.protocols, protocol))//otherwise push the new protocol to the old target.protocols

				temp[target_name] = NewTarget(target_name, new_ports, new_protocols)

			}else{
				temp[target_name] = NewTarget(target_name, []string{port}, []string{protocol})
			}
		}
	}
	for _, t := range temp{
		res = append(res, t)
	}

	return res
}

func ScanFromFile(filename string)[]string{
	file_content := get_file_content(filename)
	lines := strings.Split(file_content, "\n")
	targets := lines_to_targets(lines)
	var res []string
	for _, target := range targets{
		res = append(res, target.ScanAddressOnPorts()...)
	}
	return res
}
