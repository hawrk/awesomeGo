// Package webserver
/*
 Author: hawrkchen
 Date: 2022/4/21 9:47
 Desc:
*/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upGrader = websocket.Upgrader{
	HandshakeTimeout: 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Ping(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("upgrader err:", err)
		return
	}
	defer ws.Close()

	for {
		// 收
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read msg fail:", err)
		}
		log.Println("get messagetype:", mt, ", message:", string(message))
		// 回
		if string(message) == "我准备好了，来搞我吧" {
			for i := 0; i < 10; i++ {
				time.Sleep(time.Second * 1)
				ret := fmt.Sprintf("是吗，搞你 %d 次", i)
				log.Println("send msg :", ret)
				err = ws.WriteMessage(mt, []byte(ret))
				//ws.WriteJSON()
				if err != nil {
					break
				}
			}
		} else {
			break
		}
	}
	// 这里如果不阻塞，处理完成后，就是基本的socket返回了
}

func main() {
	r := gin.Default()
	r.GET("/v1/hello-websocket", Ping)

	r.Run(":3088")
}
