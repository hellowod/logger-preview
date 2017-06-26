package utils

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/hellowod/log-preview/logcore/log"
)

func GetPachName(url string) string {
	if len(url) <= 0 {
		log.Info("url pach is err.")
	}
	str := strings.Replace(url, "/", "", -1)
	return strings.TrimSpace(str)
}

// get file string context
func ReadFile(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Error("open file error." + path)
		panic(err)
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		log.Error("read file error." + path)
		panic(err)
	}
	return string(fd)
}

// write string to file
func WriteFile(fp string, t string) error {
	f, err := os.Create(fp)
	if err != nil {
		log.Error("create file error." + fp)
		return err
	}
	defer f.Close()
	_, errr := f.WriteString(t)
	if errr != nil {
		log.Error("write file error." + fp)
		return errr
	}

	return nil
}
