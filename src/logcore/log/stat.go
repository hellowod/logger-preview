package log

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const statFlags = Ldate | Ltime

var (
	stat     = NewLogger(nil, "", statFlags)
	statPath = ""
)

func init() {
	stat.SetPrefix("stat")
}

func Stat(a ...interface{}) {
	prepareStat()
	stat.Output(0, fmt.Sprint(a...))
}

func SetStatFilePrefix(prefix string) {
	stat.SetPrefix(prefix)
}

func SetStatFilePath(path string) {
	statPath = path
}

var lastDate int

func prepareStat() {
	newDate := currentDate()
	if lastDate != newDate {
		lastDate = newDate
		switchFile()
	}
}

func switchFile() {
	filename := fmt.Sprintf("%s_%d.log", stat.prefix, lastDate)
	file, err := os.OpenFile(filepath.Join(statPath, filename), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		Fatal("Cannot open stat file", err)
	}

	stat.SetWriter(file)
}

func currentDate() int {
	year, month, day := time.Now().Date()
	return year*10000 + int(month)*100 + day
}
