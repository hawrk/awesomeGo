// Package nhooyr
/*
 Author: hawrkchen
 Date: 2022/4/21 15:59
 Desc:
*/
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "http, hello")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, req *http.Request) {
		conn, err := websocket.Accept(w, req, nil)
		if err != nil {
			log.Println("web socket error :", err)
			return
		}
		defer conn.Close(websocket.StatusInternalError, "internal error")

		var v interface{}
		if err = wsjson.Read(context.Background(), conn, &v); err != nil {
			log.Println("error read:", err)
			return
		}
		log.Printf("rece msg:%+v", v)

		for i := 0; i < 100; i++ {
			msg := fmt.Sprintf("整活整活，第 %d 次", i)
			if err := wsjson.Write(context.Background(), conn, msg); err != nil {
				log.Println("err write :", err)
			}
			time.Sleep(time.Second * 1)
		}

		conn.Close(websocket.StatusNormalClosure, "")
	})

	http.ListenAndServe(":2021", nil)
}
