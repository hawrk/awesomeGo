// Package awesomeGo
/*
 Author: hawrkchen
 Date: 2022/3/8 10:59
 Desc:
*/
package main

import (
	"fmt"
	"github.com/spf13/cast"
	"time"
)

func main() {
	now := cast.ToInt64(time.Now().Add(-time.Minute * 1).Format("200601021504"))
	fmt.Println("now:", now)

	return

	min := time.Now().Add(-time.Minute * 1).Format("200601021504")
	fmt.Println("min:", min)
	var lastqty uint64
	lastqty = 722
	var orderqty uint64
	orderqty = 5000
	//TODO:
	//FIXME:
	//HACK:

	dealRate := float64(lastqty) / float64(orderqty)

	fmt.Println("dealRate:", dealRate)

	//c := make(chan struct{})
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	//c <- struct{}{}
	//	//close(c)
	//}()
	//<-c
	//fmt.Println("finish")

}
