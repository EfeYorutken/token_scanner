package main

import(
//	"fmt"
)

func get_cli_arguments()[]string{
	return []string{"-p", "-r", "-t", "-sG", "-sB", "-s", "-f"}
}

func get_indicies_of_args(args []string)[]int{
	var indicies []int
	cli_arguments := get_cli_arguments()
	for i := 0; i < len(args); i++{
		if contains(cli_arguments, args[i]){
			indicies = append(indicies, i)
		}
	}
	return indicies
}

func GetArgsAndResponsibilities(args []string)[][]string{
	var args_and_resps [][]string


	arg_indicies := get_indicies_of_args(args)


	for i := 0; i < len(arg_indicies)-1; i++{


		arg_index := arg_indicies[i]
		next_arg_index := arg_indicies[i+1]


		arg := args[arg_index]

		resposibilities := args[arg_index + 1 : next_arg_index]

		arg_and_resp := push(resposibilities, arg)

		args_and_resps = append(args_and_resps, arg_and_resp)
	}

	last_arg_index := arg_indicies[len(arg_indicies) - 1]
	last_statement := args[last_arg_index : ]

	last_arg := last_statement[0]
	last_responsibility := last_statement[1:]

	last_arg_and_resp := push(last_responsibility, last_arg)

	args_and_resps = append(args_and_resps, last_arg_and_resp)

	return args_and_resps
}
