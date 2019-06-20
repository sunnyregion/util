package util

// SunnyMsgStruct 统一返回的信息结构体
type SunnyMsgStruct struct {
	Msg    string      `json:msg`
	Status bool        `json:status`
	Data   interface{} `json:data`
}
