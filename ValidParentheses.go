package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {

	if len(s) > 0 {
		if strings.Contains(s, "()") {
			strings.ReplaceAll(s, "()", " FOUND ")
			fmt.Println("Inside ()")
			fmt.Println(s)
		}
		if strings.Contains(s, "{}") {
			strings.ReplaceAll(s, "{}", "")
		}
		if strings.Contains(s, "[]") {
			strings.ReplaceAll(s, "[]", "")
			fmt.Println("Inside []")
			fmt.Println(s)
		}
		//fmt.Println("Didnt go inside")
	}

	if len(s) == 0 {
		return true
	}

	return false
}
