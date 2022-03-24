// Package awesomeGo
/*
 Author: hawrkchen
 Date: 2022/3/8 10:59
 Desc:
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		//c <- struct{}{}
		//close(c)
	}()
	<-c
	fmt.Println("finish")

}
