package main

import (
	"strconv"
	"strings"
)

// Ex004 takes a string of comma-seperated numbers and returns a slice of int
func Ex004(input string) []int {
	pureStr := strings.ReplaceAll(input, " ", "")
	strLst := strings.Split(pureStr, ",")

	ret := []int{}
	for i := 0; i < len(strLst); i++ {
		intVar, _ := strconv.Atoi(strLst[i])
		ret = append(ret, intVar)
	}
	return ret
}
