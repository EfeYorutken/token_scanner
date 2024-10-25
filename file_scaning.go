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
	for _, line := range lines{
		if len(line) != 0{


			line_arr := strings.Split(line, " ")

			target_name := line_arr[0]
			target_port := line_arr[1]
			target_protocol := line_arr[2]


			//might be in stdlib
			target_names := map_to(res, func(target Target)string{
				return target.name 
			})
			target_search_res := -1
			for i := 0; i < len(target_names); i++{
				if(target_names[i] == target_name){
					target_search_res = i
					break
				}
			}

			if  target_search_res == -1{
				res = append(res, Target{target_name, []string{target_port}, []string{target_protocol}})
			}else{

				found_has_port := contains(res[target_search_res].ports, target_port)
				found_had_protocol := contains(res[target_search_res].protocols, target_protocol)

				if !found_has_port{
					res[target_search_res].ports = push(res[target_search_res].ports, target_port)
				}
				if !found_had_protocol{
					res[target_search_res].protocols = push(res[target_search_res].protocols, target_protocol)
				}

			}

		}
	}
	return res
}

func GetTargetsFromFile(filename string)[]Target{
	file_content := get_file_content(filename)
	lines := strings.Split(file_content, "\n")
	return lines_to_targets(lines)
}
