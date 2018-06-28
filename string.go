package util

import "github.com/axgle/mahonia"

// ConvertToString ... gbk装换成utf8
//str := "乱码的字符串变量"
//str  = ConvertToString(str, "gbk", "utf-8")
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
