// Package awesomeGo
/*
 Author: hawrkchen
 Date: 2022/3/8 10:59
 Desc:
*/
package main

import (
	"awesomeGo/bench"
	"awesomeGo/httpx"
	"fmt"
	"strings"
)

func main() {

	httpx.HttpMain()
	var str0 = "12345678901234567890"
	str1 := strings.Repeat(str0[:10], 1)
	fmt.Println("address str0:", &str0, ", str1:", &str1)
	fmt.Println("str1:", str1)

	return
	bench.JsoniterEncode()
	bench.JsonEncode()

}
