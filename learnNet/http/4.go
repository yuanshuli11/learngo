package main

import (
	"fmt"
	"net/http"
	"os"
)

type fooHandler struct {
	w int
}

func (h *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", `{"tt":"33"}`)
	fmt.Println("fooHandler ServeHTTP")
	//获取进程id
	pid :=os.Getpid()
	fmt.Fprintf(w,"pid: %d\n", pid)
	//父进程id
	ppid :=os.Getppid()
	fmt.Fprintf(w,"ppid: %d\n", ppid)
}
type HandlersChain []http.Handler

func main() {
	var handlers HandlersChain
	//ListenAndServe使用指定的监听地址和处理器启动一个HTTP服务端。处理器参数通常是nil，这表示采用包变量
	handlers = append(handlers, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%q", "aaa")
	}))


	http.Handle("/countss", new(fooHandler))
	//
	http.Handle("/tt", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%q", "aaa")
	}))
	http.ListenAndServe(":8787", nil)
}
