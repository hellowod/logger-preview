package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"logcore/cache"
	"logcore/log"
	"logcore/static"
	"logcore/utils"
	"net/http"
	"strings"
)

func RunServer() {
	http.HandleFunc(static.NetPathRoot, dispacher)
	err := http.ListenAndServe(static.NetURL, nil)
	if err != nil {
		log.Info("Http Error Server: ", err)
	}
	cache.LogCache.Content = newCache()
}

func dispacher(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	url := utils.GetPachName(req.URL.Path)
	if url == static.ReqLog {
		processReqLog(resp)
	}
	if url == static.ReqPut {
		processPutLog(req)
	}
}

func processReqLog(resp http.ResponseWriter) {
	if len(cache.LogCache.Content) <= 0 {
		fmt.Fprintln(resp, static.DefaultLog)
		return
	}
	tpl := template.New("htmltpl")
	tpl, _ = tpl.Parse(utils.ReadFile("templete/index.tpl"))

	buf := new(bytes.Buffer)
	err := tpl.Execute(buf, cache.LogCache)
	if err != nil {
		fmt.Fprintln(resp, static.DefaultError)
		return
	}
	fmt.Fprintln(resp, buf.String())
}

func processPutLog(req *http.Request) {
	if len(cache.LogCache.Content) > static.LogMax {
		cache.LogCache.Content = cache.LogCache.Content[static.LogSep:]
	}
	info := cache.LogInfo{}
	for k, v := range req.Form {
		key := strings.TrimSpace(k)
		value := strings.Join(v, "")
		if key == static.KeyTime {
			info.Time = value
		} else if key == static.KeyType {
			info.Type = value
		} else if key == static.KeyCondition {
			info.Condition = value
		} else if key == static.KeyStackTrace {
			info.StackTrace = value
		} else if key == static.KeyOperatingSystem {
			info.OperatingSystem = value
		} else if key == static.KeyDevice {
			info.Device = value
		} else if key == static.KeyDeviceIP {
			info.DeviceIP = value
		} else if key == static.KeyDeviceModel {
			info.DeviceModel = value
		}
	}
	cache.LogCache.Content = append(cache.LogCache.Content, info)
	log.Info(info)
}

func newCache() []cache.LogInfo {
	return make([]cache.LogInfo, static.LogMin)
}
