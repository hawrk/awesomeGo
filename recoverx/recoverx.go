// Package recoverx
/*
 Author: hawrkchen
 Date: 2022/4/6 16:46
 Desc:
*/
package recoverx

import (
	"fmt"
	"time"
)

//   1.main函数 进行了recover,main 函数panic时，会调用recover，然后该函数退出
//   2.要使 main 函数不崩溃，就要在子函数里捕获异常并recover,子函数退出，不影响 main函数
//   3. 如果子函数painc了，但是没有做recover，该panic会向上抛，也就是抛给main函数 ，main reocver该异常后，也会退出

// 4.如果子协程panic 而又没有recover时，主协程无法获得子协程异常， 整个程序会崩溃退出
// 5. 如果子协程panic recover，该子协程结束，不影响主协程
func Re_Main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in main")
		}
	}()

	go out()

	time.Sleep(time.Second * 3)
	panic("panic in main")
	select {}
}

func out() {
	for {
		tm := time.NewTimer(time.Second)
		select {
		case <-tm.C:
			fmt.Println("tick tick")
		}
	}
}
