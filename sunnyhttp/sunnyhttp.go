package sunnyhttp

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"

	SunnyUtil "github.com/sunnyregion/util"
)

//@title http请求

// ParamHttp 请求参数
type ParamHttp struct {
	BFile     bool   `json:"bfile"`     //是否是文件
	Fieldname string `json:"fieldname"` //传输参数名
	Name      string `json:"name"`      //名字
	Value     string `json:"value"`     //值
	//File      multipart.File `json:"file"`
	File io.Reader `json:"file"`
}

// String toString
func (this *ParamHttp) String() (result string) {
	b := strconv.FormatBool(this.BFile)
	result = SunnyUtil.SunnyStrJoin("{\r\nBFile：", b, "\r\nFieldname:", this.Fieldname, "\r\nFileName:", this.Name, "\r\nValue:", this.Value, "\r\n}")
	return
}

// @Param	ph []ParamHttp    true 参数
// @Param	url string        true 在配置文件里面的url地址
// @Return res *http.Response, err error
func DoPost(ctx context.Context, ph []ParamHttp, url string) (res *http.Response, err error) {
	var (
		b  bytes.Buffer
		fw io.Writer
	)
	w := multipart.NewWriter(&b)
	for _, v := range ph {
		if !v.BFile {
			fw, err = w.CreateFormField(v.Fieldname)
			if err != nil {
				return
			}
			if _, err = fw.Write([]byte(v.Value)); err != nil {
				return
			}
		} else {
			fw, err = w.CreateFormFile(v.Fieldname, v.Name)
			if err != nil {
				return
			}
			if _, err = io.Copy(fw, v.File); err != nil {
				return
			}
		}
	}

	w.Close()
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", w.FormDataContentType())
	// Submit the request
	// ctx, cancel := context.WithTimeout(context.Background(), time.Now().Add(3*time.Second))
	// defer cancel()
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err = client.Do(req)
	return
}
