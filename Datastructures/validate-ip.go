package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	ip := "256.33.195.255"
	res := validate(ip)
	fmt.Println("res : ",res)
}

func validate(ip string)bool{
	var b []string
	b = strings.Split(ip,".")
	if len(b) != 4{
		return false
	}
	for _, bt := range b{
		n, err := strconv.Atoi(bt)
		if err !=nil{
			return false
		}
		if n < 0 || n > 255{
			return false
		}
	}
	return true
}
