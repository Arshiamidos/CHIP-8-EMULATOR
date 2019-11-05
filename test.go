package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(0 == '0', byte(0), byte('0'))
	n, e := strconv.ParseInt(fmt.Sprintf("%s", "2"), 16, 8)
	if e != nil {
		panic(e)
	}
	fmt.Println(int(n) & 0xFF)
}
