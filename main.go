package main

import (
	logger "github.com/sirupsen/logrus"
	)

//import "moon/g/gsocket"

//var origin = "http://127.0.0.1:9011/"
//var url = "ws://127.0.0.1:9011/websocket/2"
type User struct {
	user string
}

func main() {
	//gsocket.Conn(url, origin,"33");
	//go gsocket.Test()
	//u, _ := url.Parse("http://127.0.0.1:8081/basic/events")


	// Second argument is the category on which we want to receive events
	//c := glpclient.NewClient(u, "farm")
	//c.Start()

	user := User{"hello"}
	logger.Warn("moon start...",user)
	//启动HTTP服务
}

