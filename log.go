package util

import (
	"log"
	"os"
)

//logStruct struct 为了把log写成多线程，节省时间
type sunnyLogStruct struct {
	data   interface{} //数据
	msg    string      //信息
	symbol string      //标志位
}

//doWritLogFile 写日志
func DoWritLogFile(data interface{}, msg string, symbol string, lf *os.File) {
	logch := make(chan sunnyLogStruct, 1)
	logch <- sunnyLogStruct{data, msg, symbol}
	go writeLogFile(logch, lf)
	defer close(logch)
}

// writeLogFile把写日志变成多线程
func writeLogFile(ch chan sunnyLogStruct, lf *os.File) {
	mych := <-ch
	message := ""
	log.SetOutput(lf)
	switch mych.symbol {
	case "str":
		message = mych.data.(string)
	default:
		message = SunnyJSONToStr(mych.data)
	}
	log.Println("\t"+message, "---"+mych.msg+"\r\n")
}
