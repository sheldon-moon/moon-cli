package log

import (
	"log"
	"os"
)

type Logger struct {
	DEBUG *log.Logger
	INFO *log.Logger
	WARN *log.Logger
	ERROR *log.Logger
}

func Ldir(path string)   {
	if err :=os.MkdirAll(path,os.ModePerm) ; err!=nil{
			log.Print("log path is error , create local file"  ,  path  , err)
			os.Create("./log.out")
		return
	}
}


func LDefault() {
	log.Println("这是默认的格式\n")
}

func LDate() {
	log.SetFlags(log.Ldate)
	log.Println("这是输出日期格式\n")
}

func LTime() {
	log.SetFlags(log.Ltime)
	log.Println("这是输出时间格式\n")
}

func LMicroseconds() {
	log.SetFlags(log.Lmicroseconds)
	log.Println("这是输出微秒格式\n")
}

func LLongFile() {
	log.SetFlags(log.Llongfile)
	log.Println("这是输出路径+文件名+行号格式\n")
}

func LShortFile() {
	log.SetFlags(log.Lshortfile)
	log.Println("这是输出文件名+行号格式\n")
}

func LUTC() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)
	log.Println("这是输出 使用标准的UTC时间格式 格式\n")
}


