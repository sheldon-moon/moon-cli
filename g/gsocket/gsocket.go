package gsocket

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
	"strings"
	"sync"
	"time"
)

var (
	ws *websocket.Conn
	lock sync.RWMutex
	state bool
)

func Conn(uri string, origin  string, token string) bool {
	_uri_ := strings.Join([]string{uri,token}, "")
	logrus.Debug("uri:", _uri_)
	conn, err := websocket.Dial(_uri_, "", origin)
	if err != nil {
		logrus.Error("websocket conn error", err)
		return false
	}
	ws = conn
	state = true;
	go clientHandler(ws);
	return true
}

func clientHandler(ws *websocket.Conn)  {
	lock.Lock()
	defer lock.Unlock()
	defer ws.Close()
	for {
		request := make([]byte, 4096);
		readLen, err := ws.Read(request)
		if err != nil {
			logrus.Error("recv msg err" , err)
			break;
		}
		if readLen == 0{
			logrus.Error("websocket 关闭")
			state = false
		}else{
			fmt.Println(string(request[:readLen]))
		}
	}
}

func Retry(uri string, flag string, token string) {

}

func WriteMsg(msg interface{}) {
	if state {
		b_msg, _ := json.Marshal(msg)
		ws.Write(b_msg)
	}
}

func Test()  {
	for ; ;  {
		WriteMsg("test")
		time.Sleep(time.Second)
	}
}

