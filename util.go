package util

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

import (
	"crypto/md5"
	"encoding/hex"

	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	//json "github.com/sunnyregion/util/sunnyjson"

	"github.com/oklog/ulid"
	//"github.com/pquerna/ffjson/ffjson"
	UUID "github.com/satori/go.uuid"
)

const EARTH_RADIUS = 6371000

//Element interface
type Element interface{}

//SunnyMd5 md5 return
func SunnyMd5(str string) (keyMd5 string) {
	m := md5.Sum([]byte(str))
	keyMd5 = hex.EncodeToString(m[:])
	return
}

//GetDateTimeFormat ...
func GetDateTimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

//SunnyTimeNow 返回格式化好的当前字符串
// style 是day返回年月日，是time返回2006-01-02 15:04:05
func SunnyTimeNow(style string) (result string) {
	switch style {
	case "day":
		result = time.Now().Format("2006-01-02")
	case "time":
		result = time.Now().Format("2006-01-02 15:04:05")
	case "sunnyday":
		result = time.Now().Format("20060102")
	case "sunnytime":
		result = time.Now().Format("150405")
	default:
		result = time.Now().Format("2006-01-02 15:04:05.999999999 +0800 CST m=+0.999999999")
	}
	return
}

//SunnyTimeToStr 把时间类型的转换为字符串
func SunnyTimeToStr(t time.Time, style string) (result string) {
	switch style {
	case "day":
		result = t.Format("2006-01-02")
	case "time":
		result = t.Format("2006-01-02 15:04:05")
	default:
		result = t.Format("2006-01-02 15:04:05.999999999 +0800 CST m=+0.999999999")
	}
	return
}

//SunnyJSONToStr 把json变成字符串
func SunnyJSONToStr(jsondata interface{}) (result string) {
	content, _ := json.Marshal(jsondata)
	result = string(content)
	return
}

//SunnyIsNotExist 如果path不存在创建
func SunnyIsNotExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}

//SunnyIsNotExists 如果path不存在创建(创建多级目录)
func SunnyIsNotExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.MkdirAll(path, os.ModePerm)
	}
}

//SliceIndex 类似python的list.index()方法，找出slice里面是否包含查询的参数
// 返回值是-1表示没有搜索到。
func SliceIndex(a Element, i interface{}) int {
	if b, ok := a.([]int); ok {
		if c, ok1 := i.(int); ok1 {
			for indexC, v := range b {
				if v == c {
					return indexC
				}
			}
		}
	}
	if b, ok := a.([]string); ok {
		if c, ok1 := i.(string); ok1 {
			for indexC, v := range b {
				if v == c {
					return indexC
				}
			}
		}
	}
	if b, ok := a.([]float64); ok {
		if c, ok1 := i.(float64); ok1 {
			for indexC, v := range b {
				if v == c {
					return indexC
				}
			}
		}
	}
	return -1
}

//SunnyTypeof 返回数据类型
func SunnyTypeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

//SunnyStr2Time 输入字符串返回时间
//待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
// @Param strTime string
// @Return time.Time
func SunnyStr2Time(strTime string) time.Time {
	timeLayout := "2006-01-02 15:04:05.999999999"          //转化所需模板
	loc, _ := time.LoadLocation("Local")                   //重要：获取时区
	t, _ := time.ParseInLocation(timeLayout, strTime, loc) //使用模板在对应时区转化为time.time类型
	return t
}

//SunnyStr2USATime 输入字符串返回UTC时间
//待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
// @Param strTime string
// @Return time.Time error
func SunnyStr2USATime(strTime string) (time.Time, error) {
	timeLayout := "2006-01-02T15:04:05.999999999Z07:00"   //转化所需模板
	loc, _ := time.LoadLocation("Local")                  //重要：获取时区
	return time.ParseInLocation(timeLayout, strTime, loc) //使用模板在对应时区转化为time.time类型
}

//SunnyParseDuration 输入字符串返回增加或者减少的时间，1m是增加一分钟，-1m是减少一分钟
// @Param strSpan string,t time.Time
// @Return time.Time
func SunnyParseDuration(strSpan string, t time.Time) time.Time {
	m, _ := time.ParseDuration(strSpan)
	return t.Add(m)
}

//SunnyCompareTime 比较时间戳大小
// @Param t,u time.Time需要比较的时间
// @Param symbol string 标志包括 lt gt eq
func SunnyCompareTime(t, u time.Time, symbol string) (b bool) {
	switch symbol {
	case "lt":
		b = t.Before(u)
	case "gt":
		b = t.After(u)
	case "eq":
		b = t.Equal(u)
	}
	return b
}

//GetULID 得到UUID
func GetULID() ulid.ULID {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	return ulid.MustNew(ulid.Timestamp(t), entropy)
	// Output: 0000XSNJG0MQJHBF4QX1EFD6Y3
}

//GetULID 得到UUID
func GetULIDV4() UUID.UUID {
	return UUID.Must(UUID.NewV4(), nil)
}

//GetULID2Str ...
func GetULID2Str() string {
	var a [16]byte
	a = GetULID()
	return strings.ToUpper(hex.EncodeToString(a[:]))
}

/**
 * 截取字符串
 * @param str 字符串
 * @param start 开始的位置
 * @param length 截取多长，可是采用负数，就是从右往前计算
 */
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	if length < 0 {
		end = len(str) + length
	} else {
		end = start + length
	}

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

type Longlati struct {
	Long float64
	Lati float64
}

func (l *Longlati) GetLong() float64 {
	return l.Long
}
func (l *Longlati) SetLong(value float64) {
	l.Long = value
}
func (l *Longlati) SetLongFromString(value string) {
	f, _ := strconv.ParseFloat(value, 64)
	l.Long = f
}
func (l *Longlati) GetLati() float64 {
	return l.Lati
}
func (l *Longlati) SetLati(value float64) {
	l.Lati = value
}
func (l *Longlati) SetLatiFromString(value string) {
	f, _ := strconv.ParseFloat(value, 64)
	l.Lati = f
}

/**
 * 取得两个经纬度之间的差值
 * oldLL 原来的坐标
 * newLL 新的坐标
 */
func Get_two_points_distance(oldLL Longlati, newLL Longlati) int {
	lng1 := oldLL.GetLong() * math.Pi / 180
	lat1 := oldLL.GetLati() * math.Pi / 180
	lng2 := newLL.GetLong() * math.Pi / 180
	lat2 := newLL.GetLati() * math.Pi / 180
	calcLongitude := lng2 - lng1
	calcLatitude := lat2 - lat1
	stepOne := math.Pow(math.Sin(calcLatitude/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(calcLongitude/2), 2)
	stepTwo := 2 * math.Asin(math.Min(1, math.Sqrt(stepOne)))
	calculatedDistance := EARTH_RADIUS * stepTwo
	return int(calculatedDistance + 0.5)
}
