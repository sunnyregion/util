package util

import (
	"github.com/axgle/mahonia"
	"github.com/sunnyregion/util/sunnyjson"
)

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

// GetI18NValue 取得I18N的值
// str := GetI18NValue(`{"zh-CN":"添加成功!","en-US":"Added successfully.","zh-TW":"添加成功!"} `,`en-US`)
func GetI18NValue(s, lang string) (result string, err error) {
	result = ``
	var (
		by  []byte
		dat map[string]interface{}
	)
	by = []byte(s)
	if err = sunnyjson.Unmarshal(by, &dat); err == nil {
		if str, ok := dat[lang].(string); ok {
			result = str
		} else {
			result = dat[`zh-CN`].(string)
		}
	}
	return result, err
}
