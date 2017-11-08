// Copyright 2017 sunnyini authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Source code and project home:
//
// https://github.com/sunnyregion/util
//
// Installation:
//
// go get  github.com/sunnyregion/util
//
// Example:
//
//		import "github.com/sunnyregion/util"
//这个文件作为我基础使用的一些函数的寄存处
package util

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/pquerna/ffjson/ffjson"
)

// md5 return
func SunnyMd5(str string) (keyMd5 string) {
	m := md5.Sum([]byte(str))
	keyMd5 = hex.EncodeToString(m[:])
	return
}

// 返回格式化好的当前字符串
// style 是day返回年月日，是time返回2006-01-02 15:04:05
func SunnyTimeNow(style string) (result string) {
	switch style {
	case "day":
		result = time.Now().Format("2006-01-02")
	case "time":
		result = time.Now().Format("2006-01-02 15:04:05")
	default:
		result = time.Now().Format("2006-01-02 15:04:05")
	}
	return
}

//// json return
//func StructToJSON(data map[interface{}]interface{}) string {
//	content, _ := ffjson.Marshal(data)
//	s := string(content)
//	return s

//}

func SunnyJsonToStr(jsondata interface{}) (result string) {
	content, _ := ffjson.Marshal(jsondata)
	result = string(content)
	return
}
