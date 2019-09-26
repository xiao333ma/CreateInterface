package main

import (
	"createInterface/utils"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
)

var userPath = utils.GetUserName() + "/Sites/"

var subPath = "www/interface/"

var defaultPath = userPath + subPath

func main() {
	path :=  os.Args[1]
	paths := strings.Split(path, "/")
	// 文件全名 test.php
	fullFileName := paths[len(paths) - 1]

	filePaths := strings.Split(path, fullFileName)
	// 拼接全部路径
	urlPath := strings.TrimPrefix(filePaths[0], "/")

	fullPath := defaultPath + urlPath
	// 获得文件名
	fileName := strings.Split(fullFileName, ".")[0]

	createPHPFile(fullPath, fileName)
	createJSONFile(fullPath, fileName)
	fmt.Printf("🍻 创建完成 可以打开以下连接查看 👇 \n")
	l, r := getURL(fullPath + fileName)
	fmt.Print(l + "\n" + r + "\n")
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

func getURL(fullPath string) (string, string){

	homeDir := strings.TrimPrefix(utils.GetUserName(), "/Users/")
	path := strings.TrimPrefix(fullPath, userPath)

	url := "/~" + homeDir + "/" + path + ".php"
	getLocalURL()


	localURL := getLocalURL() + url
	remoteURL :=  getRemoteURL() + url
	return localURL, remoteURL
}

func getLocalURL() string {
	return "http://127.0.0.1"
}

func getRemoteURL() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return "http://" + ipnet.IP.String()
			}

		}
	}
	return ""
}