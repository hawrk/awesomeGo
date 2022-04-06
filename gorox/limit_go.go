// Package gorox
/*
 Author: hawrkchen
 Date: 2022/3/31 9:59
 Desc:  限制协程并发数量， 用channel 和waitGroup实现
*/
package gorox

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

const NUM = 50

func Main_gox() {
	ch := make(chan struct{}, 5) // 限制5个协程
	for i := 0; i < NUM; i++ {
		wg.Add(1)
		ch <- struct{}{}
		go func(num int) {
			defer wg.Done()
			fmt.Println("go here,i:", num, ", time :", time.Now())
			time.Sleep(3 * time.Second)
			<-ch
		}(i)
	}
	wg.Wait()
	fmt.Println("end")
}
