package main

import (
)

func ternary[T any](cond bool, if_true T, if_false T)T{
	if cond{
		return if_true
	}
	return if_false
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

