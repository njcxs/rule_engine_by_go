package log

import (
	"io"
	"log"
	"os"
	"strings"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func init() {
	errFile, err := os.OpenFile(GetCurrentPath()+"/logs/nids.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}

	Info = log.New(io.MultiWriter(os.Stdout, errFile), "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(io.MultiWriter(os.Stdout, errFile), "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, errFile), "Error:", log.Ldate|log.Ltime|log.Lshortfile)

}
