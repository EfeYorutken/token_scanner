package main

import(
	"strconv"
)

func get_indicies_of_args(args []string)[]int{
	var res []int
	cli_arguments := get_cli_arguments()
	for i := 0; i < len(args); i++{
		if contains(cli_arguments, args[i]){
			res = append(res, i)
		}
	}
	return res
}

func get_args_and_responsibilities(args []string)[][]string{
	var res [][]string
	indicies := get_indicies_of_args(args)
	for i := 0; i < len(indicies)-1; i++{
		to_parse := push(args[indicies[i]+1:indicies[i+1]], args[indicies[i]])
		res = append(res, to_parse)
	}
	last_index := len(indicies) - 1
	to_parse := push(args[indicies[last_index]+1:], args[indicies[last_index]])
	res = append(res, to_parse)
	return res
}

func execute_tool(args []string)[]string{
	coms := get_args_and_responsibilities(args)
	target_name := args[0]
	target := NewTarget(target_name, nil, nil)

	//this one can be done much better
	for _, to_parse := range coms{
		if to_parse[0] == "-p"{
			target.ports = to_parse[1:]
		}
		if to_parse[0] == "-r"{
			lower, _ := strconv.Atoi(to_parse[1:][0])
			upper, _ := strconv.Atoi(to_parse[1:][1])
			var prts_t_scn []string
			for port := lower; port <= upper; port++{
				prts_t_scn = append(prts_t_scn, strconv.Itoa(port))
			}
			target.ports = prts_t_scn
		}
		if to_parse[0] == "-t"{
			target.protocols = to_parse[1:]
		}
	}

	return target.ScanAddressOnPorts()
}
