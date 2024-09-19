package main

import (
	//	"fmt"
)

func get_cli_arguments()[]string{
	return []string{"-p", "-r", "-t"}
}

func push[T any](arr []T, val T) []T{
	res := []T{val}
	for i := 0; i < len(arr); i++{
		res = append(res, arr[i])
	}
	return res
}

func map_to[T any, U any](from []T, fn func(T)U)[]U{
	var res []U

	for i := 0; i < len(from); i++{
		res = append(res, fn(from[i]))
	}

	return res
}

func index_of[T comparable](arr []T, val T, begin int)int{
	for i := begin; i < len(arr); i++{
		if arr[i] == val{
			return i
		}
	}
	return -1
}

func contains[T comparable	](arr []T, val T)bool{
	for i := 0; i < len(arr); i++{
		if arr[i] == val{
			return true
		}
	}
	return false
}

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
	for i := 0; i < len(coms); i++{
		command := coms[i][0]
		switch command{
		case "-p"	:
			target.ports = map_to(coms[i][1:], string_to_int)
		case "-r":
			upper := string_to_int(coms[i][1])
			lower := string_to_int(coms[i][0])
			for i := lower; i < upper+1; i++{
				target.ports = append(target.ports, i)
			}
		case "-t":
			target.protocols = coms[i][1:]
		}
	}

	return target.ScanAddressOnPorts()
}
