package util

import (
	"os"
)

// GetFileSize 取得文件大小
// @Param file *os.File 文件句柄
// @Return s int64
func GetFileSize(file *os.File) (s int64) {
	if f, err := file.Stat(); err == nil {
		s = f.Size()
	}
	return
}
