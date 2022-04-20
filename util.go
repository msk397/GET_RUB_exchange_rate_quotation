package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func barkPush(body, key, title string) string {
	mapdata := make(map[string]string)
	mapdata["body"] = body
	mapdata["device_key"] = key
	mapdata["title"] = title
	b, _ := json.Marshal(mapdata)
	// 超时时间：60秒
	client := &http.Client{Timeout: 60 * time.Second}
	pushurl := "https://barkapi.machangxin.top/push"
	header := "application/json; charset=utf-8"
	resp, err := client.Post(pushurl, header, bytes.NewBuffer(b))
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	result, _ := ioutil.ReadAll(resp.Body)
	return string(result) + "\n"
}

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func saveLog(mess string) {
	filePath := "./send.log"
	exist := FileExist(filePath)
	if !exist {
		fileCreate, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("文件创建失败", err)
		}
		err = fileCreate.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	_, err = write.WriteString(timeStr + " " + mess)
	if err != nil {
		return
	}
	//Flush将缓存的文件真正写入到文件中
	err = write.Flush()
	if err != nil {
		return
	}
}
