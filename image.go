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
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"math"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"gocv.io/x/gocv"
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

// IsOrientationZero  是否是正向的照片
// @Param f *os.File  图片文件句柄
// @Return b bool, e error
func IsOrientationZero(f *os.File) (b bool, e error) {
	exif.RegisterParsers(mknote.All...)
	x, err := exif.Decode(f)
	if err != nil {
		e = err
	} else {
		camModel, err := x.Get(exif.Orientation)
		if err == nil {
			if camModel.Val[1] == 1 {
				b = true
				e = nil
			}
		} else {
			e = errors.New(`File size is too big.`)
		}

	}

	return
}

/*
//旋转角度	   参数
//0°   	    1
//顺时针90°	6
//逆时针90°	8
//180°     	3
*/
// GetOrientation  取得照片旋转的方向
// @Param f *os.File  图片文件句柄
// @Return b bool, e error
func GetOrientation(f *os.File) (val int, e error) {
	exif.RegisterParsers(mknote.All...)
	x, err := exif.Decode(f)
	if err != nil {
		e = err
	}
	camModel, _ := x.Get(exif.Orientation)
	val = int(camModel.Val[1])
	return
}

// GetImageSize 取得图片尺寸
// @Param file *os.File 文件句柄
// @Return width,height int
func GetImageSize(file *os.File) (width, height int) {
	c, _, _ := image.DecodeConfig(file)
	width = c.Width
	height = c.Height
	return
}

//GetImageSizeAndCount 取得照片里面人脸数量和面积
//https://github.com/opencv/opencv/tree/master/data/haarcascades
//https://blog.csdn.net/yangleo1987/article/details/52858706
func GetFaceSizeAndCount(f *os.File) (faceCount int, long, width, area float64, rect image.Rectangle, e error) {
	b := make([]byte, 1024000)
	f.ReadAt(b, 0)
	img, err := gocv.IMDecode(b, gocv.IMReadColor)
	if err != nil {
		e = errors.New("This is not a image.")
	} else {
		xmlFile := "face.xml"
		classifier := gocv.NewCascadeClassifier()
		defer classifier.Close()
		if !classifier.Load(xmlFile) {
			e = errors.New("Error reading cascade file: " + xmlFile)
		} else {
			rects := classifier.DetectMultiScale(img)
			faceCount = len(rects)
			e = nil
			if faceCount < 2 {
				min := rects[faceCount-1].Min
				max := rects[faceCount-1].Max
				long = math.Abs(float64(max.X - min.X))
				width = math.Abs(float64(max.Y - min.Y))
				area = long * width
				rect = rects[0]
			} else if faceCount < 3 {
				min1 := rects[0].Min
				max1 := rects[0].Max
				l1 := math.Abs(float64(max1.X - min1.X))
				w1 := math.Abs(float64(max1.Y - min1.Y))
				a1 := l1 * w1
				min2 := rects[1].Min
				max2 := rects[1].Max
				l2 := math.Abs(float64(max2.X - min2.X))
				w2 := math.Abs(float64(max2.Y - min2.Y))
				a2 := l2 * w2

				if a1 > a2 {
					area = a1
					long = l1
					width = w1
					rect = rects[0]
				} else {
					area = a2
					long = l2
					width = w2
					rect = rects[1]
				}
			}
		}
	}
	return
}
