package common

import (
	"fmt"
	"os"
	"time"
)

func init() {
	CreateDateDir()
}

type logConfig struct {
	Gin string
	Zap string
}

var LogConfig = &logConfig{
	Zap: fmt.Sprintf("./log/%v/zapLog/%v.log", GameName, GetLogName()),
	Gin: fmt.Sprintf("./log/%v/ginLog/%v.log", GameName, GetLogName()),
}

func CreateDateDir() {

	err := os.MkdirAll(fmt.Sprintf("./log/%v/zapLog", GameName), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	err = os.MkdirAll(fmt.Sprintf("./log/%v/ginLog", GameName), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
func GetLogName() string {
	folderName := time.Now().Format("20060102150405")
	return folderName
}
