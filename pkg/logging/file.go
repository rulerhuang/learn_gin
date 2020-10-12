package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	LogSavePath = "runtime/logs"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s/%s", getLogFilePath(), suffixPath)
}

func mkDir() {
	dir, _ := os.Getwd()
	// 创建目录以及其子目录，等效 mkdir -p 命令
	err := os.MkdirAll(fmt.Sprintf("%s/%s", dir, getLogFilePath()), os.ModePerm) // 0777
	if err != nil {
		panic(err)
	}
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission:%v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("fail to OpenFile:%v", err)
	}
	return handle
}
