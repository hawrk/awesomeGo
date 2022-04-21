// Package awesomeGo
/*
 Author: hawrkchen
 Date: 2022/3/8 10:59
 Desc:
*/
package main

import (
	"fmt"
	"strconv"
	"time"
)

type Myst struct {
	Id   int
	Name string
	Sum  int
}

type MyMap struct {
	Order map[string]*Myst
}

var GlobalMap = MyMap{
	Order: make(map[string]*Myst),
}

func main() {
	date := "202204190800"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt, _ := time.ParseInLocation("200601021504", date, loc)
	out := tt.Unix()

	fmt.Println("out:", out)

	fmt.Println("now:", time.Now().Unix())
	return

	year, _ := strconv.Atoi(date[:4])
	mon, _ := strconv.Atoi(date[4:6])
	day, _ := strconv.Atoi(date[6:8])
	hour, _ := strconv.Atoi(date[8:10])
	min, _ := strconv.Atoi(date[10:])
	fmt.Println("get :", year, mon, day, hour, min)
	start := time.Date(year, time.Month(mon), day, hour, min, 0, 0, time.UTC).Add(-time.Minute * time.Duration(1)).Format("200601021504")

	fmt.Println("start:", start)

	//c := make(chan struct{})
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	//c <- struct{}{}
	//	//close(c)
	//}()
	//<-c
	//fmt.Println("finish")

}
