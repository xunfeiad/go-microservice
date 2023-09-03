package log

import (
	"fmt"
	"io"
	stlog "log" // 标准库log
	"net/http"
	"os"
)

var log *stlog.Logger

type fileLog string

func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

// 创建自定义的log文件：写入数据
func Run(destination string) {
	log = stlog.New(fileLog(destination), "[go] - ", stlog.LstdFlags|stlog.Ltime)
}

// 针对post请求body的内容全部读出来
func RegisterHandlers() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			msg, err := io.ReadAll(r.Body)
			if err != nil || len(msg) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(msg))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}

func write(message string) {
	log.Printf("%v\n", message)
}
