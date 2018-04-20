// Copyright 2018 sunny authors
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
//这个文件作为我处理一些图片使用
package util

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/jpeg"

	"github.com/disintegration/imaging"
)

func Jpg2Base64(filename string) (base64Str string) {
	src, err := imaging.Open(`./e.jpg`)
	if err != nil {
		fmt.Println(err.Error())
	}

	src = imaging.Resize(src, 150, 0, imaging.Lanczos)

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, src, nil)
	send_s3 := buf.Bytes()

	base64Str = `data:image/png;base64,` + base64.StdEncoding.EncodeToString(send_s3)
	return
}
