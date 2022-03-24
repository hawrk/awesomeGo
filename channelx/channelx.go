// Package channelx
/*
 Author: hawrkchen
 Date: 2022/3/24 19:39
 Desc:
*/
package channelx

import (
	"fmt"
	"time"
)

// 等待事件， close一个channel 就可以了
func CloseAChannel() {
	c := make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		//c <- struct{}{}   // 可以往channel里写入标识
		//close(c)   // 关闭一个channel也可以
	}()
	<-c
	fmt.Println("finish")
}

// 取最快的结果，一般用在多协程执行，取最快的就行了
func GetQuickilyResult() {
	c := make(chan string, 10) // channel 数组 定义为 chs := make([]chan string, 10)   表示有10个int 类型的channel，每个channel的缓冲区需要自定义
	for i := 0; i < cap(c); i++ {
		go dosomething(c)
	}
	out := <-c
	fmt.Println("out:", out)
}
func dosomething(result chan string) {
	// do something
	result <- "aaa"
}

// 多channel的定义和使用
func MultiChannels() {
	chs := make([]chan int, 10)
	for i := 0; i < cap(chs); i++ {
		chs[i] = make(chan int)
		go count(chs[i])
	}

	for _, ch := range chs {
		fmt.Println("ch", <-ch)
	}
}
func count(ret chan int) {
	ret <- 555
}
