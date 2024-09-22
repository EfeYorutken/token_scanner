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
	for i := 0; i < len(lines); i++{
		if len(lines[i]) != 0{
			line := lines[i]
			split_line := strings.Split(line, " ")
			target_name := split_line[0]
			port := string_to_int(split_line[1])
			protocol := string(split_line[2])
			new_target := NewTarget(target_name, []int{port}, []string{protocol})
			res = append(res, new_target)
		}
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
