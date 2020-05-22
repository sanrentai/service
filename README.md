# service
将golang编译后的程序，注册成windows系统服务
## 简单用法，可以同时在当前目录和系统日志里记录日志内容，支持INFO,ERROR,WARN.
```
package main

import (
	"fmt"
	"net/http"

	"github.com/sanrentai/service"
)

func main() {
	prg := &service.Program{
		Name:        "test",
		DisplayName: "TEST",
		Description: "this is a test",
		Runfunc:     Run,
	}
	service.Run(prg)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	service.Info("hello world")
	fmt.Fprintf(w, "hello world")
}

func Run() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
```
在以管理员身份运行  cmd  
go build  
example.exe -s install   
example.exe -s uninstall




