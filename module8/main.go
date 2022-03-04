package main

import (
	"fmt"
	"github.com/golang/glog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	initConfig() // 加载配置
	fmt.Printf("conf:%+v \n", conf)

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/prestart", preStartHandler)
	mux.HandleFunc("/ready", readyHandler)
	mux.HandleFunc("/healthz", healthZHandler)
	mux.HandleFunc("/homework", handler)
	fmt.Println("env PLAYER_INITIAL_LIVES:", os.Getenv("PLAYER_INITIAL_LIVES")) // 环境变量

	var done = make(chan struct{})
	go handleOsSig(done)
	go func() {
		err := http.ListenAndServe(":80", mux)
		if err != nil {
			glog.Error(err)
		}
		select {
		case done <- struct{}{}:
		default:
		}
	}()
	<-done
}

func handleOsSig(ch chan<- struct{}) {
	var s = make(chan os.Signal)
	signal.Notify(s, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-s
	fmt.Println("catch signal ... do something... -> sleep 5s")
	time.Sleep(5 * time.Second)
	ch <- struct{}{}
}

func handler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		for i, str := range v {
			if i == 0 {
				w.Header().Set(k, str)
			} else {
				w.Header().Add(k, str)
			}
		}
	}

	ver := os.Getenv("VERSION")
	w.Header().Set("VERSION", ver)

	ip := r.RemoteAddr
	glog.Info("client ip ", ip, ",status code ", http.StatusOK)

	res := fmt.Sprintf("hello world handler, VERSION %s , client ip %s \n", ver, ip)
	_, _ = w.Write([]byte(res))
	w.WriteHeader(http.StatusOK)
}

func healthZHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world module8 healthZHandler,Custom-Header:" + r.Header.Get("Custom-Header"))
	_, _ = w.Write([]byte("hello world module8 healthZHandler,Custom-Header:" + r.Header.Get("Custom-Header")))
	w.WriteHeader(http.StatusOK) // 4、当访问 localhost/healthz 时，应返回 200
}

func readyHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("ready check:OK")
	w.WriteHeader(http.StatusOK) // 4、当访问 localhost/healthz 时，应返回 200
}

func preStartHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("pre start:OK")
	w.WriteHeader(http.StatusOK) // 4、当访问 localhost/healthz 时，应返回 200
}
func indexHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("hello world module8 !!! index ."))
	w.WriteHeader(http.StatusOK) // 4、当访问 localhost/healthz 时，应返回 200
}
