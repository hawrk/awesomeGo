// Package httpx
/*
 Author: hawrkchen
 Date: 2022/3/22 9:39
 Desc: 关闭http服务时，如果直接kill 进程或者ctr+c 会导致未处理完成的请求直接关闭，需要调用shutdown处理完所有请求后
       才进行关闭
*/
package httpx

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func HttpMain() {
	addr := flag.String("server addr", ":8080", "server address")
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		fmt.Fprintln(w, "hello")
	})

	srv := http.Server{
		Addr:    *addr,
		Handler: handler,
	}

	processed := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT)
		<-c

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalln("server shotdown fail:", err)
		}
		log.Println("server shotdown gracefully")
		close(processed)
	}()
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(" server fail:", err)
	}
	<-processed
	fmt.Println("close successful")
}
