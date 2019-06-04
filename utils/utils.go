package utils

import (
	"fmt"
	"os"
)



// 创建文件夹，然后写入文件
func CreateFolderAndWriteToFile(fullPath string, fileFullName string, content string)  {
	ok := createFolderIfNotExists(fullPath)
	if ok {
		f, err := os.OpenFile(fullPath + "/" + fileFullName, os.O_CREATE | os.O_WRONLY, os.ModePerm)
		if err == nil {
			_, writeErr := f.WriteString(content)
			if writeErr == nil {
				fmt.Printf("🎉 " + fileFullName + " 创建成功\r\n")
			} else  {
				fmt.Printf("💥 " + fileFullName + " 创建失败\r\n")
			}
		}
	}
}

// 判断文件夹是否存在，如果不存在的话，就创建
func createFolderIfNotExists(fullPath string) bool {
	_, err := os.Stat(fullPath)
	if err != nil {
		mkdirError := os.MkdirAll(fullPath, os.ModePerm)
		if mkdirError != nil {
			return false
		} else {
			return true
		}
	} else {
		return true
	}
}