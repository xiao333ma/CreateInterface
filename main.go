package main

import (
	"createInterface/utils"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var defaultPath = "/Users/xiaoma/Sites/www/dd/"

func main() {
	path :=  os.Args[1]
	paths := strings.Split(path, "/")
	// 文件全名 test.php
	fullFileName := paths[len(paths) - 1]

	filePaths := strings.Split(path, fullFileName)
	// 拼接全部路径
	fullPath := defaultPath + filePaths[0]
	// 获得文件名
	fileName := strings.Split(fullFileName, ".")[0]

	createPHPFile(fullPath, fileName)
	createJSONFile(fullPath, fileName)
	fmt.Printf("🍻 创建完成 \n")
	commandString := "open " + fullPath + "/" + fileName + ".json" + " -a Visual\\ Studio\\ Code"
	_ = exec.Command("/bin/bash", "-c", commandString).Run()

}

func createPHPFile(fullPath string, fileName string)  {
	phpContent := `<?php
header('Content-type:application/json;charset=utf-8');
$json_string = file_get_contents('./__file__.json');
echo $json_string;
?>
`
phpContent = strings.Replace(phpContent, "__file__", fileName, -1)
utils.CreateFolderAndWriteToFile(fullPath, fileName + ".php", phpContent)
}

func createJSONFile(fullPath string, fileName string)  {
	jsonContent := `{
    "code": 10000,
    "msg": "成功",
    "friendlyMsg": "成功",
    "data": {
    }
}
`
	utils.CreateFolderAndWriteToFile(fullPath, fileName + ".json", jsonContent)

}