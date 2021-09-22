package utils

import (
	"fmt"
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, err := os.OpenFile(
		logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666,
	)
	if err != nil {
		fmt.Println("error occurred!")
		log.Fatal(err)
	}
	// 標準出力と作成したlogfileに書き込むための変数を作成
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	// logの形式を指定する(時間日付のフォーマット)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
