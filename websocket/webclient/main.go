// Package websock
/*
 Author: hawrkchen
 Date: 2022/4/21 9:23
 Desc:
*/
package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
)

type WebSocketClientManager struct {
	Conn    *websocket.Conn
	Addr    string
	Path    string
	IsAlive bool
	Timeout int
}

func NewWsClientManager(ip, port, path string, timeout int) *WebSocketClientManager {
	addrstr := ip + ":" + port
	var conn *websocket.Conn
	return &WebSocketClientManager{
		Conn:    conn,
		Addr:    addrstr,
		Path:    path,
		IsAlive: false,
		Timeout: timeout,
	}
}

func main() {
	wsc := NewWsClientManager("127.0.0.1", "2021", "/ws", 100)
	wsc.start()

	select {}
}

func (w *WebSocketClientManager) DailServer() {
	u := url.URL{
		Scheme: "ws",
		Host:   w.Addr,
		Path:   w.Path,
	}
	log.Printf("connect to %s", u.String())
	var err error
	w.Conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Printf("err:", err)
		return
	}
	w.IsAlive = true
	log.Printf("connect to %s success", u.String())
}

func (w *WebSocketClientManager) SendMsg() {
	go func() {
		msg := `{"name":"整活吗"}`
		err := w.Conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Println("write msg:", msg)
		}
	}()
}
func (w *WebSocketClientManager) SendbybyMsg() {
	go func() {
		msg := "不搞了不搞了，吃不消"
		err := w.Conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Println("write msg:", msg)
		}
	}()
}

func (w *WebSocketClientManager) ReadMsg() {
	go func() {
		for {
			if w.Conn != nil {
				_, message, err := w.Conn.ReadMessage()
				if err != nil {
					log.Println("read：", err)
					w.IsAlive = false
					break
				}
				log.Printf("recv:%s", message)
			}
		}
	}()
}

func (w *WebSocketClientManager) start() {
	for {
		if !w.IsAlive {
			w.DailServer()
		}
		w.SendMsg()
		w.ReadMsg()
		time.Sleep(time.Second * time.Duration(w.Timeout))
		//w.SendbybyMsg()
		//time.Sleep(time.Second * time.Duration(w.Timeout))
	}
}
