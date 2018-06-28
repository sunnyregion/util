package util

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

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
	"strings"

	"github.com/disintegration/imaging"
)

//Jpg2Base64 ...
func Jpg2Base64(filename string) (base64Str string, err error) {
	src, err := imaging.Open(filename)

	if err != nil {
		base64Str = ``
	} else {
		buf := new(bytes.Buffer)
		err = jpeg.Encode(buf, src, nil)
		sendS3 := buf.Bytes()

		base64Str = `data:image/jpeg;base64,` + base64.StdEncoding.EncodeToString(sendS3)
	}

	//	src = imaging.Resize(src, 750, 0, imaging.Lanczos)

	return
}

//Base642Image base64 to Image
//目前支持png、jpeg
func Base642Image(image string) (img image.Image, picType string, err error) {

	coI := strings.Index(string(image), ",")
	rawImage := string(image)[coI+1:]

	// Encoded Image DataUrl //
	unbased, _ := base64.StdEncoding.DecodeString(string(rawImage))

	res := bytes.NewReader(unbased)

	switch strings.TrimSuffix(image[5:coI], ";base64") {
	case "image/png":
		picType = "image/png"
		img, err = png.Decode(res)
	case "image/jpeg":
		picType = "image/jpeg"
		img, err = jpeg.Decode(res)
	default:
		picType = "image/jpeg"
		img, err = jpeg.Decode(res)
	}
	return
}
